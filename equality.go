package jsonequaliser

import (
	"encoding/json"
	"fmt"
	"reflect"
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

		case float64:
			_, isFloat := b[k].(float64)
			if !isFloat {
				return false, nil
			}

		case interface{}:
			aArr, aIsArray := a[k].([]interface{})

			if aIsArray {
				bArr, bIsArray := b[k].([]interface{})

				if !bIsArray {
					return false, nil
				}

				aLeaf, aIsMap := aArr[0].(map[string]interface{})
				bLeaf, bIsMap := bArr[0].(map[string]interface{})

				if aIsMap && bIsMap {
					return isStructurallyTheSame(aLeaf, bLeaf)
				} else if aIsMap && !bIsMap {
					return false, nil
				} else {
					return reflect.TypeOf(aArr[0]) == reflect.TypeOf(bArr[0]), nil
				}
			}

			aLeaf, aIsMap := a[k].(map[string]interface{})
			bLeaf, bIsMap := b[k].(map[string]interface{})

			if aIsMap && bIsMap {
				return isStructurallyTheSame(aLeaf, bLeaf)
			}
			return false, nil

		default:
			return false, fmt.Errorf("Unmatched type of json found, got a %v", reflect.TypeOf(v))
		}

	}

	return true, nil
}
