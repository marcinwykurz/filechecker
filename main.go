package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "unicode"
)



func ReadFileLines(fpath string)  []string {
 // Reading every line of a file into an array of strings
    
    readFile, err := os.Open(fpath)
 
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
        os.Exit(1)
	}
 
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
 
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
 
	readFile.Close()

    if len(fileTextLines) == 0 {
        fmt.Println("The file is empty")
        os.Exit(1)
    }

    return fileTextLines
 
 
}


func OnlyDigits(str string) bool {
// Checking whether a string contains a digit
    result := true
	
    for _, char := range str {
		if !unicode.IsDigit(char) {
			result := false
            return result
		}                   
    }

    
    return result
    
}





func main() {
   

    //Checking whether at least one argument has been provided
    if len(os.Args) > 1 {
            fpath := os.Args[1]
            fileLines := ReadFileLines(fpath)
            colSum := make(map[int]int64)
            allNumbers := make([]string,0)


            //Going through all lines of the file , to check whether they cotnain only numbers and calculate the sum of each column
            for _, eachLine := range fileLines {
                counter := 1
                for _, eachElement := range strings.Split(eachLine, ",") {
                    if OnlyDigits(eachElement) {
                        if a, err := strconv.ParseInt(eachElement, 10, 64); err == nil {
                            colSum[counter] = colSum[counter] + a
                            allNumbers = append(allNumbers, eachElement)
                            }

                    } else {
                        fmt.Println("The file doesn't contain only digits. Please remove non digits characters")
                        os.Exit(1)
                    }
                    counter++



                }
            }
            

            //Checking whether all numbers in the file are unique
            for i := 0; i < len(allNumbers)-1; i++ {
                for j := i+1; j < len(allNumbers); j++ {
                    if allNumbers[i] == allNumbers[j] { 
                        fmt.Println("The file doesn't contain unique numbers")
                        os.Exit(1)
                    }

                }    
           
            }
                
            // Checking whehter the sum of the numbers in each column is the same
            numCol := len(colSum)
            if numCol > 1 {
                for i := 1; i < numCol; i++ {
                    if colSum[i] == colSum[i+1] {

                    } else { 
                    fmt.Println("The sum of the numbers in each column is not the same")
                    os.Exit(1) 
                     }
                    }

            } else { os.Exit(0)}        // if the file contains a single column then no point to the sum
            
            os.Exit(0)

        }   else {
            fmt.Println("Please provide a path to a file as an argument")
            os.Exit(1)
        }


}
