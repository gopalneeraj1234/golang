package main

import (
	"flag"
	"log"
)

type UserFlags struct {
	charCount bool
}

func (u *UserFlags) IsCharCount() bool {
	return u.charCount
}

func ProcessFlags() *UserFlags {
	charCountFlag := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	flag.Parse()
	log.Printf("The char count flag is set to %t", *charCountFlag)
	return &UserFlags{
		charCount: *charCountFlag,
	}
}
