package service

import (
	"GolangRestfulApiExample/mongoDB"
	"GolangRestfulApiExample/errorStatus"

	log "github.com/sirupsen/logrus"
)

/*
 * Project:     GolangRestfulApiExample
 * Created by:  joaopmgd@gmail.com | JP
 * Date:        01th November 2017
 */


/*
 * All of this functions are creating a connection to the database and sending it to the request in mongoDb
 * All of the messages are formatted and returned to the main with its error status.
 */

func PostNewUserService (user mongoDB.User)(errorStatus.ErrorMessage){
	dataStore, err := mongoDB.GetConnection()
	if err.HTTPStatus != 200 {
		log.Info("Database connection error. URI: localhost.")
		return err
	}
	defer dataStore.CloseConnection()
	return dataStore.PostNewUserMongo(user)
}

func GetUserService (socialNumber string)(mongoDB.User, errorStatus.ErrorMessage){
	dataStore, err := mongoDB.GetConnection()
	if err.HTTPStatus != 200 {
		log.Info("Database connection error. URI: localhost.")
		return mongoDB.User{}, err
	}
	defer dataStore.CloseConnection()
	return dataStore.GetUserMongo(socialNumber)
}

func UpdateUserService (socialNumber string, user mongoDB.User)(errorStatus.ErrorMessage){
	dataStore, err := mongoDB.GetConnection()
	if err.HTTPStatus != 200 {
		log.Info("Database connection error. URI: localhost.")
		return err
	}
	defer dataStore.CloseConnection()
	return dataStore.UpdateUserMongo(socialNumber, user)
}

func DeleteUserService (socialNumber string)(errorStatus.ErrorMessage){
	dataStore, err := mongoDB.GetConnection()
	if err.HTTPStatus != 200 {
		log.Info("Database connection error. URI: localhost.")
		return err
	}
	defer dataStore.CloseConnection()
	return dataStore.DeleteUserMongo(socialNumber)
}