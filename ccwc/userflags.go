package main

import (
	"flag"
)

type UserFlags struct {
	charCountFlag bool
	lineCountFlag bool
	wordCountFlag bool
}

func (u *UserFlags) IsWordCount() bool {
	return u.wordCountFlag
}

func (u *UserFlags) IsCharCount() bool {
	return u.charCountFlag
}

func (u *UserFlags) IsLineCount() bool {
	return u.lineCountFlag
}

func ProcessFlags() *UserFlags {
	charCountFlag := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	lineCountFlag := flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
	wordCountFlag := flag.Bool("w", false, "The number of words in each input file is written to the standard output.")
	flag.Parse()
	//log.Printf("The char count flag is set to %t", *charCountFlag)
	//log.Printf("The line count flag is set to %t", *lineCountFlag)
	return &UserFlags{
		charCountFlag: *charCountFlag,
		lineCountFlag: *lineCountFlag,
		wordCountFlag: *wordCountFlag,
	}
}
