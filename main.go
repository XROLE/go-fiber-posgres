package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	Author string			 `json:"author"`
	Title string		 	`json:"title"`
	Publishe string 		`json:"publisher"`
}

type Reposirory struct {
	DB *gorm.DB
}

func (r *Reposirory) SetupTRoutes(app *fiber.App) {
	api := app.Group(("/api"))
	api.Post("/create_books", r.CreateBooks)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")

	if(err != nil){
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config.Storage)

	if(err != nil){
		log.Fatal("could not load the database")
	}

	r := Reposirory {
		DB : db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.listen(":8080")
}