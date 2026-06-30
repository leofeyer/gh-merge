package api

import (
	"encoding/json"
	"strconv"
)

func GetNumber(args []string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	data, err := execGh("pr", "view", "--json", "number")
	if err != nil {
		return "", err
	}

	var r struct {
		Number int `json:"number"`
	}

	if err := json.Unmarshal(data.Bytes(), &r); err != nil {
		return "", err
	}

	return strconv.Itoa(r.Number), nil
}
