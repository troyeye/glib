package cast

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
)

var errNegative = errors.New("can't cast negative value")

// ToBoolE 转换 interface 到 bool
func ToBoolE(i interface{}) (bool, error) {
	i = indirect(i)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("can't cast %#v of type %T to bool", i, i)
	}
}

// ToInt64E 转换 interface 到 int64
func ToInt64E(i interface{}) (int64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int64(s), nil
	case int64:
		return s, nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint64:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case float64:
		return int64(s), nil
	case float32:
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("can't cast %#v of type %T to int64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to int64", i, i)
	}
}

// ToInt32E 转换 interface 到 int32
func ToInt32E(i interface{}) (int32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int32(s), nil
	case int64:
		return int32(s), nil
	case int32:
		return s, nil
	case int16:
		return int32(s), nil
	case int8:
		return int32(s), nil
	case uint:
		return int32(s), nil
	case uint64:
		return int32(s), nil
	case uint32:
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint8:
		return int32(s), nil
	case float64:
		return int32(s), nil
	case float32:
		return int32(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 32)
		if err == nil {
			return int32(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v of type %T to int32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to int32", i, i)
	}
}

// ToInt16E 转换 interface 到 int16
func ToInt16E(i interface{}) (int16, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int16(s), nil
	case int64:
		return int16(s), nil
	case int32:
		return int16(s), nil
	case int16:
		return s, nil
	case int8:
		return int16(s), nil
	case uint:
		return int16(s), nil
	case uint64:
		return int16(s), nil
	case uint32:
		return int16(s), nil
	case uint16:
		return int16(s), nil
	case uint8:
		return int16(s), nil
	case float64:
		return int16(s), nil
	case float32:
		return int16(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 16)
		if err == nil {
			return int16(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v of type %T to int16", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to int16", i, i)
	}
}

// ToInt8E 转换 interface 到 int8
func ToInt8E(i interface{}) (int8, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int8(s), nil
	case int64:
		return int8(s), nil
	case int32:
		return int8(s), nil
	case int16:
		return int8(s), nil
	case int8:
		return s, nil
	case uint:
		return int8(s), nil
	case uint64:
		return int8(s), nil
	case uint32:
		return int8(s), nil
	case uint16:
		return int8(s), nil
	case uint8:
		return int8(s), nil
	case float64:
		return int8(s), nil
	case float32:
		return int8(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 8)
		if err == nil {
			return int8(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v of type %T to int8", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to int8", i, i)
	}
}

// ToIntE 转换 interface 到 int
func ToIntE(i interface{}) (int, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		return int(s), nil
	case float32:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v of type %T to int", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to int", i, i)
	}
}

// ToUint64E 转换 interface 到 uint64
func ToUint64E(i interface{}) (uint64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("can't cast %#v to uint64: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegative
		}
		return uint64(s), nil
	case int64:
		if s < 0 {
			return 0, errNegative
		}
		return uint64(s), nil
	case int32:
		if s < 0 {
			return 0, errNegative
		}
		return uint64(s), nil
	case int16:
		if s < 0 {
			return 0, errNegative
		}
		return uint64(s), nil
	case int8:
		if s < 0 {
			return 0, errNegative
		}
		return uint64(s), nil
	case uint:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case float32:
		if s < 0 {
			return 0, errNegative
		}
		return uint64(s), nil
	case float64:
		if s < 0 {
			return 0, errNegative
		}
		return uint64(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to uint64", i, i)
	}
}

// ToUint32E 转换 interface 到 uint32
func ToUint32E(i interface{}) (uint32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 32)
		if err == nil {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v to uint32: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegative
		}
		return uint32(s), nil
	case int64:
		if s < 0 {
			return 0, errNegative
		}
		return uint32(s), nil
	case int32:
		if s < 0 {
			return 0, errNegative
		}
		return uint32(s), nil
	case int16:
		if s < 0 {
			return 0, errNegative
		}
		return uint32(s), nil
	case int8:
		if s < 0 {
			return 0, errNegative
		}
		return uint32(s), nil
	case uint:
		return uint32(s), nil
	case uint64:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint16:
		return uint32(s), nil
	case uint8:
		return uint32(s), nil
	case float64:
		if s < 0 {
			return 0, errNegative
		}
		return uint32(s), nil
	case float32:
		if s < 0 {
			return 0, errNegative
		}
		return uint32(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to uint32", i, i)
	}
}

// ToUint16E 转换 interface 到 uint16
func ToUint16E(i interface{}) (uint16, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 16)
		if err == nil {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v to uint16: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegative
		}
		return uint16(s), nil
	case int64:
		if s < 0 {
			return 0, errNegative
		}
		return uint16(s), nil
	case int32:
		if s < 0 {
			return 0, errNegative
		}
		return uint16(s), nil
	case int16:
		if s < 0 {
			return 0, errNegative
		}
		return uint16(s), nil
	case int8:
		if s < 0 {
			return 0, errNegative
		}
		return uint16(s), nil
	case uint:
		return uint16(s), nil
	case uint64:
		return uint16(s), nil
	case uint32:
		return uint16(s), nil
	case uint16:
		return s, nil
	case uint8:
		return uint16(s), nil
	case float64:
		if s < 0 {
			return 0, errNegative
		}
		return uint16(s), nil
	case float32:
		if s < 0 {
			return 0, errNegative
		}
		return uint16(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to uint16", i, i)
	}
}

// ToUint8E 转换 interface 到 uint8
func ToUint8E(i interface{}) (uint8, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 8)
		if err == nil {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v to uint8: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegative
		}
		return uint8(s), nil
	case int64:
		if s < 0 {
			return 0, errNegative
		}
		return uint8(s), nil
	case int32:
		if s < 0 {
			return 0, errNegative
		}
		return uint8(s), nil
	case int16:
		if s < 0 {
			return 0, errNegative
		}
		return uint8(s), nil
	case int8:
		if s < 0 {
			return 0, errNegative
		}
		return uint8(s), nil
	case uint:
		return uint8(s), nil
	case uint64:
		return uint8(s), nil
	case uint32:
		return uint8(s), nil
	case uint16:
		return uint8(s), nil
	case uint8:
		return s, nil
	case float64:
		if s < 0 {
			return 0, errNegative
		}
		return uint8(s), nil
	case float32:
		if s < 0 {
			return 0, errNegative
		}
		return uint8(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to uint8", i, i)
	}
}

// ToUintE 转换 interface 到 uint
func ToUintE(i interface{}) (uint, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 0)
		if err == nil {
			return uint(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v to uint: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegative
		}
		return uint(s), nil
	case int64:
		if s < 0 {
			return 0, errNegative
		}
		return uint(s), nil
	case int32:
		if s < 0 {
			return 0, errNegative
		}
		return uint(s), nil
	case int16:
		if s < 0 {
			return 0, errNegative
		}
		return uint(s), nil
	case int8:
		if s < 0 {
			return 0, errNegative
		}
		return uint(s), nil
	case uint:
		return s, nil
	case uint64:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint8:
		return uint(s), nil
	case float64:
		if s < 0 {
			return 0, errNegative
		}
		return uint(s), nil
	case float32:
		if s < 0 {
			return 0, errNegative
		}
		return uint(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to uint", i, i)
	}
}

// ToFloat64E 转换 interface 到 float64
func ToFloat64E(i interface{}) (float64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("can't cast %#v of type %T to float64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to float64", i, i)
	}
}

// ToFloat32E 转换 interface 到 float32
func ToFloat32E(i interface{}) (float32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case float64:
		return float32(s), nil
	case float32:
		return s, nil
	case int:
		return float32(s), nil
	case int64:
		return float32(s), nil
	case int32:
		return float32(s), nil
	case int16:
		return float32(s), nil
	case int8:
		return float32(s), nil
	case uint:
		return float32(s), nil
	case uint64:
		return float32(s), nil
	case uint32:
		return float32(s), nil
	case uint16:
		return float32(s), nil
	case uint8:
		return float32(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return float32(v), nil
		}
		return 0, fmt.Errorf("can't cast %#v of type %T to float32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't cast %#v of type %T to float32", i, i)
	}
}

// ToStringE 转换 interface 到 string
func ToStringE(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(s, 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("can't cast %#v of type %T to string", i, i)
	}
}

// 参考Go SDK html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
// 多次解引用后，返回一个base类型 或者 nil
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

var (
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

// 参考Go SDK html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
// 多次解引用以后，返回一个base类型 或者 nil 或者 fmt.Stringer的实例 或者 error
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
