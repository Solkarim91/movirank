package models

import (
	"time"

	"github.com/lib/pq"
)

type Movie struct {
	ID    		  string          `gorm:"primary key" json:"id"`
	Title     	string 		    	`json:"title"`
	Description string     			`json:"description"`
	Director 		string 		    	`json:"director"`
	Cast        pq.StringArray  `gorm:"type:text[]" json:"cast"`
	Genre 			string 	    		`json:"genre"`
	Runtime 		float64	    		`json:"runtime"`
	Released 		int     				`json:"released"`
	Img 				string    			`json:"img"`
	Ratings     []*Rating 			`json:"ratings"`
}

type MovieInput struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Director    string          `json:"director"`
	Cast        pq.StringArray  `gorm:"type:text[]" json:"cast"`
	Genre       string          `json:"genre"`
	Runtime     float64         `json:"runtime"`
	Released    int             `json:"released"`
	Img         string          `json:"img"`
}

type UpdateMovieInput struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Director    string         `json:"director"`
	Cast        pq.StringArray `gorm:"type:text[]" json:"cast"`
	Genre       string         `json:"genre"`
	Runtime     float64        `json:"runtime"`
	Released    int            `json:"released"`
	Img         string         `json:"img"`
}

type Rating struct {
	ID        string    `json:"id"`
	MovieID   string    `json:"movieId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Message   string    `json:"message"`
	Star      int       `json:"star"`
}

type CreateRatingInput struct {
	MovieID string `json:"movieId"`
	Message string `json:"message"`
	Star    int    `json:"star"`
}