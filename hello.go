package main

import (
	"fmt"
	"os"
    "io/ioutil"
	"strconv"
	"strings"
)

func main(){
	counter := 0
	empty := ""
	block := ""
	replace := ""
	row := 0
	
	//Checking if filename was given
	if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }
    data, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }
	
	//Took metadata from file
	var metadata string = string(data)[0:5]
	
	//Checking metadata's first number and if it's 1, taking second character as well
	//If first number is different than 1, moving forward
	height, err := strconv.Atoi(metadata[0:1])
	if height == 1{
	
		//Reading metadata to get characters for empty, block and replace
		height, err = strconv.Atoi(metadata[0:2])
		empty = metadata[2:3]
		block = metadata[3:4]
		replace = metadata[4:5]
		
		//Taking 6th to 40th characters from file and splitting them with \n to check what is boards lenght
		var lenght string = string(data)[6:40]
		splits := strings.Split(lenght, "\n")
		row = len(splits[0])
		
		//Setting counter to count characters from file, so that board will be precicely like in the file
		counter = 5
	}else{
		//Same as above
		empty = metadata[1:2]
		block = metadata[2:3]
		replace = metadata[3:4]
		
		counter = 4
		
		var lenght string = string(data)[5:40]
		splits := strings.Split(lenght, "\n")
		row = len(splits[0])
	}
	
	//Creating starter board from file into matrix
	originalBoard := make([][]string, height)
	for i := range originalBoard {
		originalBoard[i] = make([] string, row)
	}

	//Creating replica of the file board using counter
	for i := 0; i < height; i++{
		for j := 0; j<row; j++{
			counter++
			if j == row-1 {
				originalBoard[i][j] = string(data)[counter:counter+1]
				counter++
			}else{
				originalBoard[i][j] = string(data)[counter:counter+1]
			}
		}
	}
	
	//Making two new matrix, first to check values inside originalBoard, 
	//second to be matrix that will be printed out
	newBoard := make([][]int, height)
	for i := range newBoard {
		newBoard[i] = make([] int, row)
	}
	finalBoard := make([][]string, height)
	for i := range finalBoard {
		finalBoard[i] = make([] string, row)
	}
	
	//Assigning values to value boards first collumn and row and coping orginalBoard to finalBoard
	//Assigning 1 if matrix[i][j] is empty space and 0 to blocks
	for i := 0; i < height; i++{
		if strings.Contains(originalBoard[i][0], block){
			newBoard[i][0] = 0
			finalBoard[i][0] = originalBoard[i][0]
		}else{
			newBoard[i][0] = 1
			finalBoard[i][0] = originalBoard[i][0]
		}
	}
	for i := 0; i < row; i++{
		if strings.Contains(originalBoard[0][i], empty){
			newBoard[0][i] = 1
			finalBoard[0][i] = originalBoard[0][1]
		}else{
			newBoard[0][i] = 0
			finalBoard[0][i] = originalBoard[0][i]
		}
	}
	//Going through values and checking where is the first highest value find in whole board
	maxCount := newBoard[0][0]
	maxI := 0
	maxJ := 0
	
	//Going through matrix to set values to whole board
	for i := 1; i < height; i++{
		for j := 1; j<row; j++{
			if strings.Contains(originalBoard[i][j], block){
				newBoard[i][j] = 0
				finalBoard[i][j] = originalBoard[i][j]
			}else{
					//Every time matrix is empty space, checking previous values next to new one
					//And taking smallest value and increasing it by 1
					newBoard[i][j] = min(newBoard[i][j-1], newBoard[i-1][j], newBoard[i-1][j-1]) + 1
					finalBoard[i][j] = originalBoard[i][j]
					if maxCount < newBoard[i][j]{
						maxCount = newBoard[i][j]
						maxI = i
						maxJ = j
					}
			}
		}
	}
	
	//Going backwards from highest value to replace empty spaces with square
	for i := maxI; i > maxI - maxCount; i--{
		for j := maxJ; j > maxJ - maxCount; j--{
			finalBoard[i][j] = replace
		}
	}
	
	//Printing whole board when square is made
	for i := 0; i < height; i++{
		for j := 0; j<row; j++{
			if j == row-1{
				fmt.Println(finalBoard[i][j])
			}else{
				fmt.Print(finalBoard[i][j])
			}
		}
	}
}



func min(a, b, c int) int{
	if a < b && a < c{
		return a
	}else if b < a && b < c{
		return b
	}else if c < a && c < b{
		return c
	}else if a == b{
		return b
	}else{
		return c
	}
}
