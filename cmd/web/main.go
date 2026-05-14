package main

import (
	"embed"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"
)

//go:embed assets/templates/index.html assets/static/styles.css
var webFS embed.FS

var tpl = template.Must(template.ParseFS(webFS, "assets/templates/index.html"))

func main() {
	flag.Parse()
	assetsRootFS, err := fs.Sub(webFS, "assets")
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	// serve styles.css so the browser can load it when the HTML references /static/styles.css.
	mux.Handle("/static/", http.FileServerFS(assetsRootFS))

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
