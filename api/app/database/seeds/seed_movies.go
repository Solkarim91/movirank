package seeds

import (
	"github.com/google/uuid"
	"github.com/solkarim91/movirank/api/app/models"
	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB, id string, title string, description string, director string, genre string, runtime float64, released int, img string) error {
        return db.Create(&models.Movie{ID: id, Title: title, Description: description, Director: director, Genre: genre, Runtime: runtime, Released: released, Img: img,}).Error
}

func SeedMovies() []Seed {
	return (
		[]Seed{
				{
					Name: "CreateMovie1",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "The Lord of the Rings: The Return of the King", "Gandalf and Aragorn lead the World of Men against Sauron's army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.", "Peter Jackson", "Fantasy", 3.21, 2003, "https://res.cloudinary.com/domtywpyc/image/upload/v1676474113/LOTR-ROTK_qer5yo.jpg"))
					},
				},
				{
					Name: "CreateMovie2",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "Titanic", "A seventeen-year-old aristocrat falls in love with a kind but poor artist aboard the luxurious, ill-fated R.M.S. Titanic.", "James Cameron", "Romance", 3.14, 1997, "https://res.cloudinary.com/domtywpyc/image/upload/v1676474113/Titanic_i62e7s.jpg"))
					},
				},
				{
					Name: "CreateMovie3",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "Avatar", "A paraplegic Marine dispatched to the moon Pandora on a unique mission becomes torn between following his orders and protecting the world he feels is his home.", "James Cameron", "Science Fiction", 2.42, 2009, "https://res.cloudinary.com/domtywpyc/image/upload/v1676474112/avatar_gqngv2.jpg"))
					},
				},
				{
					Name: "CreateMovie4",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "Star Wars: Episode VII - The Force Awakens", "As a new threat to the galaxy rises, Rey, a desert scavenger, and Finn, an ex-stormtrooper, must join Han Solo and Chewbacca to search for the one hope of restoring peace.", "J.J. Abrams", "Science Fiction", 2.18, 2015, "https://res.cloudinary.com/domtywpyc/image/upload/v1676474112/starwars-forceawakens_exrzqf.jpg"))
					},
				},
				{
					Name: "CreateMovie5",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "Avengers: Endgame", "After the devastating events of Avengers: Infinity War (2018), the universe is in ruins. With the help of remaining allies, the Avengers assemble once more in order to reverse Thanos' actions and restore balance to the universe.", "Anthony Russo", "Fantasy", 3.01, 2019, "https://res.cloudinary.com/domtywpyc/image/upload/v1676474112/avengers-endgame_jen8hk.jpg"))
					},
				},
				{
					Name: "CreateMovie6",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "Knocked Up", "For fun-loving party animal Ben Stone, the last thing he ever expected was for his one-night stand to show up on his doorstep eight weeks later to tell him she's pregnant with his child.", "Judd Apatow", "Comedy", 2.09, 2007, "https://res.cloudinary.com/domtywpyc/image/upload/v1676551286/Knocked_up_ga7ssk.jpg"))
					},
				},
				{
					Name: "CreateMovie7",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "Jurassic World", "A new theme park, built on the original site of Jurassic Park, creates a genetically modified hybrid dinosaur, the Indominus Rex, which escapes containment and goes on a killing spree.", "Colin Trevorrow", "Science Fiction", 2.04, 2015, "https://res.cloudinary.com/domtywpyc/image/upload/v1676474112/jurassic-world_qmf3tg.jpg"))
					},
				},
				{
					Name: "CreateMovie8",
					Run: func(db *gorm.DB) error {
						return (
							CreateMovie(db, uuid.New().String(), "Captain Phillips", "The true story of Captain Richard Phillips and the 2009 hijacking by Somali pirates of the U.S.-flagged MV Maersk Alabama, the first American cargo ship to be hijacked in two hundred years.", "Paul Greengrass", "Biography", 2.14, 2013, "https://res.cloudinary.com/domtywpyc/image/upload/v1676551561/Captain_phillips_nexwsm.jpg"))
					},
				},
		})
}
