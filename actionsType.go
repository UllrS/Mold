package mold

import (
	"fmt"
)

func actionMap(mold, obj map[string]interface{}, key *Key, source interface{}) interface{} {
	if key.Separator == WriteAttempt {
		if objValue, ok := obj[key.FillerKey].(map[string]interface{}); ok {
			return fillObj(objValue, mold, source)
		}
		return fillObj(mold, mold, source)

	} else if key.Separator == WriteForce {
		if objValue, ok := obj[key.FillerKey].(map[string]interface{}); ok {
			return fillObj(objValue, mold, source)
		}
		return nil

	} else if key.Separator == WriteHarsh {
		return obj[key.FillerKey]
	} else if key.Separator == WriteAttemptAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, ok := newValue.(map[string]interface{}); ok {
				return fillObj(objValue, mold, source)
			}
		}
		return fillObj(mold, mold, source)

	} else if key.Separator == WriteForceAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, suok := newValue.(map[string]interface{}); suok {
				return fillObj(objValue, mold, source)
			}
		}
		return nil
	} else if key.Separator == WriteHarshAll {
		newValue, _ := findFieldWide(source, key.FillerKey, mold)
		return newValue
	}
	return mold
}
func actionList(mold []interface{}, obj map[string]interface{}, key *Key, source interface{}) interface{} {
	if key.Separator == WriteAttempt {
		if objValue, ok := obj[key.FillerKey].([]interface{}); ok {
			return fillList(objValue, mold, source)
		}
		return fillList(mold, mold, source)
	} else if key.Separator == WriteForce {
		if objValue, ok := obj[key.FillerKey].([]interface{}); ok {
			return fillList(objValue, mold, source)
		}
		return []interface{}{}

	} else if key.Separator == WriteHarsh {
		return obj[key.FillerKey]
	} else if key.Separator == WriteAttemptAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, ok := newValue.([]interface{}); ok {
				return fillList(objValue, mold, source)
			}
		}
		return fillList(mold, mold, source)
	} else if key.Separator == WriteForceAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, ok := newValue.([]interface{}); ok {
				return fillList(objValue, mold, source)
			}
		}
		return []interface{}{}
	} else if key.Separator == WriteHarshAll {
		newValue, _ := findFieldWide(source, key.FillerKey, mold)
		return newValue
	}
	return nil
}
func actionBool(mold bool, obj map[string]interface{}, key *Key, source interface{}) interface{} {
	if key.Separator == WriteAttempt {
		if objValue, ok := obj[key.FillerKey].(bool); ok {
			return objValue
		}
		return mold
	} else if key.Separator == WriteForce {
		if objValue, ok := obj[key.FillerKey].(bool); ok {
			return objValue
		}
		return nil
	} else if key.Separator == WriteHarsh {
		return obj[key.FillerKey]
	} else if key.Separator == WriteAttemptAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, subok := newValue.(bool); subok {
				return objValue
			}
		}
		return mold
	} else if key.Separator == WriteForceAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, ok := newValue.(bool); ok {
				return objValue
			}
		}
		return nil
	} else if key.Separator == WriteHarshAll {
		newValue, _ := findFieldWide(source, key.FillerKey, mold)
		return newValue
	}
	return nil
}
func actionFloat(mold float64, obj map[string]interface{}, key *Key, source interface{}) interface{} {
	if key.Separator == WriteAttempt {
		if objValue, ok := obj[key.FillerKey]; ok {
			fmt.Println("FLOAT 64 ", key.FillerKey, ":", mold, ":", obj[key.FillerKey])
			if _, ok := objValue.(float64); ok || (objValue == nil) {
				return objValue
			}
		}
		return mold
	} else if key.Separator == WriteForce {
		if objValue, ok := obj[key.FillerKey].(float64); ok {
			return objValue
		}
		return nil
	} else if key.Separator == WriteHarsh {
		return obj[key.FillerKey]
	} else if key.Separator == WriteAttemptAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, subok := newValue.(float64); subok {
				return objValue
			}
		}
		return mold
	} else if key.Separator == WriteForceAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, ok := newValue.(float64); ok {
				return objValue
			}
		}
		return nil
	} else if key.Separator == WriteHarshAll {
		newValue, _ := findFieldWide(source, key.FillerKey, mold)
		return newValue
	}
	return nil

}
func actionString(mold string, obj map[string]interface{}, key *Key, source interface{}) interface{} {
	if key.Separator == WriteAttempt {
		if objValue, ok := obj[key.FillerKey].(string); ok {
			return objValue
		}
		return mold
	} else if key.Separator == WriteForce {
		if objValue, ok := obj[key.FillerKey].(string); ok {
			return objValue
		}
		return nil

	} else if key.Separator == WriteHarsh {
		return obj[key.FillerKey]
	} else if key.Separator == WriteAttemptAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, subok := newValue.(string); subok {
				return objValue
			}
		}
		return mold
	} else if key.Separator == WriteForceAll {
		if newValue, ok := findFieldWide(source, key.FillerKey, mold); ok {
			if objValue, ok := newValue.(string); ok {
				return objValue
			}
		}
		return nil
	} else if key.Separator == WriteHarshAll {
		newValue, _ := findFieldWide(source, key.FillerKey, mold)
		return newValue
	}
	return nil
}
func actionNil(obj map[string]interface{}, key *Key, source interface{}) interface{} {
	if key.Separator == WriteAttempt {
		return obj[key.FillerKey]
	} else if key.Separator == WriteForce {
		return obj[key.FillerKey]
	} else if key.Separator == WriteHarsh {
		return obj[key.FillerKey]
	} else if key.Separator == WriteAttemptAll {
		newValue, _ := findFieldWide(source, key.FillerKey, nil)
		return newValue
	} else if key.Separator == WriteForceAll {
		newValue, _ := findFieldWide(source, key.FillerKey, nil)
		return newValue
	} else if key.Separator == WriteHarshAll {
		newValue, _ := findFieldWide(source, key.FillerKey, nil)
		return newValue
	}
	return nil
}
