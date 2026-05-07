package main

import (
	"bufio"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"

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

	staticFS, err := fs.Sub(webFS, "assets/static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	//hooks up handlers to endpoints
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/decoder", decoderHandler)

	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	log.Printf("art-decoder web listening on http://localhost%s", *addr)
	log.Fatal(srv.ListenAndServe())
}

func render(w http.ResponseWriter, tpl *template.Template, status int, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_ = tpl.Execute(w, data)
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
