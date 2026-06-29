package api

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cli/go-gh/v2"
)

func execGh(args ...string) (bytes.Buffer, error) {
	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		message := strings.TrimSpace(stderr.String())
		if message != "" {
			return stdout, fmt.Errorf("gh execution failed: %w\n%s", err, message)
		}

		return stdout, err
	}

	return stdout, nil
}
