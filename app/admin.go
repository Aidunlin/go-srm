package app

import (
	"net/url"
)

type AdminRecord struct {
	Id        int
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func NewAdminFromRegisterForm(params url.Values) (AdminRecord, []string) {
	data := AdminRecord{}
	errors := []string{}

	firstName := params.Get("first_name")
	if len(firstName) == 0 {
		errors = append(errors, "A <strong>first name</strong> is required.")
	} else {
		data.FirstName = firstName
	}

	lastName := params.Get("last_name")
	if len(lastName) == 0 {
		errors = append(errors, "A <strong>last name</strong> is required.")
	} else {
		data.LastName = lastName
	}

	email := params.Get("email")
	if len(email) > 0 {
		data.Email = email
	} else {
		errors = append(errors, "An <strong>email address</strong> is required.")
	}

	password := params.Get("password")
	if len(password) >= 8 {
		data.Password = password
	} else {
		errors = append(errors, "A <strong>password</strong> of at least <em>8 characters</em> is required.")
	}

	return data, errors
}

func NewAdminFromLoginForm(params url.Values) (AdminRecord, []string) {
	data := AdminRecord{}
	errors := []string{}

	email := params.Get("email")
	if len(email) > 0 {
		data.Email = email
	} else {
		errors = append(errors, "An <strong>email address</strong> is required.")
	}

	password := params.Get("password")
	if len(password) > 0 {
		data.Password = password
	} else {
		errors = append(errors, "A <strong>password</strong> is required.")
	}

	return data, errors
}
