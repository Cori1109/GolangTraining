package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/cloud/datastore"
	"html/template"
	"net/http"
	"log"
	"fmt"
)

type User struct {
	Email    string
	UserName string
	Password string
}

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/login", Login)
	r.GET("/signup", Signup)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
}

func checkUserName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	var Possibility string
	json.NewDecoder(req.Body).Decode(&Possibility)
	q, err := datastore.NewQuery("Users").Filter("Username =", Possibility).Count(ctx)
	if err != nil {
		json.NewEncoder(res).Encode("false")
		return
	}
	if q >= 1 {
		json.NewEncoder(res).Encode("true")
	} else {
		json.NewEncoder(res).Encode("false")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
//	ctx := appengine.NewContext(req)
	NewUser := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: req.FormValue("password"),
	}
	log.Println(NewUser)
	fmt.Fprintln(res, NewUser.Email, NewUser.Password, NewUser.UserName)
//	json.NewDecoder(req.Body).Decode(&NewUser)
//	key := datastore.NewIncompleteKey(ctx, "Users", nil)
//	key, _ = datastore.Put(ctx, key, &NewUser)
}