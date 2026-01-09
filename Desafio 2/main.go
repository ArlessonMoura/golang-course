package main

import "fmt"


func multiplesOfThreeAndFive(limit int) { 
  for i := 1; i <= limit; i++ {
    if i%3 == 0 && i%5 == 0 {
      fmt.Println("pinpan")
    } else if i%3 == 0 {
      fmt.Println("pin")
    } else if i%5 == 0 {
      fmt.Println("pan")
    } else {
      fmt.Println(i)
    }
  }
}

func main() {
	multiplesOfThreeAndFive(100)
}