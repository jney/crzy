package pkg

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer(t *testing.T) {
	handler := &api{}
	server := httptest.NewServer(handler)
	client := server.Client()

	request, _ := http.NewRequest("Get", server.URL, nil)
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Should not return %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf(
			"Status Code should be 200, current: %d",
			response.StatusCode,
		)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Should not return %v", err)
	}
	if string(body) != "version" {
		t.Errorf("Expecting version, get: %s",string(body))
	}
}