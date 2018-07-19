package main

// @author - JackBekket (sergey ponomarev)
// this is my first application on golang

import (

  "fmt"
  "math/rand"
//  "os"
  "time"
)



// Version 1 - how much dice(d20) we need to throw?

func random(min, max int) int {
//  rand.Seed(time.Now().Unix())
  rand.Seed(time.Now().UTC().UnixNano())
  return rand.Intn((max - min)+1) + min
}




func main()  {
  //rand.Seed(time.Now().UTC().UnixNano())

  // n_rolls - number of rolls to throw
  var n_rolls int;

  fmt.Print("Input number of rolls ");
  fmt.Scan(&n_rolls);

  for i := 0; i < n_rolls; i++ {
    roll := random(1,20)
    fmt.Println(roll)
  }

/*
  if err != nil {
  fmt.Println("Ошибка ввода: ", err)
}
*/

}
