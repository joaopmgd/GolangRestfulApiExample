package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"GolangRestfulApiExample/mongoDB"
)

/*
 * Project:     GolangRestfulApiExample
 * Created by:  joao.goncalves@cielo.com.br | joao
 * Date:        06th November 2017
 */

var user_01 = mongoDB.User{
	Name: "Test_User_Name_01",
	Age: "21",
	Phones: []string{"123456789", "123456789"},
	SocialNumber: "12345-01",
	Address: address_user_01,
}

var address_user_01 = mongoDB.Address{
	StreetName: "Test_User_Street_01",
	StreetNumber: "123",
	ZipCode: "12345678",
	Country: "Brazil",
	State: "SP",
}

var user_02 = mongoDB.User{
	Name: "Test_User_Name_02",
	Age: "22",
	Phones: []string{"223456789", "223456789"},
	SocialNumber: "12345-02",
	Address: address_user_02,
}

var address_user_02 = mongoDB.Address{
	StreetName: "Test_User_Street_02",
	StreetNumber: "223",
	ZipCode: "22345678",
	Country: "United States",
	State: "CA",
}

var user_03 = mongoDB.User{
	Name: "Test_User_Name_03",
	Age: "23",
	Phones: []string{"323456789", "323456789"},
	SocialNumber: "12345-03",
	Address: address_user_03,
}

var address_user_03 = mongoDB.Address{
	StreetName: "Test_User_Street_03",
	StreetNumber: "323",
	ZipCode: "32345678",
	Country: "Brazil",
	State: "RJ",
}

var user_04 = mongoDB.User{
	Name: "Test_User_Name_04",
	Age: "24",
	Phones: []string{"423456789", "423456789"},
	SocialNumber: "12345-04",
	Address: address_user_04,
}

var address_user_04 = mongoDB.Address{
	StreetName: "Test_User_Street_04",
	StreetNumber: "423",
	ZipCode: "42345678",
	Country: "Brazil",
	State: "PR",
}

var new_name_user_01 = "Test_User_Name_01"
var fake_social_number = "00000-00"

// 1 - Try to insert 4 new unique users
func TestInsertNewUserService (t *testing.T){
	err := PostNewUserService(user_01)
	assert.Equal(t, 200, err.HTTPStatus)
	err = PostNewUserService(user_02)
	assert.Equal(t, 200, err.HTTPStatus)
	err = PostNewUserService(user_03)
	assert.Equal(t, 200, err.HTTPStatus)
	err = PostNewUserService(user_04)
	assert.Equal(t, 200, err.HTTPStatus)
}

// 2 - Try to insert a user that already exists
func TestInsertNewUserServiceAlreadyExists (t *testing.T){
	err := PostNewUserService(user_01)
	assert.Equal(t, 409, err.HTTPStatus)
}

// 3 - Get and validate all 4 users
func TestGetUserService (t *testing.T){
	user1, err := GetUserService(user_01.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	assert.Equal(t, user_01, user1)

	user2, err := GetUserService(user_01.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	assert.Equal(t, user_01, user2)

	user3, err := GetUserService(user_01.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	assert.Equal(t, user_01, user3)

	user4, err := GetUserService(user_01.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	assert.Equal(t, user_01, user4)
}

// 4 - Update a user name
func TestPatchUserService (t *testing.T){
	err := PatchUserService(user_01.SocialNumber, mongoDB.User{Name:new_name_user_01})
	assert.Equal(t, 200, err.HTTPStatus)
}

// 5 - Try to update a user that doesnt exist
func TestPatchUserServiceDoesntExist (t *testing.T){
	err := PatchUserService(fake_social_number, mongoDB.User{Name:new_name_user_01})
	assert.Equal(t, 404, err.HTTPStatus)
}

// 6 - Verifies if the user was updated
func TestGetUserServicePatched (t *testing.T){
	user1, err := GetUserService(user_01.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	assert.Equal(t, user1.Name, new_name_user_01)
}

// 7 - Try to get a user that does not exist
func TestGetUserServiceDoesntExist (t *testing.T){
	_, err := GetUserService("")
	assert.Equal(t, 404, err.HTTPStatus)
}

// 8 - Try to delete 4 users
func TestDeleteUserService (t *testing.T){
	err := DeleteUserService(user_01.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	err = DeleteUserService(user_02.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	err = DeleteUserService(user_03.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
	err = DeleteUserService(user_04.SocialNumber)
	assert.Equal(t, 200, err.HTTPStatus)
}

// 9 - Try to a user that does not exists
func TestDeleteUserServiceDoesntExist (t *testing.T){
	err := DeleteUserService(user_01.SocialNumber)
	assert.Equal(t, 404, err.HTTPStatus)
}