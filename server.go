package main

import (
    "fmt"
    "net/http"
    "io"
    "math/rand"
    "strconv"
    "time"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
//    Quizzes 
}

//
// Save user
//
func saveuser( u *User) User{
    session, err := mgo.Dial("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("goquiz").C("users")
    err = c.Insert(&u)
    if err != nil {
        fmt.Println("error saving user data into mongo");    
        panic(err)
    }

    result := User{}
    err = c.Find(bson.M{"name": u.Name }).One(&result)
    if err != nil {
        panic(err)
    }

    return result
}

func savequiz (q *Quiz){
    session, err := mgo.Dial("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("goquiz").C("quizzes")
    err = c.Insert(&q)
    if err != nil {
        fmt.Println("error saving quizzes data into mongo");    
        panic(err)
    }
}

//
// Handle post request of user params
//
func userinfo( res http.ResponseWriter, req *http.Request ){

    name := req.FormValue("name" )
    email := req.FormValue("email" )

    u := &User{name, email}

    saveuser(u);

    quizobj  := makequiz()

    encoded, err := json.Marshal( quizobj )

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

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

type Question struct {
    Question_id string `json:"question_id"`
    Question string `json:"question"`
    Answers [5]string `json:"answers"`
    Attempted bool `json:"attempted"`
}

type Quiz struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Questions [5]Question `json:"questions"`   
}

//
// Build quiz questions and answers
// Save quiz and answer's into mongo db
//
func makequiz() *Quiz{
    rand.Seed( time.Now().UTC().UnixNano())
    ques := [5]Question{}

    for i:= 0; i< 5 ; i++ {
        first := randInt(10,99)
        second := randInt(10,99)
        quetext := strconv.Itoa( first ) + " * " +  strconv.Itoa( second ) + " = ?"
        answers := [5]string{
            strconv.Itoa( first+1 * second ),
            strconv.Itoa( first+2 * second ),
            strconv.Itoa( first+3 * second ),
            strconv.Itoa( first+4 * second ),
            strconv.Itoa( first+5 * second ),
        }
        que := Question{ 
            strconv.Itoa( i+1 ),
            quetext ,
            answers,
            false,
        }
        ques[i] = que
    }    
    quizobj := &Quiz{
        "1",
        "Quiz 1",
        ques,
    }
    savequiz( quizobj )
    return quizobj
}

func getquestion( res http.ResponseWriter, req *http.Request ){
}

func score( res http.ResponseWriter, req *http.Request ){
}

//
// Render index page template
//
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
    http.HandleFunc("/question", getquestion )
    http.HandleFunc("/score", score )
    http.ListenAndServe(":9000", nil )
    fmt.Println("Server is listening to port 9000" );
}