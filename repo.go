package main

import "fmt"

var currentId int

var todos Todos

// Give us some seed data

func init() {
  RepoCreateTodo(Todo{Name: "Write presentation"})
  RepoCreateTodo(Todo{Name: "Host Meetup"})
}

func RepoFindTodo(id int) Todo {
  for _, t := range todos{
    if t.Id == id{
      return t
    }
  }
  // Return empty todo if not found
  return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
  currentId += 1
  t.Id = currentId
  todos = append(todos, t)
  return t
}

func RepoDestroyTodo(id int) error {
  for i, t := range todos{
    if t.Id == id{
      todos = append(todos[:i], todos[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}