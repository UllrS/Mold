package mold

import (
	"reflect"
)

func fillObj(obj, mold map[string]interface{}, source interface{}) map[string]interface{} {
	newJson := map[string]interface{}{}

	for jkey, val := range mold {
		key, err := splitKey(jkey)
		if err != nil {
			newJson[jkey] = val
			continue
		} else if val == nil {
			newJson[key.MoldKey] = actionNil(obj, key, source)
		} else if res, ok := val.(map[string]interface{}); ok {
			newJson[key.MoldKey] = actionMap(res, obj, key, source)
		} else if res, ok := val.([]interface{}); ok {
			newJson[key.MoldKey] = actionList(res, obj, key, source)
		} else if res, ok := val.(string); ok {
			newJson[key.MoldKey] = actionString(res, obj, key, source)
		} else if res, ok := val.(bool); ok {
			newJson[key.MoldKey] = actionBool(res, obj, key, source)
		} else if res, ok := val.(float64); ok {
			newJson[key.MoldKey] = actionFloat(res, obj, key, source)
		}
	}
	return newJson
}
func fillList(sourceList []interface{}, moldList []interface{}, source interface{}) []interface{} {
	if len(moldList) < 1 || len(sourceList) < 1 {
		return sourceList
	}
	moldFirst := moldList[0]
	var newList []interface{}
	for _, val := range sourceList {
		if reflect.TypeOf(val) != reflect.TypeOf(moldFirst) {
			continue
		}
		if res, ok := val.(map[string]interface{}); ok {
			if mold, ok := moldFirst.(map[string]interface{}); ok {
				newJson := fillObj(res, mold, source)
				newList = append(newList, newJson)
			}
		} else if res, ok := val.([]interface{}); ok {
			if mold, ok := moldFirst.([]interface{}); ok {
				newList = append(newList, fillList(res, mold, source))
			}
		} else if res, ok := val.(bool); ok {
			if _, ok := moldFirst.(bool); ok {
				newList = append(newList, res)
			}
		} else if res, ok := val.(string); ok {
			if _, ok := moldFirst.(string); ok {
				newList = append(newList, res)
			}
		} else if res, ok := val.(float64); ok {
			if _, ok := moldFirst.(float64); ok {
				newList = append(newList, res)
			}
		} else if val == nil {
			newList = append(newList, val)
		}
	}
	return newList
}

func findFieldWide(obj interface{}, field string, defValue interface{}) (interface{}, bool) {
	if sub, ok := obj.(map[string]interface{}); ok {
		result, ok := findFieldObj(sub, field, defValue)
		return result, ok
	} else if sub, ok := obj.([]interface{}); ok {
		result, ok := findFieldList(sub, field, defValue)
		return result, ok
	} else {
		return nil, false
	}

}
func findFieldObj(obj map[string]interface{}, field string, defValue interface{}) (interface{}, bool) {
	for key, val := range obj {
		if key == field {
			return val, true
		} else if sub, ok := val.(map[string]interface{}); ok {
			if f, change := findFieldObj(sub, field, defValue); change {
				return f, change
			}
		} else if sub, ok := val.([]interface{}); ok {
			for _, el := range sub {
				if subi, ok := el.(map[string]interface{}); ok {
					if f, change := findFieldObj(subi, field, defValue); change {
						return f, change
					}
				} else if subi, ok := el.([]interface{}); ok {
					if f, change := findFieldList(subi, field, defValue); change {
						return f, change
					}
				}
			}
		}
	}
	return defValue, false
}
func findFieldList(list []interface{}, field string, defValue interface{}) (interface{}, bool) {
	for _, val := range list {
		if sub, ok := val.(map[string]interface{}); ok {
			if f, change := findFieldObj(sub, field, defValue); change {
				return f, change
			}
		} else if sub, ok := val.([]interface{}); ok {
			for _, el := range sub {
				if subi, ok := el.(map[string]interface{}); ok {
					if f, change := findFieldObj(subi, field, defValue); change {
						return f, change
					}
				} else if subi, ok := el.([]interface{}); ok {
					if f, change := findFieldList(subi, field, defValue); change {
						return f, change
					}
				}
			}
		}
	}
	return defValue, false

}
