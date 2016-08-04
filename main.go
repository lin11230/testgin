package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

// DB handle mongodb
type DB struct {
	session *mgo.Session
}

type Person struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	//router := gin.New()

	router.Use(MongoDBMiddleware())

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//router.Run()
	router.Run(":3000") //for a hard coded port
}

func getting(c *gin.Context) {
	fmt.Println("I'm in Getting.")
	dbconn, ok := c.MustGet("databaseConn").(*mgo.Session)
	if !ok {
		fmt.Println("GG")
	}
	defer dbconn.Close()

	err := dbconn.Ping()
	if err != nil {
		log.Println("Ping DB failed.")
	}

	con := dbconn.DB("test").C("people")

	// Query All
	var results []Person
	err = con.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
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
func MongoDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// db init
		mongodb := "127.0.0.1"
		session, err := mgo.Dial(mongodb)
		if err != nil {
			log.Println("cannot connect to mongo, error:", err)
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c.Set("databaseConn", session)
		c.Next()

	}
}
