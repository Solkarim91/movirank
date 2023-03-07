package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/solkarim91/movirank/api/app/models"
	"github.com/solkarim91/movirank/api/graph/model"
	"gorm.io/gorm"
)

type RatingRepository interface {
	CreateRating(ratingInput *model.CreateRatingInput) (*models.Rating, error)
	UpdateRating(ratingInput *model.UpdateRatingInput) error
	DeleteRating(id string) error
	GetRatingsByMovieId(movieId string) ([]*models.Rating, error)
}

type RatingService struct {
	Db *gorm.DB
}

var _ RatingRepository = &RatingService{}

func NewRatingService(db *gorm.DB) *RatingService {
	return &RatingService{
		Db: db,
	}
}

func (b *RatingService) CreateRating(ratingInput *model.CreateRatingInput) (*models.Rating, error) {
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

func (b *RatingService) UpdateRating(ratingInput *model.UpdateRatingInput) error {
	rating := models.Rating{
		ID:        		ratingInput.ID,
		Message:      ratingInput.Message,
		Star:       	ratingInput.Star,
	}
	err := b.Db.Model(&rating).Where("id = ?", ratingInput.ID).Updates(rating).Error
	return err
}

func (b *RatingService) DeleteRating(id string) error {
	rating := &models.Rating{}
	err := b.Db.Delete(rating, "id = ?", id).Error
	return err
}

func (b *RatingService) GetRatingsByMovieId(movieId string) ([]*models.Rating, error) {
	ratings := []*models.Rating{}
	err := b.Db.Find(&ratings, "movie_id = ?", movieId).Error
	return ratings, err
}