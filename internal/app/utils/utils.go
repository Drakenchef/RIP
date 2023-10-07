package utils

import (
	"fmt"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/rs/xid"
	"strings"
)

func FindElement(slice []ds.Planet, target ds.Planet) int {
	for i, val := range slice {
		if val == target {
			return i
		}
	}

	return -1
}

func Max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func GenerateUniqueName(imageName *string) error {
	parts := strings.Split(*imageName, ".")
	if len(parts) > 1 {
		fileExt := parts[len(parts)-1]
		uniqueID := xid.New()
		*imageName = fmt.Sprintf("%s.%s", uniqueID.String(), fileExt)
		return nil
	}
	return fmt.Errorf("uncorrect file name. not fount image extension")
}
