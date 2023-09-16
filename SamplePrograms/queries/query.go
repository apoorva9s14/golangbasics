package queries

import "errors"

func nestedIfVsIfElseIf(param1 string, param2 bool) (string, error) {
	// param1 is a mandatory parameter if param2 is enabled, else it is not mandatory
	if param1 == "" && param2 {
		return param1, errors.New("param1 is empty in payload")
	} else if param1 == "" && !param2 {
		return param1, nil
	}

	//vs
	if param1 == "" {
		if param2 {
			return param1, errors.New("param1 is empty in payload")
		}
	}
	return param1, nil
}
