package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/cli/go-gh"
	"github.com/leofeyer/gh-merge/util"
)

func MergePr(pr string) error {
	err := checkStatus(pr)
	if err != nil {
		return err
	}

	subject, err := getSubject(pr)
	if err != nil {
		return err
	}

	body, err := getBody(pr)
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println(body)

	prompt := util.Confirm("Merge '"+subject+"' now?", false)
	if !prompt {
		return errors.New("Cancelled.")
	}

	args := []string{"pr", "merge", pr, "--subject", subject, "--body", body, "--squash"}

	data, _, err := gh.Exec(args...)
	if err != nil {
		if !strings.Contains(err.Error(), "not mergeable: the base branch policy prohibits the merge") {
			return err
		}

		fmt.Println("The pull request is not mergeable because the base branch policy prohibits the merge.")

		prompt = util.Confirm("Merge with admin privileges instead?", false)
		if prompt {
			args = append(args, "--admin")
		} else {
			args = append(args, "--auto")
		}

		data, _, err = gh.Exec(args...)
		if err != nil {
			return err
		}
	}

	fmt.Print(data.String())
	return nil
}

func checkStatus(pr string) error {
	data, _, err := gh.Exec("pr", "view", pr, "--json", "closed")
	if err != nil {
		return err
	}

	type Result struct {
		Closed bool `json:"closed"`
	}

	var r Result

	err = json.Unmarshal(data.Bytes(), &r)
	if err != nil {
		return err
	}

	if r.Closed {
		return errors.New("The PR is closed.")
	}

	return nil
}

func getSubject(pr string) (string, error) {
	data, _, err := gh.Exec("pr", "view", pr, "--json", "title")
	if err != nil {
		return "", err
	}

	type Result struct {
		Title string `json:"title"`
	}

	var r Result

	err = json.Unmarshal(data.Bytes(), &r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s (see #%s)", r.Title, pr), nil
}

func getBody(pr string) (string, error) {
	data, _, err := gh.Exec("pr", "view", pr, "--json", "author,body,commits")
	if err != nil {
		return "", err
	}

	type Result struct {
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

	var r Result

	err = json.Unmarshal(data.Bytes(), &r)
	if err != nil {
		return "", err
	}

	x := regexp.MustCompile("(?s)<!--.*?-->")

	ret := "Description\n-----------\n\n"
	ret += strings.TrimSpace(x.ReplaceAllString(r.Body, ""))
	ret += "\n\nCommits\n-------\n\n"

	authors := make(map[string]string)

	for i := 0; i < len(r.Commits); i++ {
		if r.Commits[i].Headline == "CS" {
			continue
		}

		if strings.HasPrefix(r.Commits[i].Headline, "Merge branch ") {
			continue
		}

		ret += fmt.Sprintf("%.8s", r.Commits[i].Oid) + " " + r.Commits[i].Headline + "\n"

		if r.Commits[i].Authors[0].Login == r.Author.Login {
			continue
		}

		authors[r.Commits[i].Authors[0].Login] = r.Commits[i].Authors[0].Email
	}

	if len(authors) > 0 {
		ret += "\n"

		for author, email := range authors {
			ret += "Co-authored-by: " + author + " <" + email + ">\n"
		}
	}

	return ret, nil
}
