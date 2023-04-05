// package mold implements support for dynamic JSON keys and pattern-matched object structure transformation
package mold

import (
	"encoding/json"
	"errors"
	"reflect"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrNotSupported   = errors.New("source type not supported")
	ErrdifferentTypes = errors.New("source and mold have different types")
)

var (
	WriteAttempt = "=-" //Passes a value to the form only if it is found in the source, matches the value type, and is on the same level as the form, otherwise leaves the form's value. Child objects are filled recursively and only dynamic values.
	WriteForce   = "=!" //Passes a value to the form only if it is found in the source, matches the value type, and is at the same level as the form, otherwise null is passed. Child objects are filled recursively and only dynamic values.
	WriteHarsh   = "==" //Passes a value to the form, regardless of types. Child objects are passed in their entirety, without recursive parsing. Recommended if you are expecting a primitive of unknown type in the value.

	WriteAttemptAll = "<-" //Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are also processed to find dynamic keys. If there is no value in the source, a value is returned in the form
	WriteForceAll   = "<!" //Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are also processed to find dynamic keys. If there is no value in the source, null is returned.
	WriteHarshAll   = "<<" //Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are passed in full, without recursive parsing. It is recommended to use if you are expecting a primitive of an unknown type in the value.
)

func Fill(source, mold []byte) ([]byte, error) {
	sourceEq, err := typeReduction(source)
	if err != nil {
		return nil, err
	}
	moldEq, err := typeReduction(mold)
	if err != nil {
		return nil, err
	}
	filler := Filler{
		Source: sourceEq,
		Mold:   moldEq,
	}
	fillMold, err := filler.Fill()
	if err != nil {
		return nil, err
	}
	fillMoldByte, err := json.Marshal(fillMold)
	if err != nil {
		return nil, err
	}
	return fillMoldByte, nil
}
func typeReduction(data []byte) (interface{}, error) {
	var equivalent interface{}
	err := json.Unmarshal(data, &equivalent)
	if err != nil {
		return nil, err
	}
	return equivalent, nil
}

type Filler struct {
	Source interface{}
	Mold   interface{}
}

func (f *Filler) Fill() (interface{}, error) {
	if reflect.TypeOf(f.Source) != reflect.TypeOf(f.Mold) {
		return nil, ErrdifferentTypes
	}
	if res, ok := f.Source.(map[string]interface{}); ok {
		return fillObj(res, f.Mold.(map[string]interface{}), f.Source), nil
	} else if res, ok := f.Source.([]interface{}); ok {
		return fillList(res, f.Mold.([]interface{}), f.Source), nil
	} else {
		return nil, ErrNotSupported
	}
}
