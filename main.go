package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// /-/css/style.css
	// / index
	http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/-/", fileServerHandler)

	//----- Block Dir -----
	//http.Handle("/-/", http.StripPrefix("/-", http.FileServer(http.Dir("public"))))
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("public")})))

	http.ListenAndServe(":8080", nil)
}

/*
func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	//h := http.FileServer(http.Dir("public"))
	//http.StripPrefix("/-", h).ServeHTTP(w, r)
}
*/

//******** Override Dir ********
//ถ้า Path ที่เข้าเป็น Dir [folder] จะไม่ให้เข้า
type noDir struct {
	http.Dir
}

func (d noDir) Open(name string) (http.File, error) {
	f, err := d.Dir.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}

//******** Override Dir ********

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//t, err := template.ParseFiles("index.tmpl")
	t, err := template.ParseFiles("index.tmpl")
	if err != nil { //ถ้ามี Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil) //Check Error
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
			w.Write([]byte(`
			<!doctype html>
			<title>Static Web Server</title>
			<link href=/-/css/style.css rel=stylesheet>
			<p class=red>
				Static web server.
			</p>
		`))
	*/
}
