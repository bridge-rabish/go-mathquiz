package main

import (
    "fmt"
    "net/http"
    "io"
    "encoding/json"
    // "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
)

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

func userinfo( res http.ResponseWriter, req *http.Request ){

    name := req.FormValue("name" )
    email := req.FormValue("email" )

    u := &User{name, email}
    jsonString := ""
    encoded, err := json.Marshal(u)
    if err != nil{
        jsonString = "{\"error\": \"error\"}"
    } else {
        jsonString = string( encoded )
    }
    res.Header().Set(
        "Content-Type",
        "application/json" ,
    )
    io.WriteString(
        res,
        jsonString,
    )
}

func question( res http.ResponseWriter, req *http.Request ){
}

func score( res http.ResponseWriter, req *http.Request ){
}

func index( res http.ResponseWriter, req *http.Request ){
    res.Header().Set(
        "Content-Type",
        "text/html" ,
    )
    io.WriteString(
        res,
        "index page",
    )
}

func main() {
    http.HandleFunc("/", index )
    http.HandleFunc("/userinfo", userinfo )
    http.HandleFunc("/question", question )
    http.HandleFunc("/score", score )
    http.ListenAndServe(":9000", nil )
    fmt.Println("Server is listening to port 9000" );
}