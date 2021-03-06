package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetNews(c *fiber.Ctx) error {
	curentTime := time.Now()
	loc, _ := time.LoadLocation("UTC")
	utcTime := curentTime.In(loc)
	today := utcTime.Format("02-01-2006")
	fmt.Println(today)
	yesterday := utcTime.AddDate(0, 0, -1).Format("02-01-2006")
	isPresent := true
	pubName := c.Query("pubName")
	format := c.Query("format")
	if err := os.Mkdir("./news/"+today, os.ModePerm); err == nil {
		fmt.Println("Changing Date")
		os.Remove("./news/" + yesterday)
	}
	if format == "html" {
		if err := os.Mkdir("./news/"+today+"/"+pubName, os.ModePerm); err == nil {
			fmt.Println("Fetching First Time")
		}
		inputFile := pubName + ".recipe"
		outputDir := "./news/" + today + "/" + pubName + "/"
		outputFile := outputDir + pubName + ".htmlz"
		fileToBeSent := outputDir + "index.html"
		renamedFile := outputDir + pubName + ".html"

		if _, err := os.Stat(renamedFile); err == nil {
			return c.Download(renamedFile)
		}
		if _, err := exec.Command("ebook-convert", inputFile, outputFile, "--htmlz-class-style", "inline").Output(); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"Message": "Oops, fetching failed"})
		}
		if _, err := exec.Command("unzip", outputFile, "-d", outputDir).Output(); err != nil {
			fmt.Println(err)
		}

		if _, err := exec.Command("mv", fileToBeSent, renamedFile).Output(); err != nil {
			fmt.Println(err)
		}
		return c.Download(renamedFile)

	}
	if _, err := os.Stat("./news/" + today + "/" + pubName + "." + format); err != nil {
		isPresent = false
	}

	if !isPresent {
		FetchNews(pubName, format, today)
	}

	return c.Download("./news/" + today + "/" + pubName + "." + format)
}

func FetchNews(pubName string, format string, today string) {
	inputFile := pubName + ".recipe"
	outputFile := "./news/" + today + "/" + pubName + "." + format
	if _, err := exec.Command("ebook-convert", inputFile, outputFile).Output(); err != nil {
		fmt.Println(err)
	}
}
