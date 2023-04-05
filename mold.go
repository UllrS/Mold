package mold

import (
	"encoding/json"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

var (
	WriteAttempt = "=-" //Передает значения в форму, только если оно найдено в источнике, соответствует тип значения и находится на одном уровне с формой, иначе оставляет значение формы. Дочерние объекты заполняются рекурсивно и только динамические значения.
	WriteForce   = "=!" //Передает значения в форму, только если оно найдено в источнике, соответствует тип значения и находится на одном уровне с формой, иначе передается значение null. Дочерние объекты заполняются рекурсивно и только динамические значения.
	WriteHarsh   = "==" //Передает значение в форму, вне зависимости от типов. Дочерние объекты передаются полностью, без рекурсивного анализа.  Рекомендуем использовать, если вы ожидаете в значении примитив неизвестного типа.

	WriteAttemptAll = "<-" //Поиск значения по ключу по всему источнику, включая дочерние и родительские объекты. Передает значение в форму, вне зависимости от типов. Дочерние объекты и массивы так же обрабатываются для поиска динамических ключей. При отсутствии значения в источнике возвращается значение в форме
	WriteForceAll   = "<!" //Поиск значения по ключу по всему источнику, включая дочерние и родительские объекты. Передает значение в форму, вне зависимости от типов. Дочерние объекты и массивы так же обрабатываются для поиска динамических ключей. При отсутствии значения в источнике возвращается значение null
	WriteHarshAll   = "<<" //Поиск значения по ключу по всему источнику, включая дочерние и родительские объекты. Передает значение в форму, вне зависимости от типов. Дочерние объекты и массивы передаются полностью, без рекурсивного анализа. Рекомендуем использовать если вы ожидаете в значении примитив неизвестного типа.
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

func (f *Filler) Fill() interface{} {
	if res, ok := f.Source.(map[string]interface{}); ok {
		return fillObj(res, f.Mold.(map[string]interface{}), f.Source)
	} else if res, ok := f.Source.([]interface{}); ok {
		return fillList(res, f.Mold.([]interface{}), f.Source)
	} else {
		return nil
	}
}
