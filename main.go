package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

// DB handle mongodb
type DB struct {
	session *mgo.Session
}

func main() {
	// db init
	mongodb := "127.0.0.1"
	session, err := mgo.Dial(mongodb)
	if err != nil {
		log.Println("cannot connect to mongo, error:", err)
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.Use(MongoDB(session))
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

func getting(c *gin.Context) {
	fmt.Println("I'm in Getting.")
	dbconn, ok := c.MustGet("dbcon").(*mgo.Session)
	if !ok {
		log.Println("connection failed.")
	}

	err := dbconn.Ping()
	if err != nil {
		log.Println("Ping DB failed.")
	}
}
func posting(c *gin.Context) {
	fmt.Println("I'm in Posting.")
}
func putting(c *gin.Context) {
	fmt.Println("I'm in Putting.")
}
func deleting(c *gin.Context) {
	fmt.Println("I'm in Deleting.")
}

// middleware
// MongoDB middleware for echo framework
func MongoDB(dbsession *mgo.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbcon", dbsession)
		c.Next()
	}
}
