package errorStatus

/*
 * Project:     GolangRestfulApiExample
 * Created by:  joaopmgd@gmail.com | JP
 * Date:        01th November 2017
 */

import log "github.com/sirupsen/logrus"

//ErrorMessage Error message interface
type ErrorMessage struct {
	HTTPStatus   int    `json:"-"`
	Code         string `json:"code"`
	Title        string `json:"title"`
	Message      string `json:"message"`
}

func ErrorMissingBodyValues ()(ErrorMessage){
	err := ErrorMessage{
		HTTPStatus: 412,
		Code: "MISSING_BODY_VALUES.ERROR",
		Title: "Body",
		Message: "Body is not complete, missing values.",
	}
	log.Info(err)
	return err
}

func ErrorDatabaseConnection ()(ErrorMessage){
	err := ErrorMessage{
		HTTPStatus:500,
		Code: "BD_CONNECTION.ERROR",
		Title: "Database",
		Message: "Error connecting to the database.",
	}
	log.Info(err)
	return err
}

func ErrorMissingHeaderValues ()(ErrorMessage){
	err := ErrorMessage{
		HTTPStatus: 412,
		Code:       "MISSING_HEADER_VALUES.ERROR",
		Title:      "Header",
		Message:    "Header is not complete, missing values.",
	}
	log.Info(err)
	return err
}

func ErrorUserNotFound ()(ErrorMessage) {
	err := ErrorMessage{
		HTTPStatus: 404,
		Code:       "USER_NOT_FOUND.ERROR",
		Title:      "Missing User",
		Message:    "User was not found.",
	}
	log.Info(err)
	return err
}

func ErrorUserDoesNotExist ()(ErrorMessage) {
	err := ErrorMessage{HTTPStatus: 404,
		Code: "USER_DOES_NOT_EXISTS.ERROR",
		Title: "Missing User",
		Message: "The user was not found.",
	}
	log.Info(err)
	return err
}

func ErrorUserAlreadyExists ()(ErrorMessage) {
	err := ErrorMessage{
		HTTPStatus: 409,
		Code:       "USER_ALREADY_EXISTS.ERROR",
		Title:      "Existing User",
		Message:    "User already exist.",
	}
	log.Info(err)
	return err
}

func RequestIsOk ()(ErrorMessage) {
	err := ErrorMessage{
		HTTPStatus: 200,
		Code: "REQUEST_IS_OK",
	}
	log.Info(err)
	return err
}