package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//AuthToken GET request: returns possible issues.
func AuthToken(c *gin.Context) {
	var UserToken APISession

	if c.BindJSON(&UserToken) == nil {
		err := CheckToken(UserToken.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Someone set up us the bomb."})
	}
}

//DelAuthToken DELETE auth request: destroys dreams.
func DelAuthToken(c *gin.Context) {
	var UserToken APISession

	if c.BindJSON(&UserToken) == nil {
		err := DelToken(UserToken.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Someone set up us the bomb."})
	}
}

//CheckToken does a DB token check, token might be destroyed in process.
func CheckToken(CheckToken string) error {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	//Search for Token
	DBSession := new(APISession)
	db.Where("token = ?", CheckToken).Find(&DBSession)
	log.Println(DBSession)
	if DBSession.UserID == "" {
		return errors.New("Token not found!")
	}

	//Check token expiration
	Now := time.Now()
	SessionAge := Now.Sub(DBSession.Time).Minutes()
	Expiration := (10 * time.Minute).Minutes()
	if SessionAge > Expiration {
		db.Delete(DBSession)
		return errors.New("Token expired and destroyed! Thanks by the way, DB is lazy.")
	}

	//Update Token Activity
	DBSession.Time = time.Now()
	db.Save(&DBSession)
	db.Close()

	return nil
}

//DelToken destoys token for you nicely.
func DelToken(DelToken string) error {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	//Search for Token
	DBSession := new(APISession)
	db.Where("token = ?", DelToken).Find(&DBSession)
	log.Println(DBSession)
	if DBSession.UserID == "" {
		return errors.New("Token not found!")
	}

	db.Delete(DBSession)
	db.Close()

	return nil
}

//HorribleDummyAuth THE WORST TOKEN CHECK POSSIBLE
func HorribleDummyAuth(CheckToken string) error {
	if CheckToken != "luek" {
		return errors.New("WOW NEE")
	}

	return nil
}
