package main

//go:generate go build -o ./data/generator generator.go
//go:generate ./data/generator ./data/fng.csv

import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/omise/go-tamboon/cipher"
)

type Record struct {
	name   string
	amount int64

	ccNumber          string
	ccCVV             string
	ccExpirationMonth int
	ccExpirationYear  int
}

func (record *Record) ParseCSV(data []string) {
	// Number,CCType,CCNumber,CVV2,CCExpires,Title,GivenName,MiddleInitial,Surname
	// 0     ,1     ,2       ,3   ,4        ,5    ,6        ,7            ,8
	record.name = strings.Join(data[5:], " ")
	record.amount = 100000 + rand.Int63n(5000000)

	record.ccNumber = data[2]
	record.ccCVV = data[3]

	record.ccExpirationMonth = int(time.Now().Month())
	record.ccExpirationYear = time.Now().Year() + 1
	if idx := strings.IndexRune(data[4], '/'); idx > 0 {
		if expMonth, err := strconv.Atoi(data[4][:idx]); err == nil {
			record.ccExpirationMonth = expMonth
		}
		if expYear, err := strconv.Atoi(data[4][idx+1:]); err == nil {
			record.ccExpirationYear = expYear
		}
	}
}

func (record *Record) CSVHeader() []string {
	return []string{
		"Name",
		"AmountSubunits",
		"CCNumber",
		"CVV",
		"ExpMonth",
		"ExpYear",
	}
}

func (record *Record) CSV() []string {
	if record == nil {
		return []string{"", "", "", "", "", ""}
	}

	fmt := func(n int64) string { return strconv.FormatInt(n, 10) }

	return []string{
		record.name,
		fmt(record.amount),
		record.ccNumber,
		record.ccCVV,
		fmt(int64(record.ccExpirationMonth)),
		fmt(int64(record.ccExpirationYear)),
	}
}

func importRecords(r io.Reader) <-chan *Record {
	ch := make(chan *Record)

	go func() {
		defer close(ch)

		reader := csv.NewReader(r)
		reader.FieldsPerRecord = 9
		reader.ReuseRecord = true

		if _, err := reader.Read(); err != nil { // header row
			log.Fatalln(err)
		}

		for {
			if row, err := reader.Read(); err != nil {
				if err == io.EOF {
					return
				} else {
					log.Fatalln(err)
				}

			} else {
				record := &Record{}
				record.ParseCSV(row)
				ch <- record
			}
		}
	}()

	return ch
}

func exportRecords(w io.Writer, ch <-chan *Record) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		var rec *Record

		writer := csv.NewWriter(w)
		if err := writer.Write(rec.CSVHeader()); err != nil {
			log.Fatalln(err)
		}

		for rec = range ch {
			writer.Write(rec.CSV())
		}
	}()

	return done
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s (fng-list)", os.Args[0])
		return
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

	log.Println("processing...")
	defer log.Println("finished.")

	encOut, err := cipher.NewRot255Writer(outFile)
	if err != nil {
		log.Fatalln(err)
	}

	encOut, err = cipher.NewRot255Writer(encOut)
	if err != nil {
		log.Fatalln(err)
	}

	ch := importRecords(fngFile)
	done := exportRecords(encOut, ch)
	<-done
}
