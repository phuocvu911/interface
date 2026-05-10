package main

import (
	m "art-interface/cmd/web/model"
	"embed"
	"html/template"
	"net/http"
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
