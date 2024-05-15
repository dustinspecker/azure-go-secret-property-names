package secretnames

import (
	"slices"
	"strings"
)

// GetSubset returns all Azure Go SDK secret property names minus provided ignore list
func GetSubset(ignoreList []string) []string {
	allSecrets := GetAll()

	return slices.DeleteFunc(allSecrets, func(secret string) bool {
		return slices.ContainsFunc(ignoreList, func(ignore string) bool {
			return strings.EqualFold(secret, ignore)
		})
	})
}
