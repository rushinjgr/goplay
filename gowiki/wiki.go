package main

import(
        "io/ioutil"
        "net/http"
        "html/template"
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

func saveHandler(w http.ResponseWriter, r *http.Request){
  title := r.URL.Path[len("/save/"):]
  body := r.FormValue("body")
  //convert formvalue to []byte before it will fit in page struct
  p := &Page{Title: title, Body: []byte(body)}
  p.save()
  http.Redirect(w,r,"/view/"+title,http.StatusFound)
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

//extracts page title from r.URL.Path
//load page title and serve up the page
func viewHandler(w http.ResponseWriter, r *http.Request){
  title := r.URL.Path[len("/view/"):]
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w,r,"/view/"+title,http.StatusFound)
    return
  }
  renderTemplate(w,"view",p)
}

func editHandler(w http.ResponseWriter, r * http.Request){
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w,r,"/edit/"+title,http.StatusFound)
    return
  }
  renderTemplate(w,"edit",p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
  t, _ := template.ParseFiles(tmpl + ".html")
  t.Execute(w,p)
}

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8080",nil)
}