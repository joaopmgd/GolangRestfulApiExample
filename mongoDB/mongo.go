package mongoDB

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	log "github.com/sirupsen/logrus"

	"GolangRestfulApiExample/errorStatus"
)

/*
 * Project:     GolangRestfulApiExample
 * Created by:  joaopmgd@gmail.com | JP
 * Date:        01th November 2017
 */

type DataStore struct {
	Session *mgo.Session
}


type User struct {
	Name			string			`json:"name" bson:"name"`
	Age				string			`json:"age" bson:"age"`
	Phones			[]string		`json:"phones" bson:"phones"`
	SocialNumber	string			`json:"socialNumber" bson:"socialNumber"`
	Address			Address			`json:"address" bson:"address"`
}

type Address struct {
	StreetName		string			`json:"streetName" bson:"streetName"`
	StreetNumber	string			`json:"streetNumber" bson:"streetNumber"`
	ZipCode			string			`json:"zipCode" bson:"zipCode"`
	Country			string			`json:"country" bson:"country"`
	State			string			`json:"state" bson:"state"`
}

// Create a connection to the mongoDb database
func GetConnection () (DataStore, errorStatus.ErrorMessage){
	log.Info("[Function] GetConnection")
	//session, err := mgo.Dial("34.234.194.15")
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Info("Error trying to fetch help information from the database. URI: localhost.")
		return DataStore{}, errorStatus.ErrorDatabaseConnection()
	}
	log.Info("Database session established.")
	return DataStore{session}, errorStatus.RequestIsOk()
}

// Close the connection to the mongoDb database
func (dataStore DataStore) CloseConnection (){
	log.Info("[Function] CloseConnection")
	dataStore.Session.Close()
	log.Info("Database session closed.")
}

// PostNewUserMongo will insert a new user in the mongoDb database.
// The parameter is the user with the validated fields.
// The return will be an ErrorMessage with the status of the operation, success or the HttpStatus.
func (dataStore DataStore) PostNewUserMongo (user User)(errorStatus.ErrorMessage){
	log.WithFields(log.Fields{
		"user": user,
	}).Info("[Function] PostNewUserMongo")
	session := dataStore.Session.Clone()
	defer session.Close()
	_, userError := dataStore.GetUserMongo(user.SocialNumber)
	if userError.HTTPStatus == 200 {
		return errorStatus.ErrorUserAlreadyExists()
	}
	c := session.DB("user").C("userData")
	err := c.Insert(user)
	if err != nil {
		em := errorStatus.ErrorDatabaseConnection()
		return em
	}
	return errorStatus.RequestIsOk()
}

// GetUserMongo recovers the complete user object with all its data.
// The parameters is just the socialNumber.
// The return will be the complete user type found in the mongoDb with the request status in the errorMessage.
func (dataStore DataStore) GetUserMongo (socialNumber string)(User, errorStatus.ErrorMessage){
	log.WithFields(log.Fields{
		"socialNumber": socialNumber,
	}).Info("[Function] GetUserMongo")
	session := dataStore.Session.Clone()
	defer session.Close()
	var user User
	c := session.DB("user").C("userData")
	err := c.Find(bson.M{"socialNumber": socialNumber}).One(&user)
	if err != nil {
		em := errorStatus.ErrorUserNotFound()
		return User{}, em
	}
	return user, errorStatus.RequestIsOk()
}

// UpdateUserMongo updates the user data.
// The parameters are the new data or the updated one.
// The return will be the status of the request in the errorMessage.
func (dataStore DataStore) UpdateUserMongo (socialNumber string, user User)(errorStatus.ErrorMessage) {
	log.WithFields(log.Fields{
		"socialNumber": socialNumber,
		"user":         user,
	}).Info("[Function] UpdateUserMongo")
	session := dataStore.Session.Clone()
	defer session.Close()
	updatedUser, userError := dataStore.GetUserMongo(socialNumber)
	if userError.HTTPStatus != 200 {
		return userError
	}
	if user.Name != "" {updatedUser.Name = user.Name}
	if user.SocialNumber != "" {updatedUser.SocialNumber = user.SocialNumber}
	if user.Age != "" {updatedUser.Age = user.Age}
	if user.Phones != nil {updatedUser.Phones = user.Phones}
	if user.Address.Country != "" {updatedUser.Address.Country = user.Address.Country}
	if user.Address.State != "" {updatedUser.Name = user.Name}
	if user.Address.StreetName != "" {updatedUser.Address.StreetName = user.Address.StreetName}
	if user.Address.ZipCode != "" {updatedUser.Address.ZipCode = user.Address.ZipCode}
	if user.Address.StreetNumber != "" {updatedUser.Address.StreetNumber = user.Address.StreetNumber}
	c := session.DB("user").C("userData")
	err := c.Update(bson.M{"socialNumber": socialNumber}, updatedUser)
	if err != nil {
		em := errorStatus.ErrorDatabaseConnection()
		return em
	}
	return errorStatus.RequestIsOk()
}

// DeleteUserMongo deletes the selected user from the mongoDb database.
// The parameter is the socialNumber.
// The return is the status of the request in the errorMessage.
func (dataStore DataStore) DeleteUserMongo (socialNumber string)(errorStatus.ErrorMessage){
	log.WithFields(log.Fields{
		"socialNumber": socialNumber,
	}).Info("[Function] DeleteUserMongo")
	session := dataStore.Session.Clone()
	defer session.Close()
	c := session.DB("user").C("userData")
	info, err := c.RemoveAll(bson.M{"socialNumber": socialNumber})
	if info.Matched == 0{
		return errorStatus.ErrorUserDoesNotExist()
	}
	if err != nil {
		em := errorStatus.ErrorDatabaseConnection()
		return em
	}
	return errorStatus.RequestIsOk()
}