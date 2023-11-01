package pkg

import (
	"os"
	"strings"
)

func IsKube() bool {
	prefix, exists := os.LookupEnv("PREFIX")
	if !exists {
		return false
	}

	return strings.Contains(prefix, "K8S")
}
