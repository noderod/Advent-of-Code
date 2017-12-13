/**
Carlos Redondo
Started December 11th, 2017
**/


package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

// Final result Part 1, checksum
var Finco int = 0

// Final result Part 2, sum of dividers without rest
var Divcou int = 0



// Function to compute the min/max of slice
func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

// Values of the above function
var minar, maxar int
var lediv1, lediv2 int



func main() {
    file, err := os.Open("A2data.txt")
    // If file is not found
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        // Working
        // Divides the line in a slice of different numbers
        divar := strings.Fields(scanner.Text())

        // Slice with present values
        var tyre []int

        // Loops through the statement
        for qq := 0; qq < len(divar); qq++ {
          acval, poserr := strconv.Atoi(divar[qq])
          // Watches for errors
          if poserr == nil{
            tyre = append(tyre, acval)
          }
        }


        // Computes the min/max of this slice
        minar, maxar = MinMax(tyre)
        Finco += maxar - minar



        //  SECOND PART


        // Value to track the for loops
        tracker := 0

        // Loops through the slice to find the 2 elements which are divisible
        for mtmt := 0; mtmt < len(tyre); mtmt ++{

          // First element
          ELE := tyre[mtmt]

          for idid := 0; idid < len(tyre); idid ++{

            // Skips if both elements are the same
            if mtmt == idid{

              continue
            }

            // Exits if condition is met
            if tracker == 1{
              break
            }

            DES := tyre[idid]

            // Creates a second slice
            var cordiv []int
            // Assigns to array and finds the min/max
            cordiv = append(append(cordiv, ELE), DES)

            lediv1, lediv2 = MinMax(cordiv)

            // Condition is met if there is no residual
            if (lediv2 % lediv1) == 0{
              Divcou += lediv2/lediv1
              // Checker
              tracker = 1

            }

          }

        }

    }

    // Prints final result
    fmt.Println("Checksum: ", Finco)
    fmt.Println("Sum of divisors: ", Divcou)
}
