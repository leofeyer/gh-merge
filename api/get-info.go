package api

import (
	"encoding/json"
)

type PrInfo struct {
	Closed  bool
	Title   string
	Author  string
	Body    string
	Commits []struct {
		Oid      string `json:"oid"`
		Headline string `json:"messageHeadline"`
		Authors  []struct {
			Login string `json:"login"`
			Email string `json:"email"`
		}
	}
}

func GetInfo(pr string) (*PrInfo, error) {
	data, err := execGh("pr", "view", pr, "--json", "closed,title,author,body,commits")
	if err != nil {
		return nil, err
	}

	var raw struct {
		Closed bool   `json:"closed"`
		Title  string `json:"title"`
		Author struct {
			Login string `json:"login"`
		}
		Body    string `json:"body"`
		Commits []struct {
			Oid      string `json:"oid"`
			Headline string `json:"messageHeadline"`
			Authors  []struct {
				Login string `json:"login"`
				Email string `json:"email"`
			}
		}
	}

	err = json.Unmarshal(data.Bytes(), &raw)
	if err != nil {
		return nil, err
	}

	return &PrInfo{
		Closed:  raw.Closed,
		Title:   raw.Title,
		Author:  raw.Author.Login,
		Body:    raw.Body,
		Commits: raw.Commits,
	}, nil
}
