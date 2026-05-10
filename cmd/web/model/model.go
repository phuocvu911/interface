package model

type pageData struct {
	Mode       string // "encode" or "decode"
	Input      string // user input from the form
	Output     string // the result of encoding or decoding the input
	StatusCode int    // HTTP status code to indicate success or type of error
	StatusText string // human-readable text of StatusCode
	Error      string // detailed error message for display in the UI if something went wrong
}
