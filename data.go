package main

import (
	"github.com/satori/go.uuid"
	"net/http"
	"time"
  "fmt"
  "log"
  _ "github.com/lib/pq"
  "database/sql"
)

var db *sql.DB

func initSQLDatabase(dbConnStr string) {
  var err error
  db, err = sql.Open("postgres", dbConnStr)
	if err != nil {
    fmt.Println(err)
		log.Fatal(err)
	}
  if err = db.Ping(); err != nil {
		fmt.Println(err)
    log.Fatal(err)
  }
  fmt.Println("Connected to SQL database.")
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Entry struct {
	ID        uuid.UUID	`json:"-"`
	ContentID uuid.UUID `json:"id"`
	Locale    string    `json:"locale"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

type Err struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func NewestEntryForContentID(contentID string) (*Entry, error) {
  var e Entry
  stmt, err := db.Prepare("select distinct entry.id, entry.content_id, locale.code as locale, entry.timestamp, entry.data from epic.entry inner join epic.locale on entry.locale_id=locale.id where entry.content_id = $1 order by entry.timestamp desc")
  if err != nil {
  	return nil, err
  }
  err = stmt.QueryRow(contentID).Scan(&e.ID, &e.ContentID, &e.Locale, &e.Timestamp, &e.Data)
  if err != nil {
  	return nil, err
  }
  return &e, nil
}

func CreateEntryForContentID(e *Entry) error {
  var localeID string
  l_stmt, err := db.Prepare("select locale.id from epic.locale where locale.code = $1")
  if err != nil {
    return err
  }
  defer l_stmt.Close()
  err = l_stmt.QueryRow(e.Locale).Scan(&localeID)
  if err != nil {
    return err
  }
  e_stmt, err := db.Prepare("insert into epic.entry (id, content_id, locale_id, timestamp, data) values ($1, $2, $3, $4, $5)")
  if err != nil {
  	return err
  }
  defer e_stmt.Close()
  _, err = e_stmt.Exec(e.ID.String(), e.ContentID.String(), localeID, e.Timestamp, e.Data)
  if err != nil {
  	return err
  }
  return nil
}

func AllContentForTag(Tag string, AppID string) ([]Entry, error) {
	stmt, err := db.Prepare("select entry.id, entry.content_id, locale.code as locale, entry.timestamp, entry.data from epic.entry inner join epic.locale on entry.locale_id=locale.id inner join epic.content on entry.content_id=content.id inner join epic.content_tag on content.id=content_tag.content_id inner join epic.tag on tag.id=content_tag.tag_id where tag.value=$1 and content_tag.application_id=$2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(Tag, AppID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]Entry, 0)
	for rows.Next() {
		e := Entry{}
		err := rows.Scan(&e.ID, &e.ContentID, &e.Locale, &e.Timestamp, &e.Data)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func CreateTag(tag string, appID string) error {
  e_stmt, err := db.Prepare("insert into epic.tag (id, application_id, value) values ($1, $2, $3)")
  if err != nil {
  	return err
  }
  defer e_stmt.Close()
  if err != nil {
  	return err
  }
  _, err = e_stmt.Exec(uuid.NewV4().String(), appID, tag)
  if err != nil {
  		return err
  }
  return nil
}

func TagContent(contentID string, tag string) error {
	stmt, err := db.Prepare("select id from epic.tag where tag = $1")
	if err != nil {
		return err
	}
	var tagID string
	err = stmt.QueryRow(tag).Scan(&tagID)
	if err != nil {
		return err
	}
	stmt, err = db.Prepare("insert into epic.content_tag (content_id, tag_id) values ($1, $2)")
  if err != nil {
  	return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(contentID, tagID)
  if err != nil {
  		return err
  }
	return nil
}

/*
UTILITIES
*/

func AppID(AppCode string) (uuid.UUID, error) {
	stmt, err := db.Prepare("select id from epic.application where code = $1")
	if err != nil {
		return uuid.Nil, err
	}
	var id uuid.UUID
	err = stmt.QueryRow(AppCode).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func getLetsEncryptCache(app string) (string, error) {
	var appID string
  stmt, err := db.Prepare("select application.id from epic.application where application.code = $1")
  if err != nil {
    return "", err
  }
  defer stmt.Close()
  err = stmt.QueryRow(app).Scan(&appID)
  if err != nil {
    return "", err
  }
	var data string
	stmt, err = db.Prepare("select value from epic.config where name='letsencrypt' and application_id=$1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	err = stmt.QueryRow(appID).Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func updateLetsEncryptCache(data string, app string) error {
  var appID string
  stmt, err := db.Prepare("select application.id from epic.application where application.code = $1")
  if err != nil {
    return err
  }
  defer stmt.Close()
  err = stmt.QueryRow(app).Scan(&appID)
  if err != nil {
    return err
  }
	stmt, err = db.Prepare("update epic.config set value=$1 where name='letsencrypt' and application_id=$2")
  if err != nil {
  	return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(data, appID)
  if err != nil {
  	return err
  }
  return nil
}

func now() string {
	return time.Now().Format(time.RFC3339Nano)
}
