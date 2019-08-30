// +build !linux

package lumberjack

import (
	"os"
	"strconv"
	"strings"
)

func chown(name string, own string) error {
	var userID int
	var groupID int

	// windows
	if os.Getuid() == -1 {
		return nil
	}

	if own == "" {
		return os.Chown(name, os.Getuid(), os.Getgid())
	} else if index := strings.Index(own, ":"); index != -1 {
		uid, err := strconv.ParseInt(own[:index], 10, 32)
		if err != nil {
			return err
		}
		gid, _ := strconv.ParseInt(own[index+1:], 10, 32)
		if err != nil {
			return err
		}

		userID = int(uid)
		groupID = int(gid)
	} else {
		uid, err := strconv.ParseInt(own, 10, 32)
		if err != nil {
			return err
		}

		userID = int(uid)
		groupID = os.Getgid()
	}

	return os.Chown(name, userID, groupID)
}
