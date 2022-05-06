package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//This is the entry point of our test cases.
func TestMain(m *testing.M) {
	fmt.Println("about to start test cases... ")
	//resty.New()
	os.Exit(m.Run())
}

//We are going to have a single test case for each return statement that we have.
func TestLoginUserTimeoutFromApi(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/login" {
			t.Errorf("Expected to request '/users/login', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,",first_name :"mustafa","last_name:"kocatepe","email":"m@gmail.com"}`))
	}))
	defer server.Close()

	baseUrl = server.URL
	repository := restUsersRepository{Client: server.Client()}
	response, err := repository.LoginUser("test@gmail.com", "password")
	if response == nil {
		t.Errorf("MSK")
	}
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
