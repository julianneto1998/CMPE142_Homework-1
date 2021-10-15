package main

//Imports you may use: "bufio", "fmt", "os", "strings"
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* This assignment is to help familiarize yourself with Golang. It includes simple string manipulation,
functions, and basic I/O functions. */

//Slices and returns the string based on the start index and end index.
//Note: Ending index is not inclusive!
func slice(str string, start int, end int) string {
	slicedStr := str[start:end]
	return slicedStr
}

/* Split and return the string into three based on its: warehouse number(first 3 characters)
warehouse zone(following 3 characters) and its Storage unit #(last 8 characters). Use slice func. */
func split(str string) (string, string, string) {
	// warehouseNum := str[0:3]
	// warehouseZone := str[3:6]
	// storageUnit := str[6:15]

	return str[0:3], str[3:6], str[6:14]
}

//Use "strings" package to implement includes
func includes(str string, target string) bool {
	return strings.Contains(target, str)
}

//In the main() function, open the file "storage_unit_numbers.txt"
//and for each line of input, write only the Storage Unit #'s(last 8 characters) in a separate "output.txt" file.
func main() {
	/* IMPORT FILE */
	file, err := os.Open("storage_unit_numbers.txt") //Opens file for reading

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() //Closes file descriptor at end of main function

	/* CREATE OUTPUT FILE */
	oFile, err2 := os.Create("output.txt")

	if err2 != nil {
		log.Fatal(err2)
	}
	defer oFile.Close()

	scanner := bufio.NewScanner(file) //New scanner is created

	//Scan() advances the Scanner to the next token
	for scanner.Scan() {
		var num, zone, unit = split(scanner.Text())
		fmt.Printf("\nWarehouse number: %s", num)
		fmt.Printf("\nWarehouse zone: %s", zone)
		fmt.Printf("\nStorage unit: %s\n", unit)
		oFile.WriteString(unit + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
