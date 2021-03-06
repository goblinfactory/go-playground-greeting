//go:build integration
// +build integration

package nethttp2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

var client30 = &http.Client{
	Timeout: 30 * time.Second,
}

// TestHTTPandJSONMarshalling ...
func TestHTTPandJSONMarshalling(t *testing.T) {

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-My-Client", "my study application")
	res, err := client30.Do(req)
	check(t, err)

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(fmt.Errorf("Unexpected status: got %v", res.Status))
	}
	fmt.Println("--- RESPONSE ---")
	fmt.Println(res.Header.Get("Content-Type"))

	// aaah, this is the first time I've seen inline class definitions! love it! keeping code clean.
	// the json attributes on the right help us refactor the dto without breaking the api.
	var data struct {
		UserID    int    `json:"userid"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	check(t, err)

	// the + below means that printf will include the property names
	json := fmt.Sprintf("%+v", data)
	assertEqual("{UserID:1 ID:1 Title:delectus aut autem Completed:false}", json)

	// and here we leave the + out and we just get the values
	json = fmt.Sprintf("%v", data)
	assertEqual("{1 1 delectus aut autem false}", json)
}

func must(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertEqual(a string, b string) {
	if a != b {
		log.Fatal(fmt.Errorf("`%s` does not equal `%s`", a, b))
	}
}

// TestMarshallingJSONResponseWithoutJSONAttributes ..
func TestMarshallingJSONResponseWithoutJSONAttributes(t *testing.T) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	res, err := client30.Do(req)
	check(t, err)
	defer res.Body.Close()

	// dont do this in production; i.e. make sure you include `json:"userid"` attributes as shown further up
	var data struct {
		UserID    int
		ID        int
		Title     string
		Completed bool
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	check(t, err)

	json := fmt.Sprintf("%+v", data)
	assertEqual("{UserID:1 ID:1 Title:delectus aut autem Completed:false}", json)

	json = fmt.Sprintf("%v", data)
	assertEqual("{1 1 delectus aut autem false}", json)
}
