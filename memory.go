package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var match_mem = regexp.MustCompile(`Mem:\s+(.+)`)

func memory_usage() (string, error) {
	data, err := exec.Command("free", "-m").Output()
	if err != nil {
		return "", err
	}

	m := match_mem.FindStringSubmatch(string(data))
	if len(m) != 2 {
		return "", fmt.Errorf("number of matches was not expected for mem submatch")
	}

	parts := strings.Fields(m[1])
	total, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}
	avail, err := strconv.Atoi(parts[5])
	if err != nil {
		return "", err
	}
	used := total - avail
	perc := 100 * used / total

	var color string
	switch {
	case perc >= 90:
		color = "#dc322f"
	case perc >= 70:
		color = "#b58900"
	default:
		color = "#859900"
	}

	return fmt.Sprintf("^fg(%s)%d%% ^i(%s)^fg()", color, perc, xbm("mem")), nil
}