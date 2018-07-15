package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Variables for command line flags
var (
	inputFileName, outputFileName string
	reportCounts                  bool
	replacements                  []string
)

func init() {
	flag.StringVar(&inputFileName, "in", "", "name of input file to process (default stdin)")
	flag.StringVar(&outputFileName, "out", "", "name of the output file to create (default stdout)")
	flag.BoolVar(&reportCounts, "counts", false, "report counts of replacements")
	flag.Parse()
	replacements = flag.Args()
}

func main() {

	checkArguments()

	lineScanner := bufio.NewScanner(getInput())
	output := getOutput()

	var lineCounter int
	replacementCounter := make(map[string]int)

	for lineScanner.Scan() {

		line := lineScanner.Text()

		for i := 0; i < len(replacements); i += 2 {

			oldString := replacements[i]
			newString := replacements[i+1]

			newLine := strings.Replace(line, oldString, newString, -1)
			if newLine != line {
				replacementCounter[oldString]++
			}
			line = newLine
		}

		_, err := fmt.Fprintln(output, line)
		if err != nil {
			log.Fatal(err)
		}

		lineCounter++
	}

	if err := lineScanner.Err(); err != nil {
		log.Fatal(err)
	}

	if reportCounts {
		_, err := fmt.Fprintf(os.Stderr, "%d lines processed\n", lineCounter)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(replacements); i += 2 {
			oldString := replacements[i]
			newString := replacements[i+1]

			_, err := fmt.Fprintf(os.Stderr, "%d lines replaced \"%s\" with \"%s\"\n", replacementCounter[oldString], oldString, newString)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// checkArguments will make sure the user has passed sane arguments. Prints
// helpful stuff then exits program if any problem found.
func checkArguments() {

	if len(replacements) < 2 || len(replacements)%2 != 0 {
		_, err := fmt.Fprintln(os.Stderr, "usage: replacestrings [-in inputfile] [-out outputfile] [-counts] oldstring1 newstring1 oldstring2 newstring2 ...")
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	for i := 0; i < len(replacements); i += 2 {
		if replacements[i] == "" {
			log.Fatal("old strings to replace must be non-empty")
		}
	}
}

// getInput will return either a new input stream from the specified (on command
// line) file name, or stdin. Exits programs if any problem.
func getInput() *os.File {
	if inputFileName == "" {
		return os.Stdin
	}

	input, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

// getOuput will return either a new output stream from the specified (on
// command line) file name, or stdout. Exits program if any problem.
func getOutput() *os.File {
	if outputFileName == "" {
		return os.Stdout
	}

	output, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}

	return output
}
