// +build !linux

package lumberjack

func chown(_ string, _ string) error {
	return nil
}
