package main

// @author - JackBekket (sergey ponomarev)
// this is my first application on golang

import (

  "fmt"
  "math/rand"
  "os"
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
  //rand.Seed(time.Now().UTC().UnixNano())

  // n_rolls - number of rolls to throw
  var input1 string;



  fmt.Print("Input number of rolls ");
  fmt.Scan(&input1);

  var input2 []string

  input2 = s.Split(input1,"d")


  var rolls_s string =input2[0]
  var sqrs_s string = input2[1]

  rolls, err := strconv.Atoi(rolls_s)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }

    sqrs, err1 := strconv.Atoi(sqrs_s)
      if err1 != nil {
          // handle error
          fmt.Println(err1)
          os.Exit(2)
      }



/*
  rolls  := strconv.Atoi(rolls_s)
  var sqrs int = strconv.Atoi(sqrs_s)
*/
  for i := 0; i < rolls; i++ {
    answer := random(1,sqrs)
    fmt.Println(answer)
  }


/*
  for i := 0; i < n_rolls; i++ {
    roll := random(1,20)
    fmt.Println(roll)
  }
*/


/*
  if err != nil {
  fmt.Println("Ошибка ввода: ", err)
}
*/

}
