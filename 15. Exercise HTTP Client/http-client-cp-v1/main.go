package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

type listOfAnimechan struct {
	listData []Animechan
}

func ClientGet() ([]Animechan, error) {
	client := &http.Client{}

	url := "https://animechan.vercel.app/api/quotes/anime?title=naruto"

	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var listOfAnimechan []Animechan

	err = json.Unmarshal(body, &listOfAnimechan)
	if err != nil {
		panic(err)
	}

	return listOfAnimechan, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)

	if err != nil {
		return Postman{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, err
	}

	var postmanResp Postman

	err = json.Unmarshal(body, &postmanResp)
	if err != nil {
		return Postman{}, err
	}

	return postmanResp, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
