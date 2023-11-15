package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"test-duaz-solusi/database"
	"test-duaz-solusi/handlers"
	"test-duaz-solusi/models"
	"test-duaz-solusi/pkg/mysql"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func getRequestBody(objek interface{}) *bytes.Buffer {
	body, _ := json.Marshal(objek)
	return bytes.NewBuffer(body)
}

func TestUserGetAll(t *testing.T) {
	app := fiber.New()

	mysql.DatabaseInit()
	database.RunMigration()

	app.Get("/users", handlers.UserHandlerGetAll)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		t.Fatal(err)
	}

	expectedMessage := "success"
	if response["message"] != expectedMessage {
		t.Fatalf("expected message %s, got %s", expectedMessage, response["message"])
	}

	if response["data"] == nil {
		t.Fatalf("expected data to be nil, got %+v", response["data"])
	}

	var users []models.User
	if data, ok := response["data"].([]interface{}); ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(dataBytes, &users)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatal("expected data to be a slice of interfaces")
	}

	if len(users) == 0 {
    	t.Fatal("expected at least one user, got none")
	}

	t.Logf("List of users: %+v", users)
}

func TestUserGetByID(t *testing.T) { 

	// to using this unit test please using this template code 
	// USER_ID=1 go test -v -run TestUserGetByID
	// run it on /your/path/test-duaz-soluzi/test

	app := fiber.New()
	mysql.DatabaseInit()
	database.RunMigration()

	app.Get("/user/:id", handlers.UserHandlerGetByID)

	userIDStr := os.Getenv("USER_ID")
	if userIDStr == "" {
		t.Skip("USER_ID cannot null, skiping test")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		t.Fatal("internal server error", err)
	}

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/%d", userID), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("user not found")
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	expectedMessage := "success"
	if response["message"] != expectedMessage {
		t.Fatalf("expected message %s, got %s", expectedMessage, response["message"])
	}

	var user models.User
	if data, ok := response["data"].(map[string]interface{}); ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(dataBytes, &user)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatal("expected data to be a map[string]interface{}")
	}
	if user.ID != userID {
		t.Fatalf("expected user ID %d, got %d", userID, user.ID)
	}
	t.Logf("User details: %+v", user)
}

func TestUserCreate(t *testing.T) {

	// to using this unit test please using this template code 
	// USER_EMAIL=JHONATHAN@GMAIL.COM USER_PASSWORD=123456 USER_FULLNAME=JHONATAN USER_GENDER=MALE USER_PHONE=0812321321321 go test -v -run TestUserCreate
	// run it on /your/path/test-duaz-soluzi/test

	app := fiber.New()
	mysql.DatabaseInit()
	database.RunMigration()
	app.Post("/user", handlers.UserHandlerCreate)


	userEmail := os.Getenv("USER_EMAIL")
	userPassword := os.Getenv("USER_PASSWORD")
	userFullName := os.Getenv("USER_FULLNAME")
	userGender := os.Getenv("USER_GENDER")
	userPhone := os.Getenv("USER_PHONE")
		
	newUser := models.User {
		Email: userEmail,
		Password: userPassword,
		FullName: userFullName,
		Gender: userGender,
		Phone: userPhone,
	}

	reqBody := getRequestBody(newUser)

	req := httptest.NewRequest(http.MethodPost, "/user", reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusCreated {
		t.Fatalf("expected status code %d, got %d", fiber.StatusCreated, resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	excpectedMessage := "success"
	if response["message"] != excpectedMessage {
		t.Fatal("User created failed")
	}

	var user models.User
	if data, ok := response["data"].(map[string]interface{}); ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}

		err = json.Unmarshal(dataBytes, &user)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatal("internal server error")
	}

	if user.Email != newUser.Email || user.Password != newUser.Password || user.FullName != newUser.FullName || 
		user.Gender != newUser.Gender || user.Phone != newUser.Phone {
			t.Logf("Expected user: %+v", newUser)
			t.Logf("Created user: %+v", user)
		}

		t.Logf("Created user success : %+v", user)
}

func TestUserDeleteByID(t *testing.T) {
	
	// to using this unit test please using this template code 
	// USER_ID=1 go test -v -run TestUserDeleteByID
	// run it on /your/path/test-duaz-soluzi/test

	app := fiber.New()
	mysql.DatabaseInit()
	database.RunMigration()

	app.Get("/user/:id", handlers.UserHandlerGetByID)
	app.Delete("/user/:id", handlers.UserDeleteById)  

	userIDStr := os.Getenv("USER_ID")
	if userIDStr == "" {
		t.Skip("USER_ID not provided, skipping test")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		t.Fatalf("failed to parse USER_ID: %s", err)
	}

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/%d", userID), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	expectedMessage := "success"
	if response["message"] != expectedMessage {
		t.Fatalf("expected message %s, got %s", expectedMessage, response["message"])
	}

	var user models.User
	if data, ok := response["data"].(map[string]interface{}); ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(dataBytes, &user)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatal("expected data to be a map[string]interface{}")
	}

	if user.ID != userID {
		t.Fatalf("expected user ID %d, got %d", userID, user.ID)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%d", userID), nil)
	deleteResp, err := app.Test(deleteReq)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteResp.Body.Close()

	// Check if the user was deleted successfully
	if deleteResp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, deleteResp.StatusCode)
	}

	// Decode the response body for delete operation
	var deleteResponse map[string]interface{}
	err = json.NewDecoder(deleteResp.Body).Decode(&deleteResponse)
	if err != nil {
		t.Fatal(err)
	}

	expectedDeleteMessage := "success"
	if deleteResponse["message"] != expectedDeleteMessage {
		t.Fatalf("expected message %s, got %s", expectedDeleteMessage, deleteResponse["message"])
	}

	// Output the deleted user details to the testing logs
	t.Logf("Deleted user details: %+v", user)
}

func TestUserUpdateByID(t *testing.T) {

	// to using this unit test please using this template code 
	// USER_ID=2 USER_FULLNAME="Jackie" USER_EMAIL="jackie@example.com" USER_PASSWORD="123456" USER_GENDER="male" USER_PHONE="123456789" go test -v -run TestUserUpdateByID
	// run it on /your/path/test-duaz-soluzi/test

	app := fiber.New()
	mysql.DatabaseInit()
	database.RunMigration()
	app.Put("/user/:id", handlers.UserHandlerUpdate)
	app.Get("/user/:id", handlers.UserHandlerGetByID) 

	userIDStr := os.Getenv("USER_ID")
	if userIDStr == "" {
		t.Skip("USER_ID not provided, skipping test")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		t.Fatal("internal server error", err)
	}

	reqGet := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/%d", userID), nil)
	respGet, err := app.Test(reqGet)
	if err != nil {
		t.Fatal(err)
	}
	defer respGet.Body.Close()

	if respGet.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, respGet.StatusCode)
	}

	var userBeforeUpdate models.User
	err = json.NewDecoder(respGet.Body).Decode(&userBeforeUpdate)
	if err != nil {
		t.Fatal(err)
	}

	userUpdate := map[string]interface{}{
		"Email":    os.Getenv("USER_EMAIL"),
		"Password": os.Getenv("USER_PASSWORD"),
		"Fullname":     os.Getenv("USER_FULLNAME"),
		"Gender":   os.Getenv("USER_GENDER"),
		"Phone":    os.Getenv("USER_PHONE"),
	}

	reqBody, err := json.Marshal(userUpdate)
	if err != nil {
		t.Fatal(err)
	}

	reqUpdate := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/user/%d", userID), bytes.NewBuffer(reqBody))
	reqUpdate.Header.Set("Content-Type", "application/json")

	respUpdate, err := app.Test(reqUpdate)
	if err != nil {
		t.Fatal(err)
	}
	defer respUpdate.Body.Close()

	if respUpdate.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, respUpdate.StatusCode)
	}

	reqGetUpdated := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/%d", userID), nil)
	respGetUpdated, err := app.Test(reqGetUpdated)
	if err != nil {
		t.Fatal(err)
	}
	defer respGetUpdated.Body.Close()

	if respGetUpdated.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, respGetUpdated.StatusCode)
	}

	var user models.User
	err = json.NewDecoder(respGetUpdated.Body).Decode(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.Email != userUpdate["Email"] || user.Password != userUpdate["Password"] || user.FullName != userUpdate["FullName"] || 
		user.Gender != userUpdate["Gender"] || user.Phone != userUpdate["Phone"] {
			t.Logf("Expected user: %+v", userUpdate)
			t.Logf("Created user: %+v", user)
		}

		t.Logf("Created user success : %+v", user)
}