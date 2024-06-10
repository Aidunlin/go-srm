package model

import (
	"fmt"
	"net/url"

	"github.com/go-mysql-org/go-mysql/mysql"
	"golang.org/x/crypto/bcrypt"
)

type AdminRecord struct {
	Id        int64
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func (s AdminRecord) ToMap() map[string]string {
	return map[string]string{
		"id":         fmt.Sprint(s.Id),
		"email":      s.Email,
		"password":   s.Password,
		"first_name": s.FirstName,
		"last_name":  s.LastName,
	}
}

func NewAdminRecordFromRegisterForm(params url.Values) (AdminRecord, []string) {
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
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			errors = append(errors, "Could not hash your <strong>password</strong>!")
		} else {
			data.Password = string(hashed)
		}
	} else {
		errors = append(errors, "A <strong>password</strong> of at least <em>8 characters</em> is required.")
	}

	return data, errors
}

func NewAdminRecordFromLoginForm(params url.Values) (AdminRecord, []string) {
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
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			errors = append(errors, "Could not hash your <strong>password</strong>!")
		} else {
			data.Password = string(hashed)
		}
	} else {
		errors = append(errors, "A <strong>password</strong> is required.")
	}

	return data, errors
}

func GetAdminRecordFromResult(result *mysql.Result, row int) AdminRecord {
	id, _ := result.GetIntByName(row, "id")
	email, _ := result.GetStringByName(row, "email")
	password, _ := result.GetStringByName(row, "password")
	firstName, _ := result.GetStringByName(row, "first_name")
	lastName, _ := result.GetStringByName(row, "last_name")

	return AdminRecord{
		Id:        id,
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
}
