package api

import (
	"fmt"
	"strings"

	"github.com/leofeyer/gh-merge/util"
)

func ThankAuthor(pr string, info *PrInfo) error {
	user, err := getUser()
	if err != nil {
		return err
	}

	if user == info.Author {
		return nil
	}

	prompt := util.Confirm("Say thank you?", true)
	if !prompt {
		return nil
	}

	data, err := execGh("pr", "comment", pr, "--body", "Thank you @"+info.Author+".")
	if err != nil {
		return err
	}

	fmt.Print(data.String())
	return nil
}

func getUser() (string, error) {
	data, err := execGh("config", "get", "user", "-h", "github.com")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(data.String()), nil
}
