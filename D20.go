package main

// @author - JackBekket (sergey ponomarev)
// this is my first application on golang

import (

  "fmt"
  "math/rand"
//  "os"
  "time"
  "strconv"

)

import s "strings"



// Version 1 - how much dice(d20) we need to throw?
// Version 2 - how much rolls by dice with how much squares we need to throw?

func random(min, max int) int {
//  rand.Seed(time.Now().Unix())
  rand.Seed(time.Now().UTC().UnixNano())
  return rand.Intn((max - min)+1) + min
}





func main()  {





  var help string

  help = `Input value in format xdy when x is how much rolls you want
  and y is how much squares you need to have on all dices.
  Mechanics is worked like this:
"2d20" means that I want to throw 2 dices and I want both dices have 20 squares
  (from "1" to "20" when 20 is critical success and 1 is a critical failure)
   `



  var input1 string;



  fmt.Print(help);

  for{

    Begin:

  fmt.Scan(&input1);

  var input2 []string

  input2 = s.Split(input1,"d")


  var rolls_s string =input2[0]
  var sqrs_s string = input2[1]

  rolls, err := strconv.Atoi(rolls_s)
    if err != nil {
        // handle error
        fmt.Println(err)
        goto Begin

    }

    sqrs, err1 := strconv.Atoi(sqrs_s)
      if err1 != nil {
          // handle error
          fmt.Println(err1)
          goto Begin
      }




  for i := 0; i < rolls; i++ {
    answer := random(1,sqrs)
    fmt.Println(answer)
  }

}

}
