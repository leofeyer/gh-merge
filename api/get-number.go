package api

import (
	"encoding/json"
	"strconv"

	"github.com/cli/go-gh"
)

func GetNumber(args []string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	data, _, err := gh.Exec("pr", "view", "--json", "number")
	if err != nil {
		return "", err
	}

	type Result struct {
		Number int `json:"number"`
	}

	var r Result

	err = json.Unmarshal(data.Bytes(), &r)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(r.Number), nil
}
