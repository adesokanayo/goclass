package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
)

type myHandler struct {
}

func (now *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/css"
		}

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Mi amigo - " + http.StatusText(404)))
	}
}

func main() {

	http.Handle("/", new(myHandler))
	http.Handle("/api/getuser/:uid", new(getuser))
	http.Handle("/login", new(myHandler))
	http.ListenAndServe(":4040", nil)

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB session Established")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//collection sample record.
	/*

		session.DB(*mgo).C(Users).Insert(u)

		u:=&Users{
			username:  "adesokanayo",
			firstname: "Odunayo",
			lastname:  "Adesokan",
			Status:    1,
			Age:       "40",
		})
	*/

}

//Users struct exported
type Users struct {
	User      string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       string `json:"age"`
	Status    int    `json:"status"`
}

type getuser struct {
}

func (g *getuser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	log.Println(params)
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are get user %s", uid)

}
