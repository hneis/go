// Package main provides ...
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hneis/go/lesson5/copier"
	"github.com/hneis/go/lesson5/grep"
)

func task3() {
	var task int
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.IntVar(&task, "t", 3, "Current task")
	flag.Parse()

	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}

func task4() {
	var task int
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.IntVar(&task, "t", 4, "Current task")
	flag.Parse()

	argsLen := len(flag.Args())

	if argsLen < 2 {
		fmt.Println("Usage: copy source destination")
	}
	source := "./source.txt"
	destination := "./destination.txt"
	var err error

	sFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer sFile.Close()

	dFile, err := os.OpenFile(destination, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dFile.Close()

	copier.NewCopier(sFile, dFile).Copy()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func classWork() {

	// var fileSource, hashFile, outFile, signedFile string

	// flag.StringVar(&fileSource, "sourceFile", "", "File source")
	// flag.StringVar(&hashFile, "hashFile", "", "Hash file")
	// flag.StringVar(&outFile, "outFile", "sign.txt", "Save to file")
	// flag.StringVar(&signedFile, "signedFile", "", "File signed")

	// flag.Parse()

	// fmt.Println(fileSource, hashFile, outFile)

	// action := flag.Arg(0)
	// switch action {
	// case "enc":
	// 	encoder, err := signature.NewEncoder(fileSource, hashFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	err = encoder.EncryptSHA256()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	err = encoder.SaveToFile(outFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// case "dec":
	// 	checker, err := signature.NewEncryptCheck(fileSource, hashFile, signedFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	err = checker.Check()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("Signature success")
	// 	}
	// default:
	// 	log.Fatal("Use enc or dec param first")
	// }
}

func task5() {
	var grepLineNumber, grepInvert bool
	var task int
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.IntVar(&task, "t", 5, "Current task")
	flag.BoolVar(&grepLineNumber, "n", false, "show line number")
	flag.BoolVar(&grepInvert, "v", false, "select no matching lines")
	flag.Parse()

	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	pattern := flag.Args()[0]
	files := (flag.Args())[1:]
	info, _ := os.Stdin.Stat()
	grep := grep.NewGrep("\x1b[39m", "\x1b[31m")
	grep.LineNumber = grepLineNumber
	grep.Invert = grepInvert
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		for _, fileName := range files {
			if fileExists(fileName) {
				rFile, err := os.Open(fileName)
				if err != nil {
					panic(err)
				}
				defer rFile.Close()

				grep.Match(bufio.NewReader(rFile), pattern)
				if grep.IsMatching() && !grep.Invert {
					fmt.Printf("%s: \n", fileName)
					grep.Print()
					fmt.Printf("\n")
				}

			}
		}
	} else {
		grep.Match(bufio.NewReader(os.Stdin), pattern)
		if grep.IsMatching() && !grep.Invert {
			grep.Print()
		}
	}
}

func main() {
	// initFlags()
	// task3()
	var task int

	flag.IntVar(&task, "t", 0, "Task")
	flag.Parse()

	switch task {
	case 1:
		fmt.Println("Task 1")
	case 2:
		fmt.Println("Task 2")
	case 3:
		task3()
	case 4:
		task4()
	case 5:
		task5()
	}

}
