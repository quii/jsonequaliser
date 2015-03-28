package jsonequaliser

import (
	"testing"
)

const simpleJSON = `{"firstname": "chris", "lastname": "james", "age": 30}`
const comparableJSON = `{"firstname": "christopher", "lastname": "james", "age": 15}`
const notSimilarJSON = `{"foo":"bar"}`

func TestItKnowsTheSameJSONIsCompatible(t *testing.T) {
	if isCompatible, _ := IsCompatible(simpleJSON, simpleJSON); !isCompatible {
		shouldBeCompatible(t, simpleJSON, simpleJSON)
	}
}

func TestItKnowsStructurallySameJSONIsCompatible(t *testing.T) {
	if isCompatible, _ := IsCompatible(simpleJSON, comparableJSON); !isCompatible {
		shouldBeCompatible(t, simpleJSON, comparableJSON)
	}
}

func TestItKnowsDifferentJSONIsDifferent(t *testing.T) {
	if equal, _ := IsCompatible(simpleJSON, notSimilarJSON); equal {
		shouldntBeCompatible(t, simpleJSON, notSimilarJSON)
	}
}

func TestItKnowsHowToHandleArrays(t *testing.T) {
	JSONWithArray := `{"foo": ["baz"]}`
	comparableJSONWithArray := `{"foo": ["bar"]}`

	if equal, _ := IsCompatible(JSONWithArray, comparableJSONWithArray); !equal {
		shouldBeCompatible(t, JSONWithArray, comparableJSONWithArray)
	}
}

func TestItDoesntMindSuperflousFieldsInB(t *testing.T) {
	extraJSON := `{"firstname":"frank", "lastname": "sinatra", "eyecolour": "blue", "age":70}`

	if equal, _ := IsCompatible(simpleJSON, extraJSON); !equal {
		shouldBeCompatible(t, simpleJSON, extraJSON)
	}
}

func TestItReturnsAnErrorForNonJson(t *testing.T) {
	if _, err := IsCompatible("nonsense", "not json"); err == nil {
		t.Error("Expected an error to be returned for bad json")
	}
}

func TestItHandlesFloatingPoints(t *testing.T) {
	floatingJSONa := `{"x": 3.14, "y": "not"}`
	floatingJSONb := `{"x": "three", "y": "not"}`

	if compatible, _ := IsCompatible(floatingJSONa, floatingJSONb); compatible {
		shouldntBeCompatible(t, floatingJSONa, floatingJSONb)
	}
}

func shouldBeCompatible(t *testing.T, a, b string) {
	t.Errorf("%s should be equal to %s", a, b)
}

func shouldntBeCompatible(t *testing.T, a, b string) {
	t.Errorf("%s should not be equal to %s", a, b)
}
