package lumberjack

import (
	"os"
	"strconv"
	"strings"
)

// os_Chown is a var so we can mock it out during tests.
var os_Chown = os.Chown

func chown(name string, own string) error {
	var userID int
	var groupID int

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

	return os_Chown(name, userID, groupID)
}
