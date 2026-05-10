package main

import (
	u "art-decoder/utils"
	m "art-interface/cmd/web/model"
	"bufio"
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//go:embed assets/templates/index.html assets/static/styles.css
var webFS embed.FS

// create new tamplate name "index.html" and parse the index.html file from the embedded
// filesystem webFS. The template.Must will panic if there is an error parsing the template,
// which is appropriate here because we want to catch any errors in our templates at startup rather than at runtime when handling requests.
var tpl = template.Must(template.New("index.html").ParseFS(webFS, "assets/templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		//w.WriteHeader(http.StatusMethodNotAllowed)  //silent
		w.Header().Set("Allow", http.MethodGet)                                                  //respond with Allow GET only
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed) //and print code 405: method not allow.
		return
	}

	//initial page load with empty form and status 200 OK. Default mode is decode.
	render(w, tpl, http.StatusOK, m.PageData{
		Mode:       "decode",
		StatusCode: http.StatusOK,
		StatusText: http.StatusText(http.StatusOK),
	})
}

func decoderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		render(w, tpl, http.StatusBadRequest, m.PageData{
			Mode:       "decode",
			StatusCode: http.StatusBadRequest,
			StatusText: http.StatusText(http.StatusBadRequest),
			Error:      "Malformed form data.",
		})
		return
	}

	mode := r.PostFormValue("mode")
	input := r.PostFormValue("data")

	//normal case would never reach this block because of radio button.
	//This is for protect the switch below as someone can send a POST request with curl or
	//Postman without the mode field or with an invalid mode value. We want to catch that and
	//respond with a 400 Bad Request and an error message in the UI rather than processing it and
	//potentially returning a 500 Internal Server Error or some other unexpected result.
	if mode != "decode" && mode != "encode" {
		render(w, tpl, http.StatusBadRequest, m.PageData{
			Mode:       "decode",
			Input:      input,
			StatusCode: http.StatusBadRequest,
			StatusText: http.StatusText(http.StatusBadRequest),
			Error:      "Malformed query: missing or invalid mode.",
		})
		return
	}

	var out string
	status := http.StatusAccepted
	switch mode {
	case "encode":
		// Empty input is valid for encode (it encodes to empty output).
		out = processLinesEncode(input)
	case "decode":
		if input == "" {
			render(w, tpl, http.StatusBadRequest, m.PageData{
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
		out = decodedText
	}

	render(w, tpl, status, m.PageData{
		Mode:       mode,
		Input:      input,
		Output:     out,
		StatusCode: status,
		StatusText: http.StatusText(status),
	})
}

func render(w http.ResponseWriter, tpl *template.Template, status int, data m.PageData) {
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
