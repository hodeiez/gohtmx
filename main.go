package main

import ( 
  "fmt"
  "net/http"
  "html/template"
  "log"

)
type Album struct {
  Band string
  Name string
}

func main(){
  fmt.Println("hello")
  h1 := func (w http.ResponseWriter, r *http.Request){


    tmpl := template.Must(template.ParseFiles("index.html"))
    content := map[string][]Album{
      "Albums":{
        {Band: "Nirvana", Name: "In utero"},
        {Band: "Love", Name: "Forever changes"},
      },
    }
    


    tmpl.Execute(w, content)
  }
  http.HandleFunc("/",h1)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
