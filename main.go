// Package plugindemo a demo plugin.
package bbwebhooks

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	OnMerge string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		OnMerge: "",
	}
}

type Bbwebhooks struct {
	next    http.Handler
	onMerge string
	name    string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Bbwebhooks{
		onMerge: config.OnMerge,
		next:    next,
		name:    name,
	}, nil
}

func (a *Bbwebhooks) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Convert the body to a string
	bodyString := string(body)

	// Print the body string to the console
	fmt.Println(bodyString)
	http.Error(rw, "Middleware ended", http.StatusBadRequest)
	// a.next.ServeHTTP(rw, req)
}
