package main

import "fmt"

const sum = 2020

func main() {
	result := 0
	substractor := 0
    s := []int{1721, 979, 366, 299, 675, 1456}
    // fmt.Println(len(s), s)
	
	// for loop:
	// substract first element from '2020'
	// then remove first element and loop again until no elements left.
	
    for len(s) > 0 {
        x := s[0]      // get the 0 index element from slice
        s = s[1:]      // remove the 0 index element from slice
        fmt.Println("substractor: ", x) // print 0 index element
		z := sum -x 
		fmt.Println("result: ", z) // print 2020-x
        // fmt.Println(len(s), s)
		// fmt.Println("---")
		
		// catch value of substraction
		// check if it match with any element
		// exit on *first* match
		for i := range s {
    		if s[i] == z {
        		// Found
				fmt.Println("match: ", z)
				result = z
				substractor = x
				break
    		}
		}
		if z != 0 {
			break
		}
    }
	fmt.Println( "multiplicator #1 ", substractor )
	fmt.Println( "multiplicator #2 ", result )
	fmt.Println( substractor * result )
}

