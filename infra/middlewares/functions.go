package middleware

import "fmt"

type MaskerFunc func(value string, skipSanitization bool) string

func SanitizeResponseBody(responseBody any, fieldName string, maskerFunc MaskerFunc, skip bool) {
	if responseBodyAsMap, success := responseBody.(map[string]interface{}); success {
		sanitizeMap(responseBodyAsMap, fieldName, maskerFunc, skip)
	} else if responseBodyAsSlice, success := responseBody.([]interface{}); success {
		for _, item := range responseBodyAsSlice {
			if itemAsMap, success := item.(map[string]interface{}); success {
				sanitizeMap(itemAsMap, fieldName, maskerFunc, skip)
			}
		}
	}
}

func sanitizeMap(responseMap map[string]interface{}, fieldName string, maskerFunc MaskerFunc, skip bool) {
	for _, item := range responseMap {
		if compoundMap, success := item.(map[string]interface{}); success {
			sanitizeMap(compoundMap, fieldName, maskerFunc, skip)
		} else if compoundSlice, success := item.([]interface{}); success {
			for _, compoundItem := range compoundSlice {
				if compoundItemAsMap, success := compoundItem.(map[string]interface{}); success {
					sanitizeMap(compoundItemAsMap, fieldName, maskerFunc, skip)
				}
			}
		} else {
			if fieldValue, contains := responseMap[fieldName]; contains {
				responseMap[fieldName] = maskerFunc(fmt.Sprintf("%v", fieldValue), skip)
			}
		}
	}
}
