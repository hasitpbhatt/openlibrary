package olclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
)

// OpenLibrary serves as an interface
type OpenLibrary struct {
	base            string
	client          *http.Client
	credentialsPath string
}

// NewOpenLibrary sets up
func NewOpenLibrary(readOnly bool) (*OpenLibrary, error) {

	cookies, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	ol := &OpenLibrary{
		base: "https://openlibrary.org/",
		client: &http.Client{
			Jar: cookies,
		},
		credentialsPath: ".credentials",
	}

	if !readOnly && ol.login() != nil {
		return nil, err
	}

	return ol, nil
}

// login attempts to login into the website
func (o *OpenLibrary) login() error {
	f, err := os.Open(o.credentialsPath)
	if err != nil {
		return err
	}
	defer f.Close()

	url := o.path("account", "login")
	resp, err := o.client.Post(url, "x-www-form-urlencoded", f)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Invalid status code: %s", err)
	}

	return nil
}

// queryTypes = "works", "isbn"
func (o *OpenLibrary) get(key, queryType string) (*Work, error) {
	url := o.path(queryType, key+".json")
	resp, err := o.client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	w := Work{}
	if err := json.Unmarshal(b, &w); err != nil {
		return nil, err
	}

	return &w, nil
}

// Get gets editions for given OLID
func (o *OpenLibrary) Get(olid string) (*Work, error) {
	return o.get(olid, "works")
}

// GetByISBN gets editions for given OLID
func (o *OpenLibrary) GetByISBN(isbn string) (*Work, error) {
	return o.get(isbn, "isbn")
}

// Path returns absolute path
func (o *OpenLibrary) path(paths ...string) string {
	return o.base + filepath.Join(paths...)
}
