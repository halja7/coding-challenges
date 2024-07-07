package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type Args struct {
	filename string
	bytes    bool
	words    bool
	lines    bool
}

type Fileinfo struct {
	filename string
	lines    int
	words    int
	bytes    int
}

func main() {
	args := parseArgs()

	fileinfo := Fileinfo{
		filename: args.filename,
	}

	var file []byte
	var err error
	if args.filename == "" {
		file, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		f, err := os.Open(args.filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer f.Close()
		file, err = io.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if args.bytes {
		bytes, err := countBytes(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fileinfo.bytes = bytes
	}

	if args.words {
		words, err := countWords(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fileinfo.words = words
	}

	if args.lines {
		lines, err := countLines(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fileinfo.lines = lines
	}

	formatOutput(fileinfo)
}

func formatOutput(fileinfo Fileinfo) {
	if fileinfo.lines > 0 {
		fmt.Printf("%d ", fileinfo.lines)
	}

	if fileinfo.words > 0 {
		fmt.Printf("%d ", fileinfo.words)
	}

	if fileinfo.bytes > 0 {
		fmt.Printf("%d ", fileinfo.bytes)
	}

	if fileinfo.bytes <= 0 && fileinfo.words <= 0 && fileinfo.lines <= 0 {
		fmt.Printf(
			"%d %d %d ",
			fileinfo.lines,
			fileinfo.words,
			fileinfo.bytes,
		)
	}

	fmt.Println(fileinfo.filename)
}

// It would be way better to do a single loop and
// pull out all of this data each time regardless
// of the flags supplied, printing only what the
// user asked for. This was my starting point and
// maybe I'll come back to it later.
func countBytes(file []byte) (int, error) {
	return len(file), nil
}

func countWords(file []byte) (int, error) {
	count := 0
	inWord := false

	for len(file) > 0 {
		r, size := utf8.DecodeRune(file)
		if unicode.IsSpace(r) {
			inWord = false
		} else if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' || isCJK(r) {
			if !inWord || isCJK(r) {
				count++
				inWord = true
			}
		} else {
			inWord = false
		}
		file = file[size:]
	}

	return count, nil
}

func isCJK(r rune) bool {
	return unicode.In(r, unicode.Han, unicode.Hiragana, unicode.Katakana, unicode.Hangul)
}

func countLines(file []byte) (int, error) {
	count := 0

	for _, b := range file {
		if b == '\n' {
			count++
		}
	}

	if len(file) > 0 && file[len(file)-1] != '\n' {
		count++
	}

	return count, nil
}

func parseArgs() *Args {
	args := &Args{}

	flag.BoolVar(&args.bytes, "bytes", false, "Count bytes")
	flag.BoolVar(&args.bytes, "c", false, "Count bytes (shorthand)")

	flag.BoolVar(&args.words, "words", false, "Count words")
	flag.BoolVar(&args.words, "w", false, "Count words (shorthand)")

	flag.BoolVar(&args.lines, "lines", false, "Count lines")
	flag.BoolVar(&args.lines, "l", false, "Count lines (shorthand)")

	flag.Parse()

	args.filename = flag.Arg(0)

	return args
}
