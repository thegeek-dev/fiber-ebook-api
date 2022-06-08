package controllers

import (
	"fiber-ebook-api/responses"
	"fmt"
	"mime/multipart"
	"os/exec"

	"net/http"
	"os"

	"path/filepath"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveEbooks(c *fiber.Ctx) error {
	format := c.Query("format")
	if format == "" {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{Status: http.StatusNotFound, Message: "Oops", Data: &fiber.Map{"data": "Please provide correct format"}})
	}
	id := uuid.New()
	if err := os.Mkdir("./uploads/"+id.String(), os.ModePerm); err != nil {
		fmt.Println(err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "Oops", Data: &fiber.Map{"data": err}})
	}
	files := form.File["ebooks"]

	for _, file := range files {
		if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s/%s", id.String(), file.Filename)); err != nil {
			fmt.Println(err)
		}
	}

	go ConvertBooks(id.String(), files, format)
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "Converted Successfully", Data: &fiber.Map{"data": id.String()}})
}

func ConvertBooks(uuid string, files []*multipart.FileHeader, format string) {

	for _, file := range files {
		pathPrefix := "./uploads/" + uuid + "/"
		filenameWithoutExtension := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
		inputFile := pathPrefix + file.Filename
		outputFile := pathPrefix + filenameWithoutExtension + "." + format

		archive := "./archives/" + uuid + ".zip"

		if _, err := exec.Command("ebook-convert", inputFile, outputFile).Output(); err != nil {
			fmt.Println(err)
		}

		if err := os.Remove(inputFile); err != nil {
			fmt.Println(err)
		}

		if _, err := exec.Command("zip", "-j", archive, outputFile).Output(); err != nil {
			fmt.Println(err)
		}

		if err := os.Remove(outputFile); err != nil {
			fmt.Println(err)
		}
	}
	os.Remove("./uploads/" + uuid)

}
