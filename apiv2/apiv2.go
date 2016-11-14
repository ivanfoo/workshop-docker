package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/toqueteos/minietcd"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	listenAddr = flag.String("listen", ":8000", "listen address")
	etcdServer = flag.String("etcd-server", "", "etcd server")
	etcdDir    = flag.String("etcd-dir", "", "etcd root dir")
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName string        `json:"fullname" binding:"required"`
	Email    string        `json:"email" binding:"required"`
	Password string        `json:"password" binding:"required"`
}

func main() {
	flag.Parse()

	if *etcdServer == "" {
		*etcdServer = os.Getenv("ETCD_SERVER")
	}

	if *etcdDir == "" {
		*etcdDir = os.Getenv("ETCD_DIR")
	}

	etcd := minietcd.New()
	if err := etcd.Dial(*etcdServer); err != nil {
		log.Fatal(err)
	}

	etcdKeys, _ := etcd.Keys(*etcdDir)
	for k, v := range etcdKeys {
		log.Println(k, v)
	}

	conn, err := mgo.Dial(etcdKeys["dbServer"])
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/users/:id", func(c *gin.Context) {
		db := conn.DB(etcdKeys["dbName"]).C("users")
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

		db := conn.DB(etcdKeys["dbName"]).C("users")
		_, err = db.Upsert(bson.M{"email": user.Email}, user)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "created",
		})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		db := conn.DB(etcdKeys["dbName"]).C("users")
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
