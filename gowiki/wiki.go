package main

import(
        "fmt" 
        "io/ioutil"
)

type Page struct{
  Title string
  Body []byte
}

//method save takes a pointer to a page as its receiver
//takes no parameters, returns value of type error
//save's page's body to a text file
//uses Title as file name
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

//constructs file name from the title parameter
//reads file's contents into a new variable body
//returns a pointer to a page literal constructed with the proper title and body values
//if error is nil, then page loaded successfully
func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil{
    return nil, err
  }
  return &Page{Title: title, Body: body},nil
}

func main() {
  p1 := &Page{Title: "TestPage",Body: []byte("This is a sample Page.")}
  p1.save()
  p2, _ := loadPage("TestPage")
  fmt.Println(string(p2.Body))
}