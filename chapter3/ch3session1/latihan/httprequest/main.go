package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	BookID   string `json:"book_id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	AuthorID int    `json:"author_id"`
}

var baseUrl = "http://localhost:8083"

func main() {
	// get()
	// post()
	auth()
}

func get() {
	url := fmt.Sprintf("%s/books", baseUrl)

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("status not OK")
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(data))

	books := make([]Book, 0)

	err = json.Unmarshal(data, &books)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(books))
	fmt.Println(books)
}

func post() {
	url := fmt.Sprintf("%s/books", baseUrl)

	client := http.Client{}

	b := Book{
		BookID:   "12",
		Title:    "Fantastic Beast and Where to Find Them",
		Price:    150000,
		AuthorID: 1,
	}

	bookData, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer(bookData)
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	fmt.Println(request)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error")
		return
	}

	dataResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var createdBook Book

	err = json.Unmarshal(dataResponse, &createdBook)

	fmt.Println(createdBook)
}

func auth() {
	url := fmt.Sprintf("%s/", baseUrl)

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		panic(err)
	}

	// authHeader := base64.StdEncoding.EncodeToString([]byte("hah:wah"))

	// request.Header.Set("Authorization", "Basic "+authHeader)

	request.SetBasicAuth("hah", "wah")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("status not OK")
		return
	}

	dataResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataResponse))
}
