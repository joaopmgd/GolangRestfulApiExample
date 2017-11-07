package main


/*
 * Project:     GolangRestfulApiExample
 * Created by:  joaopmgd@gmail.com | JP
 * Date:        01th November 2017
 */

import (
"encoding/json"
"net/http"

"github.com/gin-gonic/gin"
log "github.com/sirupsen/logrus"
"io"
"io/ioutil"
"bytes"

"GolangRestfulApiExample/errorStatus"
"GolangRestfulApiExample/mongoDB"
"GolangRestfulApiExample/service"
)

func main (){

	// Initialize Logs
	initLog()
	// Initialize the Gin Package, listening for REST calls
	r := gin.Default()
	// Verify if the server is up
	r.GET("/api/user/healthCheck", healthCheck)
	// Insert New User
	r.POST("/api/user", postNewUser)
	// Get User
	r.GET("/api/user", getUser)
	// Update User Data
	r.PATCH("/api/user", patchUser)
	// Delete User
	r.DELETE("/api/user", deleteUser)
	// Initialize Server
	r.Run()
}

// postNewUser returns an error or Accepted status.
// users with the same socialNumber will not be inserted and will return an error message.
func postNewUser(c *gin.Context) {
	log.Info("[Function] postNewUser")

	var user mongoDB.User

	json.NewDecoder(c.Request.Body).Decode(&user)

	logRequest(c.Request.Header, c.Request.Body)

	if user.Name == "" || user.SocialNumber == "" || user.Age == "" || user.Phones == nil ||
		user.Address.Country == "" || user.Address.State == "" || user.Address.StreetName == "" ||
			user.Address.ZipCode == "" || user.Address.StreetNumber == ""{

		em := errorStatus.ErrorMissingBodyValues()
		c.JSON(em.HTTPStatus, em)
		log.Info(toJSON(user),toJSON(em))
		return
	}

	err := service.PostNewUserService(user)
	if err.HTTPStatus != 200 {
		c.JSON(err.HTTPStatus, err)
	} else {
		c.JSON(err.HTTPStatus, "New User Accepted.")
	}
	log.Info(toJSON(user),toJSON(err))
}

// getUser returns the complete user data.
// The header of the request must contain the user socialNumber.
func getUser(c *gin.Context) {
	log.Info("[Function] getUser")

	socialNumber := c.Request.Header.Get("socialNumber")

	logRequest(c.Request.Header, c.Request.Body)

	if socialNumber == ""{
		em := errorStatus.ErrorMissingHeaderValues()
		c.JSON(em.HTTPStatus, em)
		log.WithFields(log.Fields{"socialNumber": socialNumber}).Info(toJSON(em))
		return
	}
	user, err := service.GetUserService(socialNumber)
	if err.HTTPStatus != 200 {
		c.JSON(err.HTTPStatus, err)
	} else {
		c.JSON(err.HTTPStatus, user)
	}
	log.Info(toJSON(user),toJSON(err))
}

// updateUser update user with the selected data.
// The body of the request must have the updated data.
// The return will be an error or Accepted.
func patchUser(c *gin.Context) {
	log.Info("[Function] patchUser")

	var user mongoDB.User

	socialNumber := c.Request.Header.Get("socialNumber")
	json.NewDecoder(c.Request.Body).Decode(&user)

	logRequest(c.Request.Header, user)

	if socialNumber != user.SocialNumber{
		err := errorStatus.ErrorSocialNumberCannotBeChanged()
		c.JSON(err.HTTPStatus, err)
		log.Info(toJSON(err))
		return
	}

	if user.Name == "" && user.Age == "" && user.Phones == nil &&
		user.Address.Country == "" && user.Address.State == "" && user.Address.StreetName == "" &&
		user.Address.ZipCode == "" && user.Address.StreetNumber == ""{
		em := errorStatus.ErrorMissingBodyValues()
		c.JSON(em.HTTPStatus, em)
		log.Info(toJSON(em))
		return
	}
	if socialNumber == ""{
		em := errorStatus.ErrorMissingHeaderValues()
		c.JSON(em.HTTPStatus, em)
		log.WithFields(log.Fields{"socialNumber": socialNumber}).Info(toJSON(em))
		return
	}
	err := service.PatchUserService(socialNumber, user)
	if err.HTTPStatus != 200 {
		c.JSON(err.HTTPStatus, err)
	} else {
		c.JSON(err.HTTPStatus, "User Permissions Updated.")
	}
	log.Info(toJSON(err))
}

// deleteUser deletes the user with the selected socialNumber.
// The return will be an error or Accepted.
func deleteUser(c *gin.Context) {
	log.Info("[Function] deleteUser")

	socialNumber := c.Request.Header.Get("socialNumber")

	logRequest(c.Request.Header, c.Request.Body)

	if socialNumber == ""{
		em := errorStatus.ErrorMissingHeaderValues()
		c.JSON(em.HTTPStatus, em)
		log.WithFields(log.Fields{"socialNumber": socialNumber}).Info(toJSON(em))
		return
	}
	err := service.DeleteUserService(socialNumber)
	if err.HTTPStatus != 200 {
		c.JSON(err.HTTPStatus, err)
	} else {
		c.JSON(err.HTTPStatus, "User Deleted.")
	}
	log.Info(toJSON(err))
}

// Api request to check if service is up
func healthCheck(c *gin.Context) {
	c.JSON(200, "ok")
}

// Init logs
func initLog(){
	// Set log format to JSON
	log.SetFormatter(&log.JSONFormatter{})
	// Will log anything that is info or above (warn, error, fatal, panic). Default.
	log.SetLevel(log.InfoLevel)
	log.Info("Initializing Log.")
}

// Transforms the object as a json object
func toJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func logRequest(header http.Header, obj interface{}){
	switch obj.(type) {
	case io.ReadCloser:
		_, obj = getStringByReadCloser(obj.(io.ReadCloser))
	}

	log.WithFields(log.Fields{
		"Header": header,
		"Body": obj,
	}).Info("[API REQUEST]")
}

func getStringByReadCloser(closer io.ReadCloser) (io.ReadCloser, string) {
	defer closer.Close()
	closerBytes, _ := ioutil.ReadAll(closer)
	bodyString := string(closerBytes)
	closer = ioutil.NopCloser(bytes.NewBuffer(closerBytes))
	return closer, bodyString
}