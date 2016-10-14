package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	//"os"
	"time"
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

type Content struct {
	ID          uuid.UUID `json:"id, omitempty"`
	AppID       uuid.UUID `json:"app-id, omitempty"`
	Name        string    `json:"name, omitempty"`
	Description string    `json:"description, omitempty"`
	Error       error     `json:"error, omitempty"`
}

type Entry struct {
	ID        uuid.UUID `json:"-"`
	ContentID uuid.UUID `json:"id"`
	Locale    string    `json:"locale"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

type Err struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type ID struct {
	ID string `json:"id"`
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

func CreateContentReservation(c *Content) error {
	c.ID = uuid.NewV4()
	timestamp := time.Now()
	e_stmt, err := db.Prepare("insert into epic.content (id, application_id, name, description, timestamp) values ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer e_stmt.Close()
	_, err = e_stmt.Exec(c.ID.String(), c.AppID.String(), c.Name, c.Description, timestamp)
	if err != nil {
		return err
	}
	return nil
}

func CreateEntryForContentID(e *Entry) error {
	locID, err := localeID(e.Locale)
	if err != nil {
		return err
	}
	e_stmt, err := db.Prepare("insert into epic.entry (id, content_id, locale_id, timestamp, data) values ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer e_stmt.Close()
	_, err = e_stmt.Exec(e.ID.String(), e.ContentID.String(), locID.String(), e.Timestamp, e.Data)
	if err != nil {
		return err
	}
	return nil
}

func AllContentForTag(Tag string, AppID string) ([]Entry, error) {
	stmt, err := db.Prepare("select entry.id, entry.content_id, locale.code as locale, entry.timestamp, entry.data from epic.entry inner join epic.locale on entry.locale_id=locale.id inner join epic.content on entry.content_id=content.id inner join epic.content_tag on content.id=content_tag.content_id inner join epic.tag on tag.id=content_tag.tag_id where tag.value=$1 and tag.application_id=$2")
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

func NewestLocalizedContentEntriesForTag(Tag string, Locale string, AppID string) ([]Entry, error) {
	sql := `
	with c as (
	select distinct content.id from epic.content
	inner join epic.content_tag on content.id=content_tag.content_id
	inner join epic.tag on tag.id=content_tag.tag_id
	where tag.value = $1 and content_tag.application_id = $3
	)
	select id, content_id, locale, timestamp, data from
	(
	select entry.id, entry.content_id, locale.code as locale, entry.timestamp, entry.data,
	rank() over (partition by entry.content_id order by entry.timestamp desc ) as rank
	from epic.entry
	inner join epic.locale on entry.locale_id=locale.id
	inner join c on entry.content_id=c.id
	group by c.id, entry.timestamp, entry.id, locale.code
	order by c.id, entry.timestamp, entry.id, locale
	) e
	where e.rank = 1 and e.locale = $2
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(Tag, Locale, AppID)
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

func TagContent(contentID string, tag string, AppID string) error {
	stmt, err := db.Prepare("select id from epic.tag where value = $1")
	if err != nil {
		return errors.New("Prepare select from tag | " + err.Error())
	}
	var tagID string
	err = stmt.QueryRow(tag).Scan(&tagID)
	if err != nil {
		return errors.New("QueryRow / Scan for tag | " + err.Error())
	}
	stmt, err = db.Prepare("insert into epic.content_tag (content_id, tag_id, application_id) values ($1, $2, $3)")
	if err != nil {
		return errors.New("Prepare insert into content_tag | " + err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(contentID, tagID, AppID)
	if err != nil {
		return errors.New("Exec insert into content_tag | " + err.Error())
	}
	return nil
}

/*
UTILITIES
*/

func AppID(appCode string) (uuid.UUID, error) {
	stmt, err := db.Prepare("select id from epic.application where code = $1")
	if err != nil {
		return uuid.Nil, err
	}
	var id uuid.UUID
	err = stmt.QueryRow(appCode).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func localeID(localeCode string) (uuid.UUID, error) {
	stmt, err := db.Prepare("select locale.id from epic.locale where locale.code = $1")
	if err != nil {
		return uuid.Nil, err
	}
	var id uuid.UUID
	err = stmt.QueryRow(localeCode).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// To be implemented ASAP.
func tagUnique(tag string) error {
	return nil
}

// DOES THIS NEED UPDATING?
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

// DOES THIS NEED UPDATING?
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
