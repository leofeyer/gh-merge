package api

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/leofeyer/gh-merge/util"
)

var htmlCommentRegex = regexp.MustCompile(`(?s)<!--.*?-->`)

func MergePr(pr string, info *PrInfo) error {
	if info.Closed {
		return errors.New("the pull request is closed")
	}

	subject := fmt.Sprintf("%s (see #%s)", info.Title, pr)
	body := buildBody(info)

	fmt.Println("")
	fmt.Println(body)

	prompt := util.Confirm("Merge '"+subject+"' now?", false)
	if !prompt {
		return errors.New("cancelled")
	}

	args := []string{"pr", "merge", pr, "--subject", subject, "--body", body, "--squash"}

	data, err := execGh(args...)
	if err != nil {
		if !strings.Contains(err.Error(), "not mergeable: the base branch policy prohibits the merge") {
			return err
		}

		fmt.Println("The pull request is not mergeable because the base branch policy prohibits the merge.")

		prompt = util.Confirm("Merge with admin privileges instead?", false)
		if !prompt {
			return errors.New("cancelled")
		}

		data, err = execGh(append(args, "--admin")...)
		if err != nil {
			return err
		}
	}

	fmt.Print(data.String())
	return nil
}

func buildBody(info *PrInfo) string {
	var ret strings.Builder
	ret.WriteString("Description\n-----------\n\n")
	ret.WriteString(strings.TrimSpace(htmlCommentRegex.ReplaceAllString(info.Body, "")))
	ret.WriteString("\n\nCommits\n-------\n\n")

	authors := make(map[string]string)

	for i := 0; i < len(info.Commits); i++ {
		if info.Commits[i].Headline == "CS" {
			continue
		}

		if info.Commits[i].Headline == "Rebuild the assets" {
			continue
		}

		if strings.HasPrefix(info.Commits[i].Headline, "Merge branch ") {
			continue
		}

		ret.WriteString(fmt.Sprintf("%.8s", info.Commits[i].Oid) + " " + info.Commits[i].Headline + "\n")

		if len(info.Commits[i].Authors) < 1 {
			continue
		}

		if info.Commits[i].Authors[0].Login == info.Author {
			continue
		}

		authors[info.Commits[i].Authors[0].Login] = info.Commits[i].Authors[0].Email
	}

	if len(authors) > 0 {
		ret.WriteString("\n")

		for author, email := range authors {
			ret.WriteString("Co-authored-by: " + author + " <" + email + ">\n")
		}
	}

	return ret.String()
}
