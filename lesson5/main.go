// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hneis/go/lesson5/signature"
)

func main() {
	var fileSource, hashFile, outFile, signedFile string

	flag.StringVar(&fileSource, "sourceFile", "", "File source")
	flag.StringVar(&hashFile, "hashFile", "", "Hash file")
	flag.StringVar(&outFile, "outFile", "sign.txt", "Save to file")
	flag.StringVar(&signedFile, "signedFile", "", "File signed")

	flag.Parse()

	fmt.Println(fileSource, hashFile, outFile)

	action := flag.Arg(0)
	switch action {
	case "enc":
		encoder, err := signature.NewEncoder(fileSource, hashFile)
		if err != nil {
			log.Fatal(err)
		}
		err = encoder.EncryptSHA256()
		if err != nil {
			log.Fatal(err)
		}

		err = encoder.SaveToFile(outFile)
		if err != nil {
			log.Fatal(err)
		}
	case "dec":
		checker, err := signature.NewEncryptCheck(fileSource, hashFile, signedFile)
		if err != nil {
			log.Fatal(err)
		}
		err = checker.Check()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Signature success")
		}
	default:
		log.Fatal("Use enc or dec param first")
	}
}
