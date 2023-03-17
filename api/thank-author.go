package api

import (
	"encoding/json"
	"fmt"

	"github.com/cli/go-gh"
	"github.com/leofeyer/gh-merge/util"
)

func ThankAuthor(pr string) error {
	user, err := getUser()
	if err != nil {
		return err
	}

	author, err := getAuthor(pr)
	if err != nil {
		return err
	}

	// We do not want to thank ourselves
	if user == author {
		return nil
	}

	prompt := util.Confirm("Say thank you?", true)
	if prompt == false {
		return nil
	}

	data, _, err := gh.Exec("pr", "comment", pr, "--body", "Thank you @"+author+".")
	if err != nil {
		return err
	}

	fmt.Print(data.String())
	return nil
}

func getUser() (string, error) {
	data, _, err := gh.Exec("config", "get", "user", "-h", "github.com")
	if err != nil {
		return "", err
	}

	return data.String(), nil
}

func getAuthor(pr string) (string, error) {
	data, _, err := gh.Exec("pr", "view", pr, "--json", "author")
	if err != nil {
		return "", err
	}

	type Result struct {
		Author struct {
			Login string `json:"login"`
		}
	}

	var r Result

	err = json.Unmarshal(data.Bytes(), &r)
	if err != nil {
		return "", err
	}

	return r.Author.Login, nil
}
