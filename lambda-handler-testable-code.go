package main

import (
    "fmt"
    "reflect"
)

// this kind of structure gives the flexibility of testing each service without the use of additional mocking frameworks or 
// passing interfaces to functions and keep the method ( or "functions" if you want to call it) signatures clean


type X struct{
  y Y
}

func (x X) fx() {
  x.y.fy()
}

//lambda.start function
func lambdastart(f interface{}){
    // what lambda.start internally does // stripped out all the additional logic
   handler := reflect.ValueOf(f)
    var args []reflect.Value
    handler.Call(args) // passing empty list as this X.fx is not taking any additional params
}

type Y struct{
  z Z
}

func (y Y) fy(){
  y.z.fz()
}

type Z struct{
  val int
}

func (z Z) fz(){
 fmt.Println("Hello, playground : ", z.val)
}


func main() {
    fmt.Println("Hello, playground")
    
    x:= X{}
    x.y.z.val = 10
    
    lambdastart(x.fx) //lambda.start(handler)
}
