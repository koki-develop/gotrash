package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func YesNo(msg string) bool {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s (y/N): ", msg)
	_ = sc.Scan()
	yn := sc.Text()

	switch strings.ToLower(yn) {
	case "y", "yes":
		return true
	default:
		return false
	}
}
