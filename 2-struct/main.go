package main

import (
    "fmt"
    "time"
  )

type user struct {
  Name string `json:"name"`
  DOB string `json:"date_of_birth"`
  City string `json:"city"`
}

// print a Hello message to user
func (u *user) SayHello() {
    fmt.Printf("Hello %v\n", u.Name)
}

// calculate age of the user
func (u *user) CalculateAge() int {
  layout := "January 2, 2006"
  date, err := time.Parse(layout, u.DOB)

  if err != nil {
    fmt.Println("Error calculating age")
    return 0
  }

  age := time.Now().Year() - date.Year()  // begin age calculation

  // if we've not yet reached user's birthday this year, subtract a year from age calculation
  if (time.Now().Month() < date.Month() || (time.Now().Month() == date.Month() && time.Now().Day() < date.Day())) {
    age--
  }

  return age;
}

// print a statement about the user
func (u *user) PrintSentence() {
    age := u.CalculateAge();
    fmt.Printf("%v who was born in %v would be %d years old today\n", u.Name, u.City, age)
}

// Say hello to user and print a statement about them
func main() {
  u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
  u.SayHello();
  u.PrintSentence();
}
