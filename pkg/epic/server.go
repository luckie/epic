package epic

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/justinas/alice"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func serveHTTP(httpPort int, host string, dbUser string, dbPassword string, dbServer string) {

	dbConnStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbServer + "/epic?sslmode=disable"
	initSQLDatabase(dbConnStr)

	r := NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "Accept"},
	})
	chain := alice.New(c.Handler, AuthHandler).Then(r)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(httpPort),
		Handler: chain,
	}

	fmt.Println("Starting HTTP Server for hosting on port " + strconv.Itoa(httpPort) + ".")
	log.Fatal(server.ListenAndServe())
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