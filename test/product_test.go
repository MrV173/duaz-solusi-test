package test

import (
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

func TestProductGetAll(t *testing.T) {
	app := fiber.New()

	mysql.DatabaseInit()
	database.RunMigration()

	app.Get("/products", handlers.ProductHandlerGetAll)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
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

	var products []models.Product
	if data, ok := response["data"].([]interface{}); ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(dataBytes, &products)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatal("expected data to be a slice of interfaces")
	}

	if len(products) == 0 {
		t.Fatal("expected at least one user, got none")
	}

	t.Logf("List of products: %+v", products)
}

func TestProductGetByID(t *testing.T) { 

	// to using this unit test please using this template code 
	// PRODUCT_ID=1 go test -v -run TestProductGetByID
	// run it on /your/path/test-duaz-soluzi/test

	app := fiber.New()
	mysql.DatabaseInit()
	database.RunMigration()

	app.Get("/product/:id", handlers.ProductHandlerGetByID)

	productIDStr := os.Getenv("PRODUCT_ID")
	if productIDStr == "" {
		t.Skip("PRODUCT_ID cannot null, skiping test")
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		t.Fatal("internal server error", err)
	}

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/product/%d", productID), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("product not found")
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

	var product models.Product
	if data, ok := response["data"].(map[string]interface{}); ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(dataBytes, &product)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatal("expected data to be a map[string]interface{}")
	}
	if product.ID != productID {
		t.Fatalf("expected user ID %d, got %d", productID, productID)
	}
	t.Logf("User details: %+v", product)
}