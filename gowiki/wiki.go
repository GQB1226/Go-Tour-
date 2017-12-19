package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	//slice
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi there,I love %s!", r.URL.Path[1:])
}

/*
 *  不使用template解析
 *
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	p, _ := loadPage(title)
	fmt.Fprintf(w, ""+
		"<h1>%s</h1>"+
		"<div>%s<div>",
		p.Title, p.Body)
}
func editHandler(w http.ResponseWriter,r *http.Request){
	title:=r.URL.Path[len("/edit/"):]
	p,err:=loadPage(title)
	if err!=nil{
		p=&Page{Title:title}
	}
	fmt.Fprintf(w,"<h1>Editing %s</h1>" +
		"<form action=\"/save/%s\" method=\"POST\">" +
			"<textarea name=\"body\">%s</textarea><br>" +
				"<input type=\"submit\" value=\"Save\">" +
					"</form>",p.Title,p.Title,p.Body)

}
*/

func render_template(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err !=nil{
		http.Redirect(w,r,"/edit/"+title,http.StatusNotFound)
		return
	}
	render_template(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	render_template(w, "edit", p)
}

func saveHandler(w http.ResponseWriter,r *http.Request)  {
	title:=r.URL.Path[len("/save/"):]
	body:=r.FormValue("body")
	p:=&Page{Title:title,Body:[]byte(body)}
	p.save()
	http.Redirect(w,r,"/view/"+title,http.StatusFound)

}

func main() {

	p1 := &Page{Title: "testPage", Body: []byte("this is a test!")}
	p1.save()
	//p2,_:=loadPage("testPage")
	//fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/",saveHandler)
	http.ListenAndServe(":8080", nil)
}
