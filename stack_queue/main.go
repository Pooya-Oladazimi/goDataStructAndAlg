package main

import (
  "fmt"
  "errors"
)


func main(){
  s := &Stack[string]{capacity: 3}
  el, err := s.push("f")
  checkError(err)
  _, err = s.push("r")
  checkError(err)
  _, err = s.push("g")
  checkError(err)
  el, err = s.pop()
  checkError(err)
  fmt.Println(s);
  fmt.Println(el);
}


type Stack[T any] struct{
  stack []T
  capacity int 
}


func (s *Stack[T]) push(element T)(T, error){
  if len(s.stack) == s.capacity{
    return element, errors.New("stack is full.")
  }
  s.stack = append(s.stack, element)
  return element,nil
}


func (s *Stack[T]) pop()(T, error){
  var nullT T 
  if len(s.stack) == 0{
    return nullT, errors.New("stack is empty.")
  }
  element := s.stack[len(s.stack) - 1]
  s.stack = s.stack[:len(s.stack) - 1]
  return element, nil 
}

func checkError(err error)bool{
  if(err != nil){
    panic(err)
  }
  return true
}
