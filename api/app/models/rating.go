package models

import "time"

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

type UpdateRatingInput struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Star    int    `json:"star"`
}