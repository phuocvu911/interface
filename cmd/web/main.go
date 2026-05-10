package main

import (
	"bufio"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"

	u "art-decoder/utils"
)

type pageData struct {
	Mode       string
	Input      string
	Output     string
	OutputHTML template.HTML
	StatusCode int
	StatusText string
	Error      string
}

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

func render(w http.ResponseWriter, tpl *template.Template, status int, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := tpl.Execute(w, data); err != nil {
		log.Printf("Error rendering template: %v", err)
	}
}

func processLinesEncode(input string) string {
	var outLines []string

	sc := bufio.NewScanner(strings.NewReader(input))
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)

	for sc.Scan() {
		line := sc.Text()
		outLines = append(outLines, u.Encode(line))
	}

	return strings.Join(outLines, "\n")
}

// processLinesDecode decodes each line independently (like the CLI --multi mode).
// For any line that fails to decode, it outputs the literal "Error" for that line.
func processLinesDecode(input string) (string, bool) {
	var outLines []string
	hadErr := false

	sc := bufio.NewScanner(strings.NewReader(input))
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)

	for sc.Scan() {
		line := sc.Text()
		decoded, err := u.Decode(line)
		if err != nil {
			hadErr = true
			outLines = append(outLines, "Error")
			continue
		}
		outLines = append(outLines, decoded)
	}

	return strings.Join(outLines, "\n"), hadErr
}
