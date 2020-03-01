package main

import (
	"fmt"
	"os"
	"strings"

	"./scrapper"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrapper(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	e := echo.New()
	gErr := e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	fmt.Println("test", gErr)
	fmt.Println("done")
}
