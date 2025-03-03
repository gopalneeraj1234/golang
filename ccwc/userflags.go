package main

import (
	"flag"
)

type UserFlags struct {
	charCountFlag      bool
	lineCountFlag      bool
	wordCountFlag      bool
	longestLineLenFlag bool
}

func (u *UserFlags) IsWordCountFlag() bool {
	return u.wordCountFlag
}

func (u *UserFlags) IsCharCountFlag() bool {
	return u.charCountFlag
}

func (u *UserFlags) IsLineCountFlag() bool {
	return u.lineCountFlag
}

func (u *UserFlags) IsLongestLineLenFlag() bool {
	return u.longestLineLenFlag
}

func ProcessFlags() *UserFlags {
	charCountFlag := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	lineCountFlag := flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
	wordCountFlag := flag.Bool("w", false, "The number of words in each input file is written to the standard output.")
	longestLineLenFlag := flag.Bool("L", false, "Write the length of the line containing the most bytes (default) or characters (when -m is provided) to standard output.  When more than one file argument is specified, the longest input line of all files is reported as the value of the final \"total\".")
	flag.Parse()
	//log.Printf("The char count flag is set to %t", *charCountFlag)
	//log.Printf("The line count flag is set to %t", *lineCountFlag)
	return &UserFlags{
		charCountFlag:      *charCountFlag,
		lineCountFlag:      *lineCountFlag,
		wordCountFlag:      *wordCountFlag,
		longestLineLenFlag: *longestLineLenFlag,
	}
}
