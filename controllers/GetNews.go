package controllers

// import (
// 	"fiber-ebook-api/responses"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// )

// func GetNews(c *fiber.Ctx) error {

// 	id := c.Params("id")
// 	files, err := ioutil.ReadDir("./news/" + id + "/")
// 	if err != nil {
// 		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{Status: http.StatusNotFound, Message: "Oops", Data: &fiber.Map{"data": err}})
// 	}
// 	filePath := "./news/" + id + "/" + files[0].Name()
// 	return c.Download(filePath)
// }
