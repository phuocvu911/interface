package main

import (
	u "art-decoder/utils"
	"embed"
	"html/template"
	"net/http"
	"strings"
)

//go:embed assets/templates/index.html assets/static/styles.css
var webFS embed.FS
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

	render(w, tpl, http.StatusOK, pageData{
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
}
