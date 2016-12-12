package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	listenAddr = flag.String("listen", ":8000", "listen address")
	dbServer   = flag.String("db", "", " db server")
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName string        `json:"fullname" binding:"required"`
	Email    string        `json:"email" binding:"required"`
	Password string        `json:"password" binding:"required"`
}

func main() {
	flag.Parse()

	if *dbServer == "" {
		*dbServer = os.Getenv("DB_SERVER")
	}

	conn, err := mgo.Dial(*dbServer)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/users/:id", func(c *gin.Context) {
		db := conn.DB("demo").C("users")
		id := bson.ObjectIdHex(c.Param("id"))
		user := User{}
		err = db.FindId(id).One(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "not found",
			})
		}
		c.JSON(http.StatusOK, user)
	})

	router.POST("/users", func(c *gin.Context) {
		var user User
		if c.BindJSON(&user) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
			})
			return
		}

		db := conn.DB("demo").C("users")
		_, err = db.Upsert(bson.M{"email": user.Email}, user)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":     user.Id,
			"status": "created",
		})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		db := conn.DB("demo").C("users")
		id := bson.ObjectIdHex(c.Param("id"))
		err = db.RemoveId(id)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "deleted",
		})
	})

	router.Run(*listenAddr)
}
