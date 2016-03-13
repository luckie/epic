package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db gorm.DB

func main() {
	startDatabase()
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func startDatabase() {
	db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Content struct {
	ID        uuid.UUID       `json:"-"`
	ContentID uuid.UUID       `json:"id"`
	Locale    string          `json:"locale"`
	Timestamp time.Time       `json:"timestamp"`
	Value     json.RawMessage `json:"value"`
}

type Err struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ContentIndex",
		"GET",
		"/content",
		ContentIndex,
	},
	Route{
		"ContentCreate",
		"POST",
		"/content",
		ContentCreate,
	},
	Route{
		"ContentRead",
		"GET",
		"/content/{id}",
		ContentRead,
	},
	Route{
		"ContentUpdate",
		"PUT",
		"/content/{id}",
		ContentUpdate,
	},
	Route{
		"ContentDelete",
		"DELETE",
		"/content/{id}",
		ContentDelete,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func ContentIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func ContentCreate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	id := uuid.NewV4()
	contentID := uuid.NewV4()
	c := Content{ID: id, ContentID: contentID, Value: body, Timestamp: time.Now()}

	if err := json.Unmarshal(body, &c); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db.NewRecord(c)
	db.Create(&c)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}

}

func ContentRead(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	id := uuid.NewV4()
	contentID := uuid.NewV4()
	c := Content{ID: id, ContentID: contentID, Value: body, Timestamp: time.Now()}

	if err := json.Unmarshal(body, &c); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db.NewRecord(c)
	db.Create(&c)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}

	fmt.Fprint(w, "To be implemented.\n")
}

func ContentUpdate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	id := uuid.NewV4()
	contentID := uuid.NewV4()
	c := Content{ID: id, ContentID: contentID, Value: body, Timestamp: time.Now()}

	if err := json.Unmarshal(body, &c); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db.NewRecord(c)
	db.Create(&c)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}

	fmt.Fprint(w, "To be implemented.\n")
}

func ContentDelete(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	id := uuid.NewV4()
	contentID := uuid.NewV4()
	c := Content{ID: id, ContentID: contentID, Value: body, Timestamp: time.Now()}

	if err := json.Unmarshal(body, &c); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db.NewRecord(c)
	db.Create(&c)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}

	fmt.Fprint(w, "To be implemented.\n")
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
