package main

import (
	"log"
	"os"

	"github.com/omise/go-tamboon/cipher"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s (fng-list)\n", os.Args[0])
	}

	var (
		fngName = os.Args[1]
		outName = fngName + ".rot255"
	)

	fngFile, err := os.Open(fngName)
	if err != nil {
		log.Fatalln(err)
	}

	defer fngFile.Close()

	outFile, err := os.Create(outName)
	if err != nil {
		log.Fatalln(err)
	}

	defer outFile.Close()

	encOut, err := cipher.NewRot255Writer(outFile)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("processing...")
	defer log.Println("finished.")
	if err := process(fngFile, encOut); err != nil {
		log.Fatalln(err)
	}
}