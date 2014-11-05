package main

import(
        "fmt" 
        "ioutil"
)

type Page struct{
  Title string
  Body []byte
}

//method save takes a pointer to a page as its receiver
//takes no parameters, returns value of type error
//save's page's body to a text file
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}


