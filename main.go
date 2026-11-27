// b36date prints a date in the compact YYMD form: a two-digit year followed by
// the month and the day, each as a single base36 digit (1-9, A-Z).
//
//	b36date             -> today, e.g. 26BQ for 2026-11-26
//	b36date 2026-11-26  -> 26BQ
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var layouts = []string{
	"2006-01-02",
	"2006/01/02",
	"20060102",
	"01/02/2006",
	"Jan 2 2006",
	"January 2 2006",
	"2 Jan 2006",
	time.RFC3339,
}

func encode(t time.Time) string {
	return fmt.Sprintf("%02d%c%c", t.Year()%100, digits[t.Month()], digits[t.Day()])
}

func parse(s string) (time.Time, error) {
	s = strings.TrimSpace(strings.ReplaceAll(s, ",", ""))
	for _, l := range layouts {
		if t, err := time.Parse(l, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized date %q (try YYYY-MM-DD)", s)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(encode(time.Now()))
		return
	}
	for _, arg := range os.Args[1:] {
		t, err := parse(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "b36date:", err)
			os.Exit(1)
		}
		fmt.Println(encode(t))
	}
}
