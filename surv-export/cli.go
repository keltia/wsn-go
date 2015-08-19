// cli.go
//
// Everything related to command-line flag handling
//
// Copyright 2015 © by Ollivier Robert for the EEC
//

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// cli
	fVerbose	bool
	fOutput		string
	fTimeout   int64
	fsTimeout  string
)

// my usage string
const (
	cliUsage	= `
Usage: %s [-o FILE] [-i N(s|mn|h|d)] [-v] feeds...
`
)

// Redefine Usage
var Usage = func() {
        fmt.Fprintf(os.Stderr, cliUsage, os.Args[0])
        flag.PrintDefaults()
}

// called by flag.Parse()
func init() {
	// cli
	flag.BoolVar(&fVerbose, "v", false, "Set verbose flag.")
	flag.StringVar(&fsTimeout, "i", "60s", "Stop after N s/mn/h/days")
	flag.StringVar(&fOutput, "o", "", "Specify output FILE.")
}
