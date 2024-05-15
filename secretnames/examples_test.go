package secretnames_test

import (
	"fmt"
	"slices"

	"github.com/dustinspecker/azure-go-secret-property-names/secretnames"
)

func ExampleGetAll() {
	// GetAll returns all Azure Go SDK secret property names
	allSecretNames := secretnames.GetAll()

	fmt.Println(len(allSecretNames) > 0)

	// Output:
	// true
}

func ExampleGetSubset() {
	// You may decide that GetAll returns too many secret property names for your use case.
	// GetSubset allows you to filter out secret property names you don't care about.
	// ignoreList is matched case insensitively
	ignoreList := []string{
		"id",
	}
	subset := secretnames.GetSubset(ignoreList)

	fmt.Println(len(subset) > 0)
	fmt.Println(!slices.Contains(subset, "id"))

	// Output:
	// true
	// true
}
