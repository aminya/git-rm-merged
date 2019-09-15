package main

import (
	"os/exec"
	"strings"
)

func getMergedBranches() ([]string, error) {
	var branches []string
	out, err := exec.Command("git", "branch", "--merged").Output()
	if err != nil {
		return nil, err
	}

	for _, branch := range strings.Split(string(out), "\n") {
		branch = strings.TrimSpace(branch)
		if (branch == "") {
			continue
		}
		// * develop のようなカレントブランチは対象外
		if strings.HasPrefix(branch, "*") {
			continue
		}
		// master/developは対象外
		if strings.Contains(branch, "master") || strings.Contains(branch, "develop") {
			continue
		}

		branches = append(branches, branch)
	}
	return branches, nil
}

func delBranch(name string) error {
	return exec.Command("git", "branch", "-d", name).Run()
}
