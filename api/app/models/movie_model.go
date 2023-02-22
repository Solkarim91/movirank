package models

type Movie struct {
    ID    		  int    	`gorm:"primary key;autoIncrement" json:"id"`
    Title     	string 	`json:"title"`
    Description string 	`json:"description"`
    Genre 			string 	`json:"genre"`
		Runtime 		float64	`json:"runtime"`
		Released 		int 		`json:"released"`
		Rating 			int 		`json:"rating"`
		Img 				string 	`json:"img"`
}

type Comment struct {
		ID 		int    `gorm:"primary key;autoIncrement" json:"id"`
		User 	string `json:"user"` 
		Text 	string `json:"text"`
}