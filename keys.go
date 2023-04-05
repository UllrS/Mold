package mold

import (
	"strings"
)

type Key struct {
	FillerKey string
	MoldKey   string
	Separator string
}

func parsingSeparator(key string) (string, error) {
	if strings.Contains(key, WriteAttempt) {
		return WriteAttempt, nil
	} else if strings.Contains(key, WriteForce) {
		return WriteForce, nil
	} else if strings.Contains(key, WriteHarsh) {
		return WriteHarsh, nil
	} else if strings.Contains(key, WriteAttemptAll) {
		return WriteAttemptAll, nil
	} else if strings.Contains(key, WriteForceAll) {
		return WriteForceAll, nil
	} else if strings.Contains(key, WriteHarshAll) {
		return WriteHarshAll, nil
	} else {
		return "", ErrNotFound
	}
}
func splitKey(key string) (*Key, error) {
	separator, err := parsingSeparator(key)
	if err != nil {
		return nil, err
	}
	parseKeys := strings.Split(key, separator)
	moldKey, fillerKey := parseKeys[0], parseKeys[1]
	if fillerKey == "" {
		fillerKey = moldKey
	}
	keym := Key{
		MoldKey:   moldKey,
		FillerKey: fillerKey,
		Separator: separator,
	}
	return &keym, nil
}
