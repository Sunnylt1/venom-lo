package venom

import (
	"fmt"
)

var (
	executors = map[string]Executor{}
)

// RegisterExecutor register Test Executors
func RegisterExecutor(name string, e Executor) {
	executors[name] = e
}

// getExecutorWrap initializes a test by name
// no type -> exec is default
func getExecutorWrap(t map[string]interface{}) (*executorWrap, error) {

	var name string
	var retry, delay int

	if itype, ok := t["type"]; ok {
		name = fmt.Sprintf("%s", itype)
	}

	if name == "" {
		name = "exec"
	}

	retry, errRetry := getAttrInt(t, "retry")
	if errRetry != nil {
		return nil, errRetry
	}
	delay, errDelay := getAttrInt(t, "delay")
	if errDelay != nil {
		return nil, errDelay
	}

	if e, ok := executors[name]; ok {
		ew := &executorWrap{
			executor: e,
			retry:    retry,
			delay:    delay,
		}
		return ew, nil
	}

	return nil, fmt.Errorf("type '%s' is not implemented", name)
}

func getAttrInt(t map[string]interface{}, name string) (int, error) {
	var out int
	if i, ok := t["retry"]; ok {
		var ok bool
		out, ok = i.(int)
		if !ok {
			return -1, fmt.Errorf("attribute %s '%s' is not an integer", name, i)
		}
	}
	return out, nil
}
