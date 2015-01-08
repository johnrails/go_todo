package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"

)


func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEcoder(w).Encode(todos); err != nil{
    panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoID := vars["todoId"]
  fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
  var todo Todo
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048567))
  if err != nil{
    panic(err)
  }

  if err := r.Body.Close(); err !=nil {
    panic(err)
  }

  if err := json.Unmarshal(body, &todo); err != nil{
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(422) // unprocessable entity
    if err := json.NewEcoder(w).Encode(err); err != nil{
      panic(err)
    }
  }

  t := RepoCreateTodo(todo)
  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEcoder(w).Encode(t); err != nil{
    panic(err)
  }
}
