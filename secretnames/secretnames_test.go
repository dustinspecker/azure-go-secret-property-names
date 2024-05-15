package secretnames_test

import (
	"slices"
	"testing"

	"github.com/dustinspecker/azure-go-secret-property-names/secretnames"
)

func TestGetAll(t *testing.T) {
	t.Parallel()

	allSecretNames := secretnames.GetAll()

	if len(allSecretNames) == 0 {
		t.Error("Expected GetAll to return at least one element")
	}
}

func TestGetSubset(t *testing.T) {
	t.Parallel()

	allWords := secretnames.GetAll()
	if !slices.Contains(allWords, "id") {
		t.Fatalf("Expected GetAll to contain 'id', unable to validate subset works")
	}

	ignoreList := []string{
		"id",
	}
	subset := secretnames.GetSubset(ignoreList)

	if len(subset) == 0 {
		t.Error("Expected GetSubset to return at least one element")
	}

	if slices.Contains(subset, "id") {
		t.Error("Expected GetSubset to not contain 'id'")
	}
}

func TestGetSubsetIsCaseInsensitive(t *testing.T) {
	t.Parallel()

	ignoreList := []string{
		"ID",
	}
	subset := secretnames.GetSubset(ignoreList)

	if slices.Contains(subset, "id") {
		t.Error("Expected GetSubset to filter out 'id' case insensitively")
	}
}
