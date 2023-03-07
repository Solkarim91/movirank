package graph

import "github.com/solkarim91/movirank/api/app/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MovieRepository repository.MovieRepository
	RatingRepository repository.RatingRepository
}
