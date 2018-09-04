package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {

	// If running in Heroku...
	var port = os.Getenv("PORT")

	fmt.Println(port)

	// Check for env PORT. If not, set it manually.
	if port == "" {
		port = "8787"
	}

	// Create Echo instance for server routing stuff
	e := echo.New()

	// Structs
	type Response struct {
		Message string
	}

	// For the index page
	e.Static("/", "./src/")

	// GET from Our Mannna for Verse Of The Day (VOTD)
	// when client sends request to this path
	e.GET("/votd/", func(c echo.Context) error {
		fmt.Println("\nSending out for VOTD...")
		resp := MakeRequest()
		return c.JSON(http.StatusOK, resp)
	})

	// Log the listening port
	e.Logger.Fatal(e.Start(":" + port))

}

// URL to get VOTD from
const votdUrl string = "https://beta.ourmanna.com/api/v1/get/?format=json"

// Interfaces for JSON export
type Details struct {
	Text      string `json:"text"`
	Reference string `json:"reference"`
	Version   string `json:"version"`
	VerseUrl  string `json:"verseUrl"`
}
type Verse struct {
	Details Details `json:"details"`
}
type OurMannaAPIResponse struct {
	Verse Verse `json:"verse"`
}

// Convert request body to JSON using our interfaces
func getVOTD(body []byte) (*OurMannaAPIResponse, error) {
	var o = new(OurMannaAPIResponse)
	err := json.Unmarshal(body, &o)
	if err != nil {
		fmt.Println("Error getting VOTD:", err)
	}
	return o, err
}

// Request-makin' function
func MakeRequest() *OurMannaAPIResponse {

	resp, err := http.Get(votdUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	s, err := getVOTD([]byte(body))

	// log.Println(s)
	return s
}
