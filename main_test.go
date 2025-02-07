package main

import (
        "testing"
         "fmt"
        "net/http"
        "net/http/httptest"
          "github.com/gin-gonic/gin"
       )


func TestGetMultiplicaByID(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodGet, "/v1/multiplica/4/5", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    // Check to see if the response was what you expected
    if w.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
}


func TestGetMultiplicaByIDZeros(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodGet, "/v1/multiplica/0/0", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    // Check to see if the response was what you expected
    if w.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
}


func TestGetMultiplicaByIDFailFIrstParameter(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodGet, "/v1/multiplica/GG/5", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    // Check to see if the response was what you expected
    if w.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
}


func TestGetMultiplicaByIDFailSecondParameter(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodGet, "/v1/multiplica/4/YYY", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    // Check to see if the response was what you expected
    if w.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
}


func TestGetMultiplicaByIDFailBothParameter(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodGet, "/v1/multiplica/TTT/YYY", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    // Check to see if the response was what you expected
    if w.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
}
