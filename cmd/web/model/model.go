package model

// field names must be exported (start with uppercase) to be accessible in the template.
// f.e: Mode can be  accessed in the template as {{.Mode}}
type PageData struct {
	Mode       string // "encode" or "decode"
	Input      string // user input from the form
	Output     string // the result of encoding or decoding the input
	StatusCode int    // HTTP status code to indicate success or type of error
	StatusText string // human-readable text of StatusCode
	Error      string // detailed error message for display in the UI if something went wrong
}
