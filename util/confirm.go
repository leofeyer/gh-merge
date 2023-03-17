package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Confirm(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)

		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)

		if s == "" {
			return def
		}

		if s == "y" {
			return true
		}

		if s == "n" {
			return false
		}
	}
}
