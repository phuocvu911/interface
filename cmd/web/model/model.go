package model

import "html/template"

type pageData struct {
	Mode       string
	Input      string
	Output     string
	OutputHTML template.HTML
	StatusCode int
	StatusText string
	Error      string
}
