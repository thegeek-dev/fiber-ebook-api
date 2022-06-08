package controllers

import (
	"fiber-ebook-api/responses"

	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetEbooks(c *fiber.Ctx) error {
	noZip := false
	noDir := false
	uuid := c.Query("uuid")
	filePathZip := "./archives/" + uuid + ".zip"
	filePathDir := "./uploads/" + uuid

	if _, err := os.Stat(filePathZip); err != nil {
		noZip = true
	}
	if _, err := os.Stat(filePathDir); err != nil {
		noDir = true
	}

	if noZip && noDir {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{Status: http.StatusNotFound, Message: "Oops", Data: &fiber.Map{"data": "Looks like the conversion failed"}})
	}
	if noZip && !noDir {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{Status: http.StatusNotFound, Message: "Oops", Data: &fiber.Map{"data": "Conversion Still in process"}})
	}
	if !noZip && !noDir {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{Status: http.StatusNotFound, Message: "Oops", Data: &fiber.Map{"data": "Conversion Still in process"}})
	}
	return c.Download(filePathZip)

	// return c.Status(http.StatusAccepted).JSON(responses.UserResponse{Status: http.StatusAccepted, Message: "Buhhooo", Data: &fiber.Map{"data": "You should be prompted to download"}})
}
