// A tutorial from https://golang.org/doc/articles/wiki/

package main

import (
	"fmt"
	"io/ioutil"
)

// define data structures
type Page struct {
	Title string
	// using []byte rather than string
	// because it is the expected type by the io library
	Body []byte
}

// This is a method named save
// takes p, a pointer to Page.
// Takes no parameter and return a value of type error
// this function saves Page's Body to a text file.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// create new page
	p1 := &Page{Title: "TestPage", Body: []byte("Sample Page")}

	// save page into a text file
	p1.save()

	// load page
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
