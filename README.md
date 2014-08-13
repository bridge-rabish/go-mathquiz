Simple math quiz
=============================================

It's a simple application, created on learning golang.

## Idea

The quiz will collect users information, and create a random set of 5 arithmetic quiz
to check how fast a user can do a arithmetic operator  in 10 seconds. 

# API

## User information collect, create arithmetic quiz

Request
```
POST /userinfo
name : "ABC",
email : "abc@abc.com"
```

Response
```
[{
    quiz_id : 'zxxwsds',
    question_id : 1,
    question : "15 * 32 = ",
    answers : [
        1235,
        1230,
        1230,
        1245
    ],
    attempted : "false"
},
..
..
]
```

## Save answer for a question
Request
```
POST /question
quiz_id : "zxxwsds"
question_id : 1
answer : 1
```
Response
```
{
    quiz_id : 'zxxwsds',
    question_id : 1,
    attempted : true
}
```

## Display score

Request 
```
GET /score/zxxwsds
```

Response
```
{
    quiz_id : 'zxxwsds',
    score : 50
}
```

### Display answers and wrong questions

Request 
```
GET /score/zxxwsds
```

Response
```
[{
    quiz_id : 'zxxwsds',
    question_id : 1,
    question : "15 * 32 = ",
    answers : [
        1235,
        1230,
        1230,
        1245
    ],
    correct : 1235,
    input : 1245,
    attempted : "true"
},
..
..
]
```

## Run application

```
$ go run server.go
```

# License

 MIT