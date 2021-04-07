package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Let's create our Jokes struct. This will contain information about a Joke

// Joke contains information about a single Joke
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

// We'll create a list of jokes
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

func main() {
	router := setupServer()

	// Start and run the server
	router.Run(":3000")
}

func setupServer() *gin.Engine {
	// Set the router as the default one shipped with Gin
	r := gin.Default()

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	initializeRoutes(r)

	return r
}

func initializeRoutes(r *gin.Engine) *gin.RouterGroup {
	// Setup route group for the API
	api := r.Group("/api")
	{
		// set up the API routes
		api.GET("/", pingHandler)
		// Handle GET requests at /api/jokes
		// will retrieve a list of jokes a user can see
		api.GET("/jokes", JokeHandler)
		// Handle POST requests at /api/jokes/like/<jokeID>
		// will capture likes sent to a particular joke
		api.POST("/jokes/like/:jokeID", LikeJoke)
	}

	return api
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, jokes)
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
	// confirm Joke ID sent is valid
	// remember to import the `strconv` package
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		// find joke, and increment likes
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes += 1
			}
		}
		// return a pointer to the updated jokes list
		c.JSON(http.StatusOK, &jokes)
	} else {
		// Joke ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}
