package fromFile

// fromFile -- run the cache from a fromFile of commands

import (
	"encoding/csv"
	"errors"
	"github.com/davecb/stupidestCache/src/stupidestCache"
	"io"
	"log"
	"os"
	"strings"
)

var eShort = errors.New("csv record was too short")
var eBlank = errors.New("csv record was blank")

func Run(filename string) {

	if filename == "" {
		log.Fatalf("No load-test .csv fromFile provided, halting.\n")
	}
	exercise(filename)
}

// code for running from a fromFile

func exercise(filename string) {
	var record []string
	var operation, key, value string
	var err error
	var cache = stupidestCache.New()
	defer cache.Close()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening %s: %s, halting.", filename, err)
	}
	defer f.Close() // nolint

	r := csv.NewReader(f)
	r.Comma = ' '
	r.Comment = '#'
	r.FieldsPerRecord = -1 // ignore differences

loop:
	for {
		record, err = r.Read()
		switch {
		case err == io.EOF:
			break loop
		case err != nil:
			log.Fatalf("failure in csv.Read(), halting. err = %v, record = %q\n", err, record)
		}
		log.Printf("stupid read %q\n", record)

		operation, key, value, err = parseCsv(record)
		if err == eBlank {
			// ignore blank csv lines, we don't mind
			continue
		}
		if err != nil {
			log.Fatalf("parse error, halting. err = %v, record = %q\n", err, record)
		}
		log.Printf("operation = %q, key = %q, value = %q, err = %v\n", operation, key, value, err)

		switch operation {
		case "g":
			x, present := cache.Get(key)
			InterpretGet(present, x, key, value)

		case "p":
			err = cache.Put(key, value)
			InterpretPut(err, key, value)

		default:
			log.Fatalf(`ill-formed csv line. Need either "p" or "g". record = %q`+"\n", record)
		}
	}
	// all done
}

// InterpretPut looks to see if we did a put operation correctly
func InterpretPut(err error, key string, value string) {
	if err != nil {
		log.Printf("couldn't put key = %q, value = %q, err = %v\n", key, value, err)
	} else {
		log.Printf("put suceeded\n")
	}
}

// InterpretGet looks to see if we could get an expected value from the cache
func InterpretGet(present bool, x, key, value string) {
	if !present {
		log.Printf("key/value absent, returned !present\n")
	}
	if x != value {
		log.Printf("comparison failed. key = %q, stupidestCache = %q, input = %q\n", key, x, value)
	} else {
		log.Printf("get and comparison suceeded\n")
	}
}

// parseCsv will parse three or more field csv files for this experiment
func parseCsv(record []string) (string, string, string, error) {
	switch len(record) {
	case 0:
		return "", "", "", eBlank // technically an error, albeit harmless
	case 1:
		return record[0], "", "", eShort
	case 2:
		return record[0], record[1], "", eShort
	default:
		return record[0], record[1], strings.Join(record[2:], " "), nil
	}
}
