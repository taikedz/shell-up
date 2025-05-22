package shup

import (
	"os"
	"fmt"
	"errors"
	"strings"
	"os/user"
	"path/filepath"
)

func homePath(target string) (string, error) {
	curuser, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(curuser.HomeDir, target[2:]), nil
}

func AbsPath(target string) (string, error) {
	if strings.Index(target, "~/") == 0 {
		path, err := homePath(target)
		if err != nil {
			return "", err
		}
		target = path
	} else {
		path, err := filepath.Abs(target)
		if err != nil {
			return "", err
		}
		target = path
	}

	if ! pathAccessible(target) {
		return "", fmt.Errorf("Could not access %s", target)
	}

	end_target, err := filepath.EvalSymlinks(target)
	if err != nil {
		return "", err
	}
	return end_target, nil
}

func pathAccessible(target string) bool {
	_, err := os.Stat(target)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	// If we get any other error we cannot know whether the item exists
	// https://stackoverflow.com/a/12518877/2703818

	// But if we cannot successfully stat it, it may as well not exist
	return false
}