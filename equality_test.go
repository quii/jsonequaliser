package jsonequaliser

import (
	"testing"
)

const simpleJSON = `{"firstname": "chris", "lastname": "james", "age": 30}`
const comparableJSON = `{"firstname": "christopher", "lastname": "james", "age": 15}`
const notSimilarJSON = `{"foo":"bar"}`

func TestItKnowsTheSameJSONIsEqual(t *testing.T) {
	if isEqual, _ := IsCompatible(simpleJSON, simpleJSON); !isEqual {
		t.Error("Should be equal")
	}
}

func TestItKnowsStructurallySameJSONIsEqual(t *testing.T) {
	if isEqual, _ := IsCompatible(simpleJSON, comparableJSON); !isEqual {
		t.Error("Should be equal")
	}
}

func TestItKnowsDifferentJSONIsDifferent(t *testing.T) {
	if equal, _ := IsCompatible(simpleJSON, notSimilarJSON); equal {
		t.Error("Should not be equal")
	}
}

func TestItKnowsHowToHandleArrays(t *testing.T) {
	JSONWithArray := `{"foo": ["baz"]}`
	comparableJSONWithArray := `{"foo": ["bar"]}`

	if equal, _ := IsCompatible(JSONWithArray, comparableJSONWithArray); !equal {
		t.Error("Should be equal")
	}
}

func TestItDoesntMindSuperflousFieldsInB(t *testing.T) {
	extraJSON := `{"firstname":"frank", "lastname": "sinatra", "eyecolour": "blue", "age":70}`

	if equal, _ := IsCompatible(simpleJSON, extraJSON); !equal {
		shouldBeEqualErrorMsg(t, simpleJSON, extraJSON)
	}
}

func TestItReturnsAnErrorForNonJson(t *testing.T) {
	if _, err := IsCompatible("nonsense", "not json"); err == nil {
		t.Error("Expected an error to be returned for bad json")
	}
}

func shouldBeEqualErrorMsg(t *testing.T, a, b string) {
	t.Errorf("%s should be equal to %s", a, b)
}
