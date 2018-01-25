package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		res := struct{ Status string }{"ok"}
		jData, _ := json.Marshal(&res)

		w.WriteHeader(http.StatusOK)
		w.Write(jData)
	})

	http.HandleFunc("/ajs-proxy.min.js", func(w http.ResponseWriter, r *http.Request) {
		snippet, err := ioutil.ReadFile("static/ajs-proxy.min.js")
		os.Stdout.Write(snippet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(http.StatusOK)
		w.Write(snippet)
	})

	u, _ := url.Parse("https://api.segment.io/")
	http.Handle("/", httputil.NewSingleHostReverseProxy(u))

	// Start the server
	log.Printf("Initializing Proxy")
	log.Fatal(http.ListenAndServe(":80", nil))
}
