 # GO WEB 

 This Go program is a simple web service using the Gin framework. It manages a collection of albums and provides
 endpoints to retrieve the list of albums, retrieve a specific album by its ID, and add a new album to the collection.


To send a response to a GET request fo an album, you can use the `getAlbums` function. This function retrieves the list of albums and sends it back to the client in JSON format. Here is an example of how you can implement this function:

```go
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
```

In this function:
- `c *gin.Context` is the context for the request, which provides methods to handle the request and response.
- `c.IndentedJSON` is used to send a JSON response with indentation for readability.
- `http.StatusOK` is the HTTP status code for a successful response.
- `albums` is the slice of Album structs that contains the list of albums.

To handle a GET request for a specific album by its ID, you can use the `getAlbumID` function:

```go
// getAlbumID responds with the album whose ID matches the id parameter sent by the client.
func getAlbumID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```

In this function:
- `c.Param("id")` retrieves the ID parameter from the request URL.
- The function loops through the `albums` slice to find an album with the matching ID.
- If a matching album is found, it sends the album as a JSON response with a status code of `http.StatusOK`.
- If no matching album is found, it sends a 404 status code with an error message.

These functions should be registered as handlers for the respective endpoints in the `main` function:

```go
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumID)
    router.Run("localhost:8080")
}
```

 The main components of the code are:
 - The `Album` struct: Defines the structure of an album with fields for ID, Title, Artist, and Price.
 - The `album` slice: A pre-populated list of albums.
 - The `main` function: Sets up the Gin router, defines the endpoints, and starts the web server on localhost:8080.
 - The `getAlbums` function: Handles GET requests to the "/albums" endpoint and returns the list of albums in JSON format.
 - The `postAlbum` function: Handles POST requests to the "/albums" endpoint, adds a new album to the collection, and returns the new album in JSON format.
 - The `getAlbumID` function: Handles GET requests to the "/albums/:id" endpoint, retrieves an album by its ID, and returns it in JSON format. If the album is not found, it returns a 404 status with an error message.


 album is a slice of Album structs that contains a list of albums with their respective details.
 Each Album struct includes the following fields:
 - ID: A unique identifier for the album.
 - Title: The title of the album.
 - Artist: The artist who created the album.
 - Price: The price of the album.

