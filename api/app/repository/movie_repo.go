package repository

import (
	"github.com/google/uuid"
	"github.com/solkarim91/movirank/api/app/models"
	"github.com/solkarim91/movirank/api/graph/model"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movieInput *model.MovieInput) (*models.Movie, error)
	UpdateMovie(movieInput *model.MovieInput, id string) error
	DeleteMovie(id string) error
	GetMovie(id string) (*models.Movie, error)
	GetMovies() ([]*model.Movie, error)
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

func (b *MovieService) CreateMovie(movieInput *model.MovieInput) (*models.Movie, error) {
	movie := &models.Movie{
		ID: 					uuid.New().String(),
		Title:     		movieInput.Title,
		Description:  movieInput.Description,
		Director: 	  movieInput.Director,
		Genre:     		movieInput.Genre,
		Runtime:      movieInput.Runtime,
		Released:     movieInput.Released,
	}
	err := b.Db.Create(&movie).Error

	return movie, err
}

func (b *MovieService) UpdateMovie(movieInput *model.MovieInput, id string) error {
	movie := models.Movie{
		ID:        		id,
		Title:     		movieInput.Title,
		Description:  movieInput.Description,
		Director: 	  movieInput.Director,
		Genre:    	  movieInput.Genre,
		Runtime:      movieInput.Runtime,
		Released:     movieInput.Released,
	}
	err := b.Db.Model(&movie).Where("id = ?", id).Updates(movie).Error
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

func (b *MovieService) GetMovies() ([]*model.Movie, error) {
	movies := []*model.Movie{}
	err := b.Db.Find(&movies).Error
	return movies, err

}