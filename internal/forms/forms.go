package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in POST request and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	// better way but not straightforward
	return r.Form.Get(field) != ""
	// begginner way but straigtforward
	// if r.Form.Get(field) == "" {
	// 	return false
	// }
	// return true
}
