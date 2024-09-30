package scripts

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func GitClone(projectName, templateType, templateURL string) error {
	if templateType == "" || templateURL == "" {
		return fmt.Errorf("project template not found")
	}

	currentDir, _ := os.Getwd()
	folder := filepath.Join(currentDir, projectName)

	_, errPlainClone := git.PlainClone(
		folder,
		false,
		&git.CloneOptions{
			URL: getAbsoluteURL(templateURL),
		},
	)
	if errPlainClone != nil {
		return fmt.Errorf("repository `%v` was not cloned", templateURL)
	}

	// Removeing .git directory from cloned repo
	err := os.RemoveAll(filepath.Join(folder, ".git"))
	if err != nil {
		return fmt.Errorf("failed to remove .git directory: %v", err)
	}

	githubFiles := []string{
		"README.md",
		".github",
		".gitignore",
		"LICENSE",
	}

	for _, file := range githubFiles {
		err := os.RemoveAll(filepath.Join(folder, file))
		if err != nil {
			return fmt.Errorf("failed to remove %s: %v", file, err)
		}
	}

	return nil
}

func getAbsoluteURL(templateURL string) string {
	templateURL = strings.TrimSpace(templateURL)
	u, _ := url.Parse(templateURL)

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	return u.String()
}
