package jsonequaliser

import (
	"testing"
)

const simpleJSON = `{"firstname": "chris", "lastname": "james", "age": 30}`
const comparableJSON = `{"firstname": "christopher", "lastname": "james", "age": 15}`
const notSimilarJSON = `{"foo":"bar"}`

func TestItKnowsTheSameJSONIsEqual(t *testing.T) {
	if isEqual, _ := IsEqual(simpleJSON, simpleJSON); !isEqual {
		t.Error("Should be equal")
	}
}

func TestItKnowsStructurallySameJSONIsEqual(t *testing.T) {
	if isEqual, _ := IsEqual(simpleJSON, comparableJSON); !isEqual {
		t.Error("Should be equal")
	}
}

func TestItKnowsDifferentJSONIsDifferent(t *testing.T) {
	if equal, _ := IsEqual(simpleJSON, notSimilarJSON); equal {
		t.Error("Should not be equal")
	}
}

func TestItKnowsHowToHandleArrays(t *testing.T) {
	JSONWithArray := `{"foo": ["baz"]}`
	comparableJSONWithArray := `{"foo": ["bar"]}`

	if equal, _ := IsEqual(JSONWithArray, comparableJSONWithArray); !equal {
		t.Error("Should be equal")
	}
}

func TestItReturnsAnErrorForNonJson(t *testing.T) {
	if _, err := IsEqual("nonsense", "not json"); err == nil {
		t.Error("Expected an error to be returned for bad json")
	}
}
