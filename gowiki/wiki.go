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
  title, err := getTtile(w,r)
  if err != nil {
    return
  }
  body := r.FormValue("body")
  //convert formvalue to []byte before it will fit in page struct
  p := &Page{Title: title, Body: []byte(body)}
  err := p.save()
  if err != nil{
    http.Error(w, err.Error(),http.StatusInternalServerError)
    return
  }
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
  title, err := getTtile(w,r)
  if err != nil {
    return
  }
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w,r,"/view/"+title,http.StatusFound)
    return
  }
  renderTemplate(w,"view",p)
}

func editHandler(w http.ResponseWriter, r * http.Request){
  title, err := getTtile(w,r)
  if err != nil {
    return
  }
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w,r,"/edit/"+title,http.StatusFound)
    return
  }
  renderTemplate(w,"edit",p)
}

var templates = template.Must(template.ParseFiles("edit.html","view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8080",nil)
}