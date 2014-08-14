package main

import (
    "fmt"
    "net/http"
    "io"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

func saveUserinfo( u *User) User{
    session, err := mgo.Dial("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    fmt.Println("saving data into mongo");
    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("goquiz").C("users")
    err = c.Insert(&u)
    if err != nil {
        fmt.Println("error saving data into mongo");    
        panic(err)
    }

    result := User{}
    err = c.Find(bson.M{"name": u.Name }).One(&result)
    if err != nil {
        panic(err)
    }

    return result
}
func userinfo( res http.ResponseWriter, req *http.Request ){

    name := req.FormValue("name" )
    email := req.FormValue("email" )

    u := &User{name, email}

    saveUserinfo(u);
    encoded, err := json.Marshal( u )

    jsonString := ""
    if err == nil{
        jsonString = string( encoded )
    } else {
        jsonString = "{\"error\": \"error\"}"
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