package repository

import (
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
