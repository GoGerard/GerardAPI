package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//UpdatePicture updates picture in DB and sets time to now.
func UpdatePicture(c *gin.Context) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var UpdatePicture Picture

	if c.BindJSON(&UpdatePicture) == nil {
		UpdatePicture.CreateTime = time.Now()
		out := db.Save(&UpdatePicture)
		c.JSON(http.StatusCreated, out)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR: Picture not saved"})
	}

	db.Close()
}
