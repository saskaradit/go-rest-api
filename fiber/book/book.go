package book

import (
	"saskara/rest-api-playground/fiber/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Rating string `json:"Rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)

}
func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No book found")
	}

	newBook := new(Book)
	if err := c.BodyParser(newBook); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Model(&book).Update(&newBook)
	c.JSON(book)
}
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found")
	}
	db.Delete(&book)
	c.Send("Book successfully deleted")
}
