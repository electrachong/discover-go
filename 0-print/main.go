package main

import (
  "fmt"
  "time"
  )

func main () {
  fmt.Println("Hello, we are Holberton School")
  fmt.Printf("the date is %v\n", time.Now())
  fmt.Printf("the year is %v\n", time.Now().Year()) // print the year
  fmt.Printf("the month is %v\n", time.Now().Month()) // print the month
  fmt.Printf("the day is %v\n", time.Now().Day()) // print the day
}
