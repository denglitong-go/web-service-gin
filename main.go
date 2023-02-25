// Gin simplifies many tasks associated with building web applications,
// including web services. In this tutorial, you'll use Gin to route requests,
// retrieve request details, and marshal JSON response.
// In this tutorial, you will build a RESTFUL API server with two endpoints.
// Includes:
// 	1.Design API endpoints.
// 	2.Create a folder for your code.
//  3.Create the data.
// 	4.Write a handler to return all items.
//  5.Write a handler to add a new item.
//  6.Write a handler to return a specific item.
//
// Endpoints in this tutorial:
// /albums
// 	GET - Get a list of all albums, returned as JSON.
// 	POST - Add a new album from request data sent as JSON.
// /albums/:id
//	GET - Get an album by its ID, returning the album data as JSON.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var (
	albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Cliford Brown", Artist: "Sarah Varghan", Price: 39.99},
	}
)

func main() {
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(albums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}
