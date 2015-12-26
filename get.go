package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllPictures receives []Picture{} from database
func GetAllPictures(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	Pictures := []Picture{}
	db.Preload("Tags").Find(&Pictures)

	c.JSON(http.StatusOK, Pictures)

	db.Close()
}

//GetAllTags receives []Tag{} from database
func GetAllTags(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	Tags := []Tag{}
	db.Find(&Tags)

	c.JSON(http.StatusOK, Tags)

	db.Close()
}

//GetIDPicture returns requested Picture{}
func GetIDPicture(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	Picture := Picture{}
	id := c.Param("id")

	db.Preload("Tags").First(&Picture, id)

	c.JSON(http.StatusOK, Picture)
	db.Close()
}

// //GetIDTag returns requested Tag{}
// func GetIDTag(c *gin.Context) {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	Tag := Tag{}
// 	id := c.Param("id")
//
// 	db.Preload("Pictures").First(&Tag, id)
//
// 	c.JSON(http.StatusOK, Tag)
//
// 	db.Close()
// }

// //GetPictureTags returns requested []Tags{} from Picture{}
// func GetPictureTags(c *gin.Context) {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	Picture := Picture{}
// 	id := c.Param("id")
//
// 	db.Preload("Tags").First(&Picture, id)
//
// 	c.JSON(http.StatusOK, Picture)
//
// 	db.Close()
// }
//
// //GetTagPictures returns requested []Pictures{} from Tag{}
// func GetTagPictures(c *gin.Context) {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	Tag := Tag{}
// 	id := c.Param("id")
//
// 	db.Preload("Pictures").First(&Tag, id)
//
// 	c.JSON(http.StatusOK, Tag.Pictures)
//
// 	db.Close()
// }
