package jsonequaliser

import (
	"encoding/json"
)

// IsCompatible checks that two json strings are structurally the same so that they are compatible. The first string should be your "correct" json, if there are extra fields in B then they will still be seen as compatible
func IsCompatible(a, b string) (bool, error) {

	aMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(a), &aMap); err != nil {
		return false, err
	}

	bMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(b), &bMap); err != nil {
		return false, err
	}

	return isStructurallyTheSame(aMap, bMap)

}

func isStructurallyTheSame(a, b map[string]interface{}) (bool, error) {
	for k, v := range a {

		if b[k] == nil {
			return false, nil
		}

		switch v.(type) {
		case string:
			_, isString := b[k].(string)
			if !isString {
				return false, nil
			}
		case int:
			_, isInt := b[k].(int)
			if !isInt {
				return false, nil
			}
		case bool:
			_, isBool := b[k].(bool)
			if !isBool {
				return false, nil
			}

		case []interface{}:
			aLeaf, _ := a[k].(map[string]interface{})
			bLeaf, _ := b[k].(map[string]interface{})
			return isStructurallyTheSame(aLeaf, bLeaf)
		}

	}

	return true, nil
}

/*
{"firstname": "chris", "lastname": "james", "age": 30}
{"firstname":"frank", "lastname": "sinatra", "eyecolour": "blue", "age":70}
*/