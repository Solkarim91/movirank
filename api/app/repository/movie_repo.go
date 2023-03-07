package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/solkarim91/movirank/api/app/models"
	"github.com/solkarim91/movirank/api/graph/model"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movieInput *model.CreateMovieInput) (*models.Movie, error)
	UpdateMovie(movieInput *model.UpdateMovieInput) error
	DeleteMovie(id string) error
	GetMovie(id string) (*models.Movie, error)
	GetMovies() ([]*models.Movie, error)
	CreateRating(ratingInput *model.CreateRatingInput) (*models.Rating, error)
	UpdateRating(ratingInput *model.UpdateRatingInput) error
	DeleteRating(id string) error
	GetRatingsByMovieId(movieId string) ([]*models.Rating, error)
}

type MovieService struct {
	Db *gorm.DB
}

var _ MovieRepository = &MovieService{}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{
		Db: db,
	}
}

func (b *MovieService) CreateMovie(movieInput *model.CreateMovieInput) (*models.Movie, error) {
	movie := &models.Movie{
		ID: 					uuid.New().String(),
		Title:     		movieInput.Title,
		Description:  movieInput.Description,
		Director: 	  movieInput.Director,
		Cast: 			  movieInput.Cast,
		Genre:     		movieInput.Genre,
		Runtime:      movieInput.Runtime,
		Released:     movieInput.Released,
		Img: 					movieInput.Img,		
	}
	err := b.Db.Create(&movie).Error

	return movie, err
}

func (b *MovieService) UpdateMovie(movieInput *model.UpdateMovieInput) error {
	movie := models.Movie{
		ID:        		movieInput.ID,
		Title:     		movieInput.Title,
		Description:  movieInput.Description,
		Director: 	  movieInput.Director,
		Cast: 			  movieInput.Cast,
		Genre:    	  movieInput.Genre,
		Runtime:      movieInput.Runtime,
		Released:     movieInput.Released,
		Img: 					movieInput.Img,
	}
	err := b.Db.Model(&movie).Where("id = ?", movieInput.ID).Updates(movie).Error
	return err
}

func (b *MovieService) DeleteMovie(id string) error {
	movie := &models.Movie{}
	err := b.Db.Delete(movie, "id = ?", id).Error
	return err
}

func (b *MovieService) GetMovie(id string) (*models.Movie, error) {
	movie := &models.Movie{}
	err := b.Db.Where("id = ?", id).First(movie).Error
	return movie, err
}

func (b *MovieService) GetMovies() ([]*models.Movie, error) {
	movies := []*models.Movie{}
	err := b.Db.Find(&movies).Error
	return movies, err
}

func (b *MovieService) CreateRating(ratingInput *model.CreateRatingInput) (*models.Rating, error) {
	movie := &models.Movie{}
	err := b.Db.Where("id = ?", ratingInput.MovieID).First(movie).Error
	now := time.Now().UTC()

	rating := models.Rating{
		ID:        uuid.New().String(),
		CreatedAt: now,
		UpdatedAt: now,
		Message:   ratingInput.Message,
		Star:			 ratingInput.Star,
	}

	movie.Ratings = append(movie.Ratings, &rating)

	err2 := b.Db.Model(movie).Where("id = ?", ratingInput.MovieID).Updates(movie).Error

	if err != nil {
		return nil, err
	}

	if err2 != nil {
		return nil, err2
	}

	return &rating, nil
}

func (b *MovieService) UpdateRating(ratingInput *model.UpdateRatingInput) error {
	rating := models.Rating{
		ID:        		ratingInput.ID,
		Message:      ratingInput.Message,
		Star:       	ratingInput.Star,
	}
	err := b.Db.Model(&rating).Where("id = ?", ratingInput.ID).Updates(rating).Error
	return err
}

func (b *MovieService) DeleteRating(id string) error {
	rating := &models.Rating{}
	err := b.Db.Delete(rating, "id = ?", id).Error
	return err
}

func (b *MovieService) GetRatingsByMovieId(movieId string) ([]*models.Rating, error) {
	ratings := []*models.Rating{}
	err := b.Db.Find(&ratings, "movie_id = ?", movieId).Error
	return ratings, err
}