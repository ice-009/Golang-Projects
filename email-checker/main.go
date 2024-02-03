package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord\n ")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal("err", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v_spf1") {
			hasSPF = true
			spfRecord = record
			break
		}

	}
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)
	if len(dmarcRecords) > 0 {
		dmarcRecord = dmarcRecords[0]
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(string(record), "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = string(record)
			break
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, hasMx, hasSPF, hasDMARC, spfRecord, dmarcRecord)
}
