package jsonequaliser

import "testing"

func TestEmptyJSONArray(t *testing.T) {
	_, err := getJSONNodeFromString("[]")

	if _, ok := err.(*emptyJSONArrayError); !ok {
		t.Fatal("Expected an emptyJSONArrayError")
	}
}
