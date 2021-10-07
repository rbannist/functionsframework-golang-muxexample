package muxservice

import (
	"net/http"
)

var mux = newMux()

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Default)
	mux.HandleFunc("/one", One)
	mux.HandleFunc("/two", Two)

	return mux
}

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

func Default(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go to '/one' or '/two' for mux routing example"))
}

func One(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello one"))
}

func Two(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello two"))
}
