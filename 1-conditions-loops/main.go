package main

import (
  "fmt"
  "math/rand"
)

// Test conditional statements and loops with random integer
func main() {
  randomNumber := rand.Intn(101)
  var greater bool = randomNumber > 50 // can also write: 'greater := randomNumber > 50'
  school := "Holberton School"
  isHolbertonSchool := school == "Holberton School"
  beautifulWeather := true
  holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

  // Test if random number is > 50
  if greater {
    fmt.Printf("my random number is %v and is greater than 50\n", randomNumber)
  }  else {
    fmt.Printf("my random number is %v and is less than 50\n", randomNumber)
  }

  // Check if school is "Holberton School" or not
  if isHolbertonSchool {
    fmt.Println("I am a student of the Holberton School")
  }  else {
    fmt.Println("I am not a student of the Holberton School")
  }

  // Check how the weather is
  if beautifulWeather {
    fmt.Println("It's a beautiful weather!")
  } else {
    fmt.Println("The weather isn't so great today.")
  }

  // Iterate over slice containing Holberton founders
  for _, v := range holbertonFounders {
    fmt.Printf("%s is a founder\n", v)
  }
}
