package jsonequaliser

import (
	"testing"
)

const simpleJSON = `{"firstname": "chris", "lastname": "james", "age": 30}`
const comparableJSON = `{"firstname": "christopher", "lastname": "james", "age": 15}`
const notSimilarJSON = `{"foo":"bar"}`

func TestItKnowsTheSameJSONIsCompatible(t *testing.T) {
	assertCompatible(t, simpleJSON, simpleJSON)
}

func TestItKnowsStructurallySameJSONIsCompatible(t *testing.T) {
	assertCompatible(t, simpleJSON, comparableJSON)
}

func TestItKnowsDifferentJSONIsIncompatible(t *testing.T) {
	assertIncompatible(t, simpleJSON, notSimilarJSON)
}

func TestItKnowsHowToHandleArrays(t *testing.T) {
	JSONWithArray := `{"foo": ["baz"]}`
	comparableJSONWithArray := `{"foo": ["bar"]}`
	assertCompatible(t, JSONWithArray, comparableJSONWithArray)
}

func TestItDoesntMindSuperflousFieldsInB(t *testing.T) {
	extraJSON := `{"firstname":"frank", "lastname": "sinatra", "extra field": "blue", "age":70}`
	assertCompatible(t, simpleJSON, extraJSON)
}

func TestItReturnsAnErrorForNonJson(t *testing.T) {
	if _, err := IsCompatible("nonsense", "not json"); err == nil {
		t.Error("Expected an error to be returned for bad json")
	}
}

func TestFloatingPoints(t *testing.T) {
	floatingJSONa := `{"x": 3.14, "y": "not"}`
	floatingJSONb := `{"x": "three", "y": "not"}`

	assertIncompatible(t, floatingJSONa, floatingJSONb)
}

// func TestNestedStructures(t *testing.T) {
// 	a := `{"hello": [{"x": 1, "y": "a"},{"x": 2, "y": "b"}]`
// 	b := `{"hello": [{"x": 10, "y": "b"}]`
// 	assertCompatible(t, a, b)
// }

func assertCompatible(t *testing.T, a, b string) {
	if compatible, _ := IsCompatible(a, b); !compatible {
		t.Errorf("%s should be compatible with %s", a, b)
	}
}

func assertIncompatible(t *testing.T, a, b string) {
	if compatible, _ := IsCompatible(a, b); compatible {
		t.Errorf("%s should not be compatible with %s", a, b)
	}
}
