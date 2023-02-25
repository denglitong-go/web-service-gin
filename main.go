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
	"github.com/gin-gonic/gin"
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
	// the Go default way without Gin
	// http.HandleFunc("/", defaultHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// the Gin way
	router := gin.Default()

	router.GET("/albums", getAlbums, requestInterceptor)
	router.GET("/albums/:id", getAlbumByID, requestInterceptor)
	router.POST("/albums", postAlbums)

	log.Fatal(router.Run(":8080"))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(albums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func requestInterceptor(c *gin.Context) {
	// TODO to intercept c.ResponseWriter
	log.Printf("request: %v\n", c.Request)

	// the Context.IndentedJSON/JSON is not a terminal operation,
	// the JSON result will be appended to the response writer if there are
	// multiple handler are registered of the same request path.
	// c.IndentedJSON(http.StatusOK, albums[0])
}

// getAlbums responds with the list of all albums as JSON
// gin.Context is the most important part of Gin. It carries request details, validates and
// serializes JSON, and more.
func getAlbums(c *gin.Context) {
	// Call Context.IndentedJSON to serialize the struct into JSON
	// and add it to the response.
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// User Context.BindJSON to bind the request body to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// URL query param, GET /album?id=1
	// id := c.Query("id")
	// URL path param, serve /album/:id, GET /album/1 => id=1
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
