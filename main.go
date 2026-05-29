package main

import (
	"fmt"
	"net/url"
	"strconv"
	"unicode"

	flag "github.com/spf13/pflag"
	"golang.org/x/text/unicode/norm"
)

var base int
var from, to string
var similar string

var verbose, ascii bool

var urlEncoding bool

func main() {
	parseFlag()

	f, err_f := strconv.ParseInt(from, 16, 32)
	t, err_t := strconv.ParseInt(to, 16, 32)

	if err_f != nil {
		fmt.Println("[-] error :", err_f)
		return
	}
	if err_t != nil {
		fmt.Println("[-] error :", err_t)
		return
	}
	if verbose {

		fmt.Println(f, t)
	}
	//handle similar
	if similar != "" {
		for i := 0; i <= unicode.MaxRune; i++ {
			if norm.NFKD.String(string(i)) == similar ||
				norm.NFKC.String(string(i)) == similar {
				if ascii {
					fmt.Println(string(i))
				} else {
					r := strconv.FormatInt(int64(i), base)
					if len(r) == 1 {
						a := "0" + r
						r = a
					}
					fmt.Println(r)
				}
			}
		}
	} else {
		//handle generate hex
		for i := f; i <= t; i++ {
			if ascii {
				fmt.Println(string(i))
			} else if urlEncoding {
				fmt.Println(url.QueryEscape(string(i)))
			} else {

				r := strconv.FormatInt(i, base)
				if len(r) == 1 {
					a := "0" + r
					r = a
				}
				fmt.Println(r)
			}
		}
	}
}

func parseFlag() {
	flag.StringVarP(&from, "from", "f", "0", "min hex number")
	flag.StringVarP(&to, "to", "t", "7f", "max hex number")
	flag.StringVarP(&similar, "similar", "s", "", "get all similar character")

	flag.BoolVarP(&verbose, "verbose", "v", false, "show verbose")
	flag.BoolVarP(&ascii, "ascii", "a", false, "show ascii")
	flag.BoolVarP(&urlEncoding, "url", "u", false, "show url encoded")

	flag.IntVarP(&base, "base", "b", 16, "format number with specific base ")
	flag.Parse()
}
