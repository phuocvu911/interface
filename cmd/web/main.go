package main

import (
	"bufio"
	"embed"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"

	u "art-decoder/utils"
)

var (
	addr  = flag.String("addr", ":8080", "server listen address")
	paint = flag.Bool("paint", false, "colorize decoded output (bonus)")
)

//go:embed assets/templates/index.html assets/static/styles.css
var webFS embed.FS

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

	tpl := template.Must(template.New("index.html").ParseFS(webFS, "assets/templates/index.html"))

	staticFS, err := fs.Sub(webFS, "assets/static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if r.Method != http.MethodGet {
			//w.WriteHeader(http.StatusMethodNotAllowed)  //silent
			w.Header().Set("Allow", http.MethodGet) //respond with Allow GET only
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed) //and print code 405: method not allow.
			return
		}

		render(w, tpl, http.StatusOK, pageData{
			Mode:       "decode",
			StatusCode: http.StatusOK,
			StatusText: http.StatusText(http.StatusOK),
		})
	})

	mux.HandleFunc("/decoder", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			render(w, tpl, http.StatusBadRequest, pageData{
				Mode:       "decode",
				StatusCode: http.StatusBadRequest,
				StatusText: http.StatusText(http.StatusBadRequest),
				Error:      "Malformed form data.",
			})
			return
		}

		mode := strings.ToLower(strings.TrimSpace(r.PostFormValue("mode")))
		input := r.PostFormValue("data")

		if mode != "decode" && mode != "encode" {
			render(w, tpl, http.StatusBadRequest, pageData{
				Mode:       "decode",
				Input:      input,
				StatusCode: http.StatusBadRequest,
				StatusText: http.StatusText(http.StatusBadRequest),
				Error:      "Malformed query: missing or invalid mode.",
			})
			return
		}

		var out string
		var outHTML template.HTML
		status := http.StatusAccepted
		switch mode {
		case "encode":
			// Empty input is valid for encode (it encodes to empty output).
			out = processLinesEncode(input)
		case "decode":
			if input == "" {
				render(w, tpl, http.StatusBadRequest, pageData{
					Mode:       mode,
					StatusCode: http.StatusBadRequest,
					StatusText: http.StatusText(http.StatusBadRequest),
					Error:      "Malformed query: missing input.",
				})
				return
			}
			decodedText, hadErr := processLinesDecode(input)
			if hadErr {
				status = http.StatusBadRequest
			}
			if *paint && !hadErr {
				outHTML = template.HTML(u.PaintLineHTML(decodedText))
			} else {
				out = decodedText
			}
		}

		render(w, tpl, status, pageData{
			Mode:       mode,
			Input:      input,
			Output:     out,
			OutputHTML: outHTML,
			StatusCode: status,
			StatusText: http.StatusText(status),
		})
	})

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
