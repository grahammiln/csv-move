// Copyright © 2019 Graham Miln <https://miln.eu>. All rights reserved.
// Covered by BSD 3-Clause "New" or "Revised" License; see LICENSE file
package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
)

// Written for https://apple.stackexchange.com/questions/363131/
//
// Sample input:
// /Volumes/Backup Plus/Salesforce/Order Test/OR-1000050081,/Volumes/Backup Plus/Salesforce/Opportunity Test/Q-RGORD20170207-1526

const mvPath = "/bin/mv"

var reallyMove bool

func main() {
	flag.BoolVar(&reallyMove, "move", false, "perform the move")
	flag.Parse()

	// Read from stdin - expects LF line endings
	r := csv.NewReader(os.Stdin)
	r.FieldsPerRecord = 2

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%q %s %q %q\n", mvPath, "-v", line[0], line[1])
		if !reallyMove {
			continue
		}

		cmd := exec.Command(mvPath, "-v", line[0], line[1])
		if err := cmd.Run(); err != nil {
			log.Fatalf("move failed: %s", err.Error())
		}
	}
}
