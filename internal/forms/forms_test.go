package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("my-field", "value")
	form = New(postedData)

	has = form.Has("my-field")
	if !has {
		t.Error("shows form doesn't have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("non-existent-field", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("non-existent-field")
	if isError == "" {
		t.Error("should have an error, but did not")
	}

	postedData = url.Values{}
	postedData.Add("my-field", "value")
	form = New(postedData)

	isError = form.Errors.Get("my-field")
	if isError != "" {
		t.Error("should not have an error, but did")
	}

	postedData = url.Values{}
	postedData.Add("my-field", "value")
	form = New(postedData)

	form.MinLength("my-field", 3)
	if !form.Valid() {
		t.Error("shows min length of 3 not met when it is")
	}

	form.MinLength("my-field", 100)
	if form.Valid() {
		t.Error("form shows min length of 100 when data is shorter")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.IsEmail("non-existent-field")
	if form.Valid() {
		t.Error("form shows vaild email for non-existent field")
	}

	postedData := url.Values{}
	postedData.Add("email", "me@here.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedData = url.Values{}
	postedData.Add("email", "invalid-email")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got nvalid email when we should not have")
	}
}
