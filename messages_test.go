package jsonequaliser

import "testing"

func TestItReturnsAllProblems(t *testing.T) {

	A := `{"x": true, "y": 10, "z": "hello"}`
	B := `{"x": 10, "y": true}`

	messages, err := IsCompatible(A, B)

	if len(messages) == 0 || err != nil {
		t.Fatal("Should be incompatible and not fail", err)
	}

	if len(messages) != 3 {
		t.Fatal("Expected 3 error messages, but got", messages)
	}
}

func TestTheMessagesAreGreatWithNestedFields(t *testing.T) {
	A := `{"x": { "y": true}}`
	B := `{"x": { "y": 10}}`
	assertMessage(t, A, B, "x->y", msgNotBool)
}

func TestTheMessagesAreGreatWithDeeplyNestedFields(t *testing.T) {
	A := `{"x": { "y": { "z": true}}}`
	B := `{"x": { "y": { "z": 10}}}`
	assertMessage(t, A, B, "x->y->z", msgNotBool)
}

func TestItReportsEmptyArray(t *testing.T) {
	A := `{"x" : ["hello"]}`
	B := `{"x" : []}`
	assertMessage(t, A, B, "x", msgEmptyArray)
}

func TestItReportsDifferentArrayType(t *testing.T) {
	A := `{"x" : ["hello"]}`
	B := `{"x" : [1]}`
	assertMessage(t, A, B, "x", msgDifferentArrayType)
}

func TestItReportsNotAMap(t *testing.T) {
	A := `{"x" : {"y": true}}`
	B := `{"x" : true}`
	assertMessage(t, A, B, "x", msgNotMap)
}

func TestItReportsMissingFields(t *testing.T) {
	A := `{"x" : "hello"}`
	B := `{"y" : 1}`
	assertMessage(t, A, B, "x", msgFieldMissing)
}

func TestItReportsNonStrings(t *testing.T) {
	A := `{"x" : "hello"}`
	B := `{"x" : 1}`
	assertMessage(t, A, B, "x", msgNotString)
}

func TestItReportsNonBools(t *testing.T) {
	A := `{"x" : true}`
	B := `{"x" : 1}`
	assertMessage(t, A, B, "x", msgNotBool)
}

func TestItReportsNonFloats(t *testing.T) {
	A := `{"x" : 1}`
	B := `{"x" : true}`
	assertMessage(t, A, B, "x", msgNotFloat)
}

func TestItReportsEmptyArrays(t *testing.T) {
	A := `[]`
	B := `[]`
	assertMessage(t, A, B, "rootArray", msgEmptyRootArray)
}

func assertMessage(t *testing.T, a, b, key, expectedMessage string) {
	messages, err := IsCompatible(a, b)

	if len(messages) == 0 || err != nil {
		t.Fatal("Should be incompatible and not fail", err)
	}

	if messages[key] != expectedMessage {
		t.Error("Message for", key, "was not", expectedMessage, "it was", messages[key])
	}
}
