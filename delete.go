package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//DeletePicture deletes a picture
func DeletePicture(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	Picture := Picture{}
	id := c.Param("id")

	db.Preload("Tags").First(&Picture, id)
	db.Delete(&Picture)

	c.JSON(http.StatusOK, gin.H{"status": "Picture has been deleted!"})

	db.Close()
}

//DeleteTag deletes a tag
func DeleteTag(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	Tag := Tag{}
	id := c.Param("id")

	db.Preload("Picture").First(&Tag, id)
	db.Delete(&Tag)

	c.JSON(http.StatusOK, gin.H{"status": "Tag has been deleted!"})

	db.Close()
}

//DisconnectTag deletes Tag from Picture
//AppendTag adds tag to picture
func DisconnectTag(c *gin.Context) {
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
	db.Model(&Picture).Association("Tags").Delete(&Tag)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
	db.Close()
}
