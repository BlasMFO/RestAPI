package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var c = make(chan []byte, 1)

// Home -> home route
func Home(ctx *fiber.Ctx) (err error) {
	fmt.Println("Home -> func")
	if err != nil {
		return ctx.SendStatus(404)
	}
	return ctx.SendString("This endpoint is -> Home")
}

// SelectMovie -> get the recommendations of the selected movie
func SelectMovie(ctx *fiber.Ctx) (err error) {
	if err != nil {
		return ctx.SendStatus(500)
	}
	fmt.Println("SelectMovie -> func")

	movieParam := ctx.Params("user_movie")
	fmt.Println(movieParam)

	movie := usermovie{Movie: movieParam}

	movieJSON, _ := json.Marshal(movie)

	// Write the JSON file with the user_movie name
	err = ioutil.WriteFile("/rest/usermovie.json", movieJSON, 0644)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(500)
	}

	// Wrtite the number 1 into read (file to allow the python script process the request)
	data := []byte("1")
	err = ioutil.WriteFile("/rest/read", data, 0644)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(500)
	}

	// Erases old recommendation
	_ = os.Remove("/rest/rcmd_movies.json")

	go func() {
		chanIs := readOutput()
		c <- chanIs
	}()

	select {
	case res := <-c:
		result := rcmd{}
		_ = json.Unmarshal(res, &result)
		return ctx.JSON(result)
	case <-time.After(6 * time.Second):
		fmt.Println("Timeout")
		return ctx.Status(501).SendString("No recommendations found :(")
	}
}

func readOutput() []byte {
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
