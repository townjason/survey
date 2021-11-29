package hit

import (
	"api/util/log"
	"reflect"
	"strconv"
	"unicode"
)

func callFn(f interface{}) interface{} {
	if f != nil {
		t := reflect.TypeOf(f)
		if t.Kind() == reflect.Func && t.NumIn() == 0 {
			function := reflect.ValueOf(f)
			in := make([]reflect.Value, 0)
			out := function.Call(in)
			if num := len(out); num > 0 {
				list := make([]interface{}, num)
				for i, value := range out {
					list[i] = value.Interface()
				}
				if num == 1 {
					return list[0]
				}
				return list
			}
			return nil
		}
	}
	return f
}

func isZero(f interface{}) bool  {
	v := reflect.ValueOf(f)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		str := v.String()
		if str == "" {
			return true
		}
		zero, err := strconv.ParseFloat(str, 10)
		log.Error(err)
		if zero == 0 && err == nil {
			return true
		}
		boolean, err := strconv.ParseBool(str)
		log.Error(err)
		return boolean == false && err == nil
	default:
		return false
	}
}

// TrimZero 去掉\u0000字符
func TrimZero(s string) string {
	str := make([]rune, 0, len(s))
	for _, v := range []rune(s) {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			continue
		}

		str = append(str, v)
	}
	return string(str)
}

// If - (a ? b : c) Or (a && b)
func If(args ...interface{}) interface{} {
	var condition = callFn(args[0])
	if len(args) == 1 {
		return condition
	}
	var trueVal = args[1]
	var falseVal interface{}
	if len(args) > 2 {
		falseVal = args[2]
	} else {
		falseVal = nil
	}
	if condition == nil {
		return callFn(falseVal)
	} else if v, ok := condition.(bool); ok {
		if v == false {
			return callFn(falseVal)
		}
	} else if isZero(condition) {
		return callFn(falseVal)
	} else if v, ok := condition.(error); ok {
		if v != nil {
			return condition
		}
	}
	return callFn(trueVal)
}

