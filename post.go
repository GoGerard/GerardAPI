package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//CreatePicture creates a picture
func CreatePicture(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var NewPicture Picture

	if c.BindJSON(&NewPicture) == nil {
		NewPicture.CreateTime = time.Now()
		out := db.Create(&NewPicture)
		c.JSON(http.StatusCreated, out)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR: Picture not created"})
	}

	db.Close()
}

//CreateTag creates a Tag
func CreateTag(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var NewTag Tag

	if c.BindJSON(&NewTag) == nil {
		out := db.Create(&NewTag)
		c.JSON(http.StatusCreated, out)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR: Tag not created"})
	}

	db.Close()
}

//AppendTag adds tag to picture
func AppendTag(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	Picture := Picture{}
	Tag := Tag{}
	PictureID := c.Param("id")
	TagID := c.Param("tag")

	err = db.First(&Picture, PictureID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR: Picture not found"})
		db.Close()
		return
	}
	err = db.First(&Tag, TagID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR: Tag not found"})
		db.Close()
		return
	}
	db.Model(&Picture).Association("Tags").Append(&Tag)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
	db.Close()
}
