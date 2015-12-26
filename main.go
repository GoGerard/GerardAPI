package main

import (
	"log"
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

//API = Default API path
var API = "/api/"

//CORSConfig = HTTP CORS Config
var CORSConfig = cors.Config{
	AbortOnError:    false,
	AllowAllOrigins: true,
	AllowedMethods:  []string{"GET", "POST", "PUT", "PATCH", "HEAD", "DELETE"},
	AllowedHeaders:  []string{"Content-Type", "X-Requested-With", "Authorization"},
	//ExposedHeaders:   "",
	AllowCredentials: false,
	MaxAge:           12 * time.Hour,
}

func main() {
	log.Println("* Initialize DB")

	err := InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Println("* Start API")

	router := gin.Default()

	router.Use(cors.New(CORSConfig))

	auth := router.Group("/")
	auth.Use(AuthTokenMiddleware())

	//AUTH
	router.POST(API+"auth", AuthToken)
	router.PUT(API+"auth", DelAuthToken)

	//GET
	auth.GET(API+"pictures", GetAllPictures)
	auth.GET(API+"tags", GetAllTags)
	auth.GET(API+"pictures/:id", GetIDPicture)
	//router.GET(API+"tags/:id", GetIDTag)
	//router.GET(API+"pictures/:id/tags", GetPictureTags)
	//router.GET(API+"tags/:id/pictures", GetTagPictures)

	//POST
	auth.POST(API+"pictures", CreatePicture)
	auth.POST(API+"tags", CreateTag)

	//PUT
	auth.PUT(API+"pictures", UpdatePicture)
	auth.PUT(API+"pictures/:id/add/:tag", AppendTag)

	//	router.PUT(API+"tags"/:id, UpdateTag)
	//	router.PUT(API+"pictures/:id/tags", UpdatePictureTags)

	//DELETE
	auth.DELETE(API+"pictures/:id", DeletePicture)
	auth.DELETE(API+"tags/:id", DeleteTag)
	auth.DELETE(API+"pictures/:id/del/:tag", DisconnectTag)

	//	router.DELETE(API+"pictures/:id/tags", DeletePictureTag)

	router.Run(":8080")

}
