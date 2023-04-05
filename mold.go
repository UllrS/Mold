// package mold implements support for dynamic JSON keys and pattern-matched object structure transformation
package mold

import (
	"encoding/json"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

var (
	WriteAttempt = "=-"
	WriteForce   = "=!"
	WriteHarsh   = "=="

	WriteAttemptAll = "<-"
	WriteForceAll   = "<!"
	WriteHarshAll   = "<<"
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
	fillMold := filler.Fill()
	fillMoldByte, err := json.Marshal(fillMold)
	return fillMoldByte, err
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

func (f *Filler) Fill() interface{} {
	if res, ok := f.Source.(map[string]interface{}); ok {
		return fillObj(res, f.Mold.(map[string]interface{}), f.Source)
	} else if res, ok := f.Source.([]interface{}); ok {
		return fillList(res, f.Mold.([]interface{}), f.Source)
	} else {
		return nil
	}
}
