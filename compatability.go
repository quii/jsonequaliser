package jsonequaliser

import (
	"fmt"
	"reflect"
)

type jsonNode map[string]interface{}

// IsCompatible checks that two json strings are structurally the same so that they are compatible. The first string should be your "correct" json, if there are extra fields in B then they will still be seen as compatible
func IsCompatible(a, b string) (messages map[string]string, err error) {

	aMap, err := getJSONNodeFromString(a)
	bMap, err := getJSONNodeFromString(b)

	if err != nil {
		return
	}

	return isStructurallyTheSame(aMap, bMap, make(map[string]string))
}

var (
	msgFieldMissing       = "Missing field"
	msgNotString          = "Field is not a string in other JSON"
	msgNotBool            = "Field is not a boolean in other JSON"
	msgNotFloat           = "Field is not a float in other JSON"
	msgEmptyArray         = "Array in other JSON is empty so I cant check"
	msgNotMap             = "Not a map in other JSON"
	msgDifferentArrayType = "Type of array is different"
)

func isStructurallyTheSame(a, b jsonNode, messages map[string]string) (map[string]string, error) {
	for jsonFieldName, v := range a {

		if fieldMissingIn(b, jsonFieldName) {
			messages[jsonFieldName] = msgFieldMissing
		}

		if a[jsonFieldName] == nil {
			break
		}

		switch v.(type) {
		case string:
			if !isString(b, jsonFieldName) {
				messages[jsonFieldName] = msgNotString
				break
			}
		case bool:
			if !isBool(b, jsonFieldName) {
				messages[jsonFieldName] = msgNotBool
				break
			}
		case float64:
			if !isFloat(b, jsonFieldName) {
				messages[jsonFieldName] = msgNotFloat
				break
			}

		case interface{}:

			aArr, aIsArray := a[jsonFieldName].([]interface{})

			bArr, bIsArray := b[jsonFieldName].([]interface{})

			if aIsArray && len(aArr) == 0 {
				break
			}

			if !bIsArray && aIsArray || aIsArray && len(bArr) == 0 {
				messages[jsonFieldName] = msgEmptyArray
				break
			}

			var aLeaf, bLeaf jsonNode
			var aIsMap, bIsMap bool

			if aIsArray && bIsArray {
				aLeaf, aIsMap = aArr[0].(map[string]interface{})
				bLeaf, bIsMap = bArr[0].(map[string]interface{})
			} else {
				aLeaf, aIsMap = a[jsonFieldName].(map[string]interface{})
				bLeaf, bIsMap = b[jsonFieldName].(map[string]interface{})
			}

			if aIsMap && bIsMap {
				messages, err := isStructurallyTheSame(aLeaf, bLeaf, messages)
				if err != nil {
					return messages, err
				}
				break
			} else if aIsMap && !bIsMap {
				messages[jsonFieldName] = msgNotMap
				break
			} else if reflect.TypeOf(aArr[0]) != reflect.TypeOf(bArr[0]) {
				messages[jsonFieldName] = msgDifferentArrayType
				break
			}
		default:
			return messages, fmt.Errorf("Unmatched type of json found, got a %v", reflect.TypeOf(v))
		}
	}

	return messages, nil
}
