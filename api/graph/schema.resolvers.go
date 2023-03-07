package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/solkarim91/movirank/api/graph/model"
)

// CreateMovie is the resolver for the CreateMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.CreateMovieInput) (*model.Movie, error) {
	movie, err := r.MovieRepository.CreateMovie(&input)
	movieCreated := &model.Movie{
		Title:       movie.Title,
		Description: movie.Description,
		Director:    movie.Director,
		Cast:        movie.Cast,
		Genre:       movie.Genre,
		Runtime:     movie.Runtime,
		Released:    movie.Released,
		Img:         movie.Img,
		ID:          movie.ID,
	}
	if err != nil {
		return nil, err
	}
	return movieCreated, nil
}

// DeleteMovie is the resolver for the DeleteMovie field.
func (r *mutationResolver) DeleteMovie(ctx context.Context, id string) (string, error) {
	err := r.MovieRepository.DeleteMovie(id)
	if err != nil {
		return "", err
	}
	successMessage := "successfully deleted"
	return successMessage, nil
}

// UpdateMovie is the resolver for the UpdateMovie field.
func (r *mutationResolver) UpdateMovie(ctx context.Context, input model.UpdateMovieInput) (string, error) {
	err := r.MovieRepository.UpdateMovie(&input)
	if err != nil {
		return "", err
	}
	successMessage := "successfully updated"
	return successMessage, nil
}

// CreateRating is the resolver for the createRating field.
func (r *mutationResolver) CreateRating(ctx context.Context, input model.CreateRatingInput) (*model.Rating, error) {
	rating, err := r.MovieRepository.CreateRating(&input)
	ratingCreated := &model.Rating{
		ID:        rating.ID,
		MovieID:   rating.MovieID,
		CreatedAt: rating.CreatedAt,
		UpdatedAt: rating.UpdatedAt,
		Message:   rating.Message,
		Star:      rating.Star,
	}

	if err != nil {
		return nil, err
	}

	return ratingCreated, nil
}

// GetMovies is the resolver for the getMovies field.
func (r *queryResolver) GetMovies(ctx context.Context) ([]*model.Movie, error) {
	movies, err := r.MovieRepository.GetMovies()
	var allMovies []*model.Movie

	for _, movie := range movies {
		mov := &model.Movie{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Director:    movie.Director,
			Cast:        movie.Cast,
			Genre:       movie.Genre,
			Runtime:     movie.Runtime,
			Released:    movie.Released,
			Img:         movie.Img,
		}
		allMovies = append(allMovies, mov)
	}

	if err != nil {
		return nil, err
	}
	return allMovies, nil
}

// GetOneMovie is the resolver for the getOneMovie field.
func (r *queryResolver) GetOneMovie(ctx context.Context, id string) (*model.Movie, error) {
	movie, err := r.MovieRepository.GetMovie(id)
	selectedMovie := &model.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Director:    movie.Director,
		Cast:        movie.Cast,
		Genre:       movie.Genre,
		Runtime:     movie.Runtime,
		Released:    movie.Released,
		Img:         movie.Img,
	}
	if err != nil {
		return nil, err
	}
	return selectedMovie, nil
}

// GetRatingsByMovieID is the resolver for the getRatingsByMovieId field.
func (r *queryResolver) GetRatingsByMovieID(ctx context.Context, movieID string) ([]*model.Rating, error) {
	ratings, err := r.MovieRepository.GetRatingsByMovieId(movieID)
	var ratingsByMovieId []*model.Rating

	for _, rating := range ratings {
		rat := &model.Rating{
			ID:          rating.ID,
			MovieID: 		 rating.MovieID,
			CreatedAt: rating.CreatedAt,
			UpdatedAt: rating.UpdatedAt,
			Message: rating.Message,
			Star: rating.Star,
		}
		ratingsByMovieId = append(ratingsByMovieId, rat)
	}

	if err != nil {
		return nil, err
	}
	return ratingsByMovieId, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
