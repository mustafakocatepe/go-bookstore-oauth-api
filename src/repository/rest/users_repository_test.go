package rest

import (
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"os"
	"testing"
)

//This is the entry point of our test cases.
func TestMain(m *testing.M) {
	fmt.Println("about to start test cases... ")
	resty.New()
	os.Exit(m.Run())
}

//We are going to have a single test case for each return statement that we have.

func TestLoginUserTimeoutFromApi(t *testing.T) {

	repository := restUsersRepository{}
	response, err := repository.LoginUser("test@gmail.com", "password")
	println(response)
	println(err)

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}

//Notes : And since we don't want to make an actual API call to the API because we are in the test environment, we want to mock this response.
