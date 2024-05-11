package main

import (
  "fmt"
  "io/ioutil"
  "regexp"
)

// This checks wheather there is any error or not 
func check(err error){
  if err != nil{
    fmt.Println(err)
    panic(err)
  }
}


