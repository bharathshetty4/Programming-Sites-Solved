package main

// mux router to authenticate basic auth [base64 encoded] using middleware

/*
curl request:

curl --request GET \
  --url http://localhost:9900/form \
  --header 'authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ='

*/
import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/form", use(myHandler, basicAuth))
	http.Handle("/", r)

	http.ListenAndServe(":9900", nil)
}

// use provides a cleaner interface for chaining middleware for single routes.
// Middleware functions are simple HTTP handlers (w http.ResponseWriter, r *http.Request)
//
//  r.HandleFunc("/login", use(loginHandler, rateLimit, csrf))
//  r.HandleFunc("/form", use(formHandler, csrf))
//  r.HandleFunc("/about", aboutHandler)
//

func use(h http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticated!"))
	return
}

// Leverages nemo's answer in http://stackoverflow.com/a/21937924/556573
func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		if pair[0] != "username" || pair[1] != "password" {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	}
}
