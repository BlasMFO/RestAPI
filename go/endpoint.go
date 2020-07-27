package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gofiber/fiber"
)

// Home -> home route
func Home(c *fiber.Ctx) {
	fmt.Println("Home -> func")

	c.Send("This endpoint is -> Home")
}

// SelectMovie -> get the recommendations of the selected movie
func SelectMovie(c *fiber.Ctx) {
	fmt.Println("SelectMovie -> func")

	movieParam := c.Params("user_movie")
	fmt.Println(movieParam)

	movie := usermovie{Movie: movieParam}

	movieJSON, _ := json.Marshal(movie)

	// Write the JSON file with the user_movie name
	err := ioutil.WriteFile("/rest/usermovie.json", movieJSON, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Wrtite the number 1 into read (file to allow the python script process the request)
	data := []byte("1")
	err = ioutil.WriteFile("/rest/read", data, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// Erases old recommendation
	_ = os.Remove("/rest/rcmd_movies.json")

	chanRcmd := make(chan []byte, 1)

	go func() {
		chanIs := readOut()
		chanRcmd <- chanIs
	}()

	select {
	case res := <-chanRcmd:
		result := rcmd{}
		_ = json.Unmarshal(res, &result)
		c.JSON(result)
	case <-time.After(6 * time.Second):
		c.Status(504).Send("No recommendations found :(")
		fmt.Println("Timeout")
	}
}

func readOut() []byte {
	err := errors.New("no file read yet")
	jsonFile := new(os.File)
	byteValue := make([]byte, 0)

	for err != nil {
		time.Sleep(10 * time.Millisecond)
		jsonFile, err = os.Open("/rest/rcmd_movies.json")
		if err != nil {
			defer jsonFile.Close()
			continue
		}
		byteValue, err = ioutil.ReadAll(jsonFile)
		if err != nil {
			defer jsonFile.Close()
			continue
		}
	}
	return byteValue
}
