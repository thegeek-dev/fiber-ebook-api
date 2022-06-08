package controllers

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FindPub(c *fiber.Ctx) error {
	searchTerm := c.Query("searchTerm")
	file, err := os.Open("./NewsRecipes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	foundRecipes := []string{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		if strings.Contains(strings.ToUpper(scanner.Text()), strings.ToUpper(searchTerm)) {
			foundRecipes = append(foundRecipes, scanner.Text())
		}
	}

	return c.Status(http.StatusOK).JSON(foundRecipes)
}
