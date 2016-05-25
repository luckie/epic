package main

import (
	"net/http"
	"crypto/tls"
	"net"
	"rsc.io/letsencrypt"
	//"github.com/gorilla/mux"
	"log"
	"time"
	"fmt"
  "strconv"

  "github.com/justinas/alice"
  // "github.com/rs/cors"
)

//func initServer(url string, port int) {
//  r := NewRouter()
//  if port != 443 {
//    fmt.Println("Starting HTTP Server on port " + strconv.Itoa(port) + ".")
//    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
//  } else {
//    fmt.Println("Starting HTTPS Server on port 443 with a Let's Encrypt TLS certificate.")
//    var m letsencrypt.Manager
//    if err := m.CacheFile("letsencrypt.cache"); err != nil {
//      log.Fatal(err)
//    }
//    log.Fatal(Serve(&m, r))
//  }
//}

//func Serve(m *letsencrypt.Manager, r *mux.Router) error {
//
//  l, err := net.Listen("tcp", ":http")
//	if err != nil {
//		return err
//	}
//	defer l.Close()
//	go http.Serve(l, http.HandlerFunc(redirectHTTP))

//	return serveHTTPS(m, r)
//}

func Serve(url string, port int) error {
  r := NewRouter()
  chain := alice.New(AuthHandler).Then(r)

  if port != 443 {
    fmt.Println("Starting HTTP Server on port " + strconv.Itoa(port) + ".")
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
  } else {
    fmt.Println("Starting HTTPS Server on port 443 with a Let's Encrypt TLS certificate.")
    var m letsencrypt.Manager
    //if err := m.CacheFile("letsencrypt.cache"); err != nil {
    if err := CacheFile(&m); err != nil {
			fmt.Println(err)
      log.Fatal(err)
    }

    l, err := net.Listen("tcp", ":http")
    if err != nil {
			fmt.Println(err)
    	return err
    }
    defer l.Close()
    go http.Serve(l, http.HandlerFunc(redirectHTTP))

    //return serveHTTPS(&m, &r)
    log.Fatal(serveHTTPS(&m, chain))

  }
  return nil

}

//func serveHTTPS(m *letsencrypt.Manager, r *mux.Router) error {
func serveHTTPS(m *letsencrypt.Manager, chain http.Handler) error {
  srv := &http.Server{
		Addr:    ":https",
		//Handler: r,
    Handler: chain,
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}
	return srv.ListenAndServeTLS("", "")
}

// RedirectHTTP is an HTTP handler (suitable for use with http.HandleFunc)
// that responds to all requests by redirecting to the same URL served over HTTPS.
// It should only be invoked for requests received over HTTP.
func redirectHTTP(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil || r.Host == "" {
		http.Error(w, "not found", 404)
	}

	u := r.URL
	u.Host = r.Host
	u.Scheme = "https"
	http.Redirect(w, r, u.String(), 302)
}

func CacheFile(m *letsencrypt.Manager) error {
	data, err := getLetsEncryptCache("schutt")
	if err != nil {
		return err
	}
	if len(data) > 0 {
		if err := m.Unmarshal(data); err != nil {
			return err
		}
	}
	go func() {
		for range m.Watch() {
			//err := ioutil.WriteFile(name, []byte(m.Marshal()), 0600)
      err := updateLetsEncryptCache(m.Marshal(), "schutt")
			if err != nil {
				log.Printf("writing letsencrypt cache: %v", err)
        fmt.Println("data marshal error: " + err.Error())
			}
		}
	}()
	return nil
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
