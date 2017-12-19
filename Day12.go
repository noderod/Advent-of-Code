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


// Tells if an element is already inside a slice
func isin(lopez int, tocheck []int) bool {

  for lant := 0; lant < len(tocheck); lant++{

    if tocheck[lant] == lopez{

      return true
    }
  }
  return false
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



    // Counts all programs which are only connected to themselves
    // Although it adds more code, it significantly reduces runtime for large number of lines
    var lonelies []int

    // Reads the list, again
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
        panic("Problem reading the number")
      }

      // Computes the second number
      // Divides the second part
      section2 := replacer.Replace(stringSlice[1])
      section3 := strings.Fields(section2)

      // It only makes sense to check if there is only one element on each side
      if len(section3) == 1{
        secel, poserr := strconv.Atoi(section3[0])

        if poserr != nil{
          panic("Error reading number")
        }

        if secel == first_elem{
          lonelies = append(lonelies, first_elem)
        }
      }
    }

    fmt.Println("Programs only connected to themselves: ", len(lonelies))


    //##########################################################################
    // For the rest of part 2, we will loop through the entire file over and over
    // And check until all programs have been assigned to a group

    // Final list of groups, each group contains the program it is connected to
    // The alone points are their own thing

    // Creates a slice of slices
    var allgro [][]int

    // Adds the first group
    allgro = append(allgro, fincon)

    file_OR, err := os.Open("A12data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file_OR.Close()
    scanner_OR := bufio.NewScanner(file_OR)

    for scanner_OR.Scan() {

      // Reads the present number and checks if it has already been accounted for
      stringSlice := strings.Split(scanner_OR.Text(), "<->")
      first_elem, err_OR := strconv.Atoi(replacer.Replace(stringSlice[0]))

      fmt.Println("Now checking: #", first_elem)

      if err_OR != nil{
        panic(err_OR)
      }

      // Checks that it is not in the lonely group
      if isin(first_elem, lonelies) == true{
        continue
      }

      elem_accounted := false

      // Checks that it is not in any group so far
      for waltz := 0; waltz < len(allgro); waltz++{
        if isin(first_elem, allgro[waltz]) == true{
          elem_accounted = true
          break
        }
      }

      if elem_accounted == true{
        continue
      }

      // If all these conditions are met, then the point forms its own group
      var new_group []int

      new_group = append(new_group, first_elem)

      // Repeats the same algorithm as above to find all the points connected to it
      Still_data := true
      var PLEN int = 0


      for Still_data == true{

        file22, err22 := os.Open("A12data.txt")
        // If file is not found
        if err22 != nil {
            log.Fatal(err)
        }
        defer file22.Close()
        scanner22 := bufio.NewScanner(file22)

        for scanner22.Scan() {

          // Divides the line into first and second side of the pipe
          SLL22 := strings.Split(scanner22.Text(), "<->")
          // Turns the first element into a number
          FEL, errDD := strconv.Atoi(replacer.Replace(SLL22[0]))
          if errDD != nil{
            panic(errDD)
          }

          // Computes the second number
          // Divides the second part
          tion2 := replacer.Replace(SLL22[1])
          tion3 := strings.Fields(tion2)

          // Loops through the already analyzed file
          for dc := 0; dc < len(new_group); dc++ {

            if new_group[dc] == FEL{

              for lol := 0; lol < len(tion3); lol++ {

                secel, poserr := strconv.Atoi(tion3[lol])
                if poserr != nil{
                  panic("Error in number conversion")
                }

                new_group = append(new_group, secel)
                new_group = removeDuplicates(new_group)
              }
              break
            }
          }
        }

        NeLEN := len(new_group)

        if PLEN == NeLEN{
          break
        }

        PLEN = NeLEN
      }

      // Adds to the number of groups
      allgro = append(allgro, new_group)
    }

    fmt.Println("Groups: ", len(allgro)+len(lonelies))
}
