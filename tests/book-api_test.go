package test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	resp, err := http.Get("http://localhost:4040/api/v1/book")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading body : %v", err)
	}

	fmt.Printf("Body : %s\n", body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestHelloFunc(t *testing.T) {
	resp, err := http.Get("http://localhost:4040/")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading body : %v", err)
	}

	fmt.Printf("Body : %s\n", body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
