package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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
	if r.Form.Get(field) == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		val := f.Get(field)
		if strings.TrimSpace(val) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks for string minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	if len(r.Form.Get(field)) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
