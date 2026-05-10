package main

import (
	"flag"
	"io/fs"
	"log"
	"net/http"
	"time"
)

func main() {
	flag.Parse()

	// make public URL shorter, hiding developer detail of file structure and also to prevent directory traversal attack (e.g. /static/../../secret.txt).
	staticFS, err := fs.Sub(webFS, "assets/static")
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	// serve style.css so the browser can load it when the HTML references /static/styles.css. The http.FileServerFS will serve files from the staticFS, which is rooted at assets/static, so it can only serve files that are in that directory (e.g. styles.css) and cannot access files outside of it (e.g. assets/templates/index.html or any other file on the system).
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServerFS(staticFS)))

	//hooks up handlers to endpoints
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/decoder", decoderHandler)

	srv := &http.Server{
		Addr:         *addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("art-interface web listening on http://localhost%s", *addr)
	log.Fatal(srv.ListenAndServe())
}
