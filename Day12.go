/*******************************************************************************
Carlos Redondo
December 12th, 2017
*******************************************************************************/


package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)



var Still_data = true // While loop separator

// Slice with final list of connected values
var fincon []int

// Substitutes in string
var replacer = strings.NewReplacer(",", " ", " ", "")

// Removes the duplicates of a slice
// Returns the slice without duplicates

func removeDuplicates(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

var l1 int = 0 // Previous length


func main() {

    // Loops indefinitely as long as there are programs to connect

    // The zero is already by default
    fincon = append(fincon, 0)

    for Still_data == true{

      // Reads the file line by line
      file, err := os.Open("A12data.txt")
      // If file is not found
      if err != nil {
          log.Fatal(err)
      }
      defer file.Close()
      scanner := bufio.NewScanner(file)

      for scanner.Scan() {

        // Divides the line into first and second side of the pipe
        stringSlice := strings.Split(scanner.Text(), "<->")
        // Turns the first element into a number
        first_elem, poserr := strconv.Atoi(replacer.Replace(stringSlice[0]))

        if poserr != nil{

          fmt.Println(first_elem)
          fmt.Println(poserr)
          fmt.Println("Error message below")
          panic("Problem reading the number")
        }

        // Divides the second part
        section2 := replacer.Replace(stringSlice[1])
        section3 := strings.Fields(section2)

        // Loops through the already analyzed file
        for qq := 0; qq < len(fincon); qq++ {

          // No point if the value has already been added

          if fincon[qq] == first_elem{

            // Adds the element to the list if they are not there already
            for lol := 0; lol < len(section3); lol++ {

              secel, poserr := strconv.Atoi(section3[lol])

              if poserr != nil{

                panic("Error in number conversion")
              }

              fincon = append(fincon, secel)
              // Simplify the final result by removing duplicates
              fincon = removeDuplicates(fincon)

            }

            break
          }

        }

      }

      // Computes the length
      l2 := len(fincon)

      // Program ends when no new programs have been added
      if l1 == l2{

        fmt.Println("Programs connected: ", len(fincon))
        break
      }

      // The old becomes new
      l1 = l2
    }
}
