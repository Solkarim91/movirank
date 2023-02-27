package models

import (
	"fmt"
	"io"
	"strconv"
)

type Movie struct {
    ID    		  string    	`gorm:"primary key;autoIncrement" json:"id"`
    Title     	string 			`json:"title"`
    Description string 			`json:"description"`
    Genre 			string 			`json:"genre"`
		Runtime 		float64			`json:"runtime"`
		Released 		int 				`json:"released"`
		Img 				string 			`json:"img"`
		// Ratings     []Rating  	`json:"ratings"`
		// Comments    []Comment 	`json:"comments"`
}

type Comment struct {
		ID 		string `json:"id"`
		User 	*User  `json:"user"`
		Text 	string `json:"text"`
}

type Rating struct {
		ID   	string  `json:"id"`
		User	*User 	`json:"user"`
		Star 	int   	`json:"star"`
}

type User struct {
		ID 				string   `json:"id"`
		UserType  UserType `json:"userType"`
		Username 	string   `json:"username"`
		Email     string   `json:"email"`
}

type UserType string

const (
	UserTypeAdmin    UserType = "ADMIN"
	UserTypeStandard UserType = "STANDARD"
)

var AllUserType = []UserType{
	UserTypeAdmin,
	UserTypeStandard,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeAdmin, UserTypeStandard:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
