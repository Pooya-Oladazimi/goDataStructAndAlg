package main

import (
  "fmt"
  "errors"
)


func main(){
  q := Newqueue[int](3)
  _, err := q.enqueue(10)
  checkError(err)
  _, err = q.enqueue(20)
  checkError(err)
  _, err = q.enqueue(30)
  checkError(err)
  element, deqerr := q.dequeue()
  checkError(deqerr)
  element, deqerr = q.dequeue()
  element, deqerr = q.dequeue()
  _, err = q.enqueue(390)
  fmt.Println(element)
  fmt.Println(q.q)
}


type queue[T any] struct{
  q []T
  front int
  rear int 
} 


func Newqueue[T any](capacity int) *queue[T]{
  return &queue[T]{
    q: make([]T, capacity),
    front: -1,
    rear: -1,
  }
}


func (q *queue[T]) isEmpty() (bool){
  if q.front == -1 && q.rear == -1{
    return true
  }
  return false
}


func (q *queue[T]) enqueue(element T) (T, error){
  if q.isEmpty(){
    q.rear = 0
    q.front = 0
    q.q[q.rear] = element
    return element,nil
  }
  q.rear = (q.rear + 1) % len(q.q)
  if q.rear == q.front{
    return element,errors.New("queue is full.")
  }
  q.q[q.rear] = element
  return element,nil
}


func (q *queue[T]) dequeue()(T, error){
  var element T
  if q.isEmpty(){
    return element,errors.New("queue is empty.")
  } 
  element = q.q[q.front]
  if q.front == q.rear{
    q.front = -1
    q.rear = -1
  }else{
    q.front = (q.front + 1) % len(q.q)
  }
  return element, nil
}


func checkError(err error){
  if err != nil{
    panic(err)
  }
}
