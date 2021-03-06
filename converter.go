package utils

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrConvertFail = errors.New("convert data type is failure")
)

func Int(in interface{}) (int, error) {
	return ToInt(in)
}

func MustInt(in interface{}) int {
	val, _ := ToInt(in)
	return val
}

func ToInt(in interface{}) (iVal int, err error) {
	switch tVal := in.(type) {
	case nil:
		iVal = 0
	case int:
		iVal = tVal
	case int8:
		iVal = int(tVal)
	case int16:
		iVal = int(tVal)
	case int32:
		iVal = int(tVal)
	case int64:
		iVal = int(tVal)
	case uint:
		iVal = int(tVal)
	case uint8:
		iVal = int(tVal)
	case uint16:
		iVal = int(tVal)
	case uint32:
		iVal = int(tVal)
	case uint64:
		iVal = int(tVal)
	case float32:
		iVal = int(tVal)
	case float64:
		iVal = int(tVal)
	case string:
		iVal, err = strconv.Atoi(strings.TrimSpace(tVal))
	default:
		err = ErrConvertFail
	}
	return
}

func Uint(in interface{}) (uint, error) {
	return ToUint(in)
}

func MustUint(in interface{}) uint {
	val, _ := ToUint(in)
	return val
}

func ToUint(in interface{}) (u uint, err error) {
	switch tVal := in.(type) {
	case nil:
		u = 0
	case int:
		u = uint(tVal)
	case int8:
		u = uint(tVal)
	case int16:
		u = uint(tVal)
	case int32:
		u = uint(tVal)
	case int64:
		u = uint(tVal)
	case uint:
		u = tVal
	case uint8:
		u = uint(tVal)
	case uint16:
		u = uint(tVal)
	case uint32:
		u = uint(tVal)
	case uint64:
		u = uint(tVal)
	case float32:
		u = uint(tVal)
	case float64:
		u = uint(tVal)
	case string:
		var u64 uint64
		u64, err = strconv.ParseUint(strings.TrimSpace(tVal), 10, 0)
		u = uint(u64)
	default:
		err = ErrConvertFail
	}
	return
}

func Uint64(in interface{}) (uint64, error) {
	return ToUint64(in)
}

func MustUint64(in interface{}) uint64 {
	val, _ := ToUint64(in)
	return val
}

func ToUint64(in interface{}) (u64 uint64, err error) {
	switch tVal := in.(type) {
	case nil:
		u64 = 0
	case int:
		u64 = uint64(tVal)
	case int8:
		u64 = uint64(tVal)
	case int16:
		u64 = uint64(tVal)
	case int32:
		u64 = uint64(tVal)
	case int64:
		u64 = uint64(tVal)
	case uint:
		u64 = uint64(tVal)
	case uint8:
		u64 = uint64(tVal)
	case uint16:
		u64 = uint64(tVal)
	case uint32:
		u64 = uint64(tVal)
	case uint64:
		u64 = tVal
	case float32:
		u64 = uint64(tVal)
	case float64:
		u64 = uint64(tVal)
	case string:
		u64, err = strconv.ParseUint(strings.TrimSpace(tVal), 10, 0)
	default:
		err = ErrConvertFail
	}
	return
}

func Int64(in interface{}) (int64, error) {
	return ToInt64(in)
}

func MustInt64(in interface{}) int64 {
	i64, _ := ToInt64(in)
	return i64
}

func ToInt64(in interface{}) (i64 int64, err error) {
	switch tVal := in.(type) {
	case nil:
		i64 = 0
	case string:
		i64, err = strconv.ParseInt(strings.TrimSpace(tVal), 10, 0)
	case int:
		i64 = int64(tVal)
	case int8:
		i64 = int64(tVal)
	case int16:
		i64 = int64(tVal)
	case int32:
		i64 = int64(tVal)
	case int64:
		i64 = tVal
	case uint:
		i64 = int64(tVal)
	case uint8:
		i64 = int64(tVal)
	case uint16:
		i64 = int64(tVal)
	case uint32:
		i64 = int64(tVal)
	case uint64:
		i64 = int64(tVal)
	case float32:
		i64 = int64(tVal)
	case float64:
		i64 = int64(tVal)
	default:
		err = ErrConvertFail
	}
	return
}

func Float64(in interface{}) (float64, error) {
	return ToFloat64(in)
}

func ToFloat64(in interface{}) (f64 float64, err error) {
	switch tVal := in.(type) {
	case nil:
		f64 = 0
	case string:
		f64, err = strconv.ParseFloat(strings.TrimSpace(tVal), 0)
	case int:
		f64 = float64(tVal)
	case int8:
		f64 = float64(tVal)
	case int16:
		f64 = float64(tVal)
	case int32:
		f64 = float64(tVal)
	case int64:
		f64 = float64(tVal)
	case uint:
		f64 = float64(tVal)
	case uint8:
		f64 = float64(tVal)
	case uint16:
		f64 = float64(tVal)
	case uint32:
		f64 = float64(tVal)
	case uint64:
		f64 = float64(tVal)
	case float32:
		f64 = float64(tVal)
	case float64:
		f64 = tVal
	default:
		err = ErrConvertFail
	}
	return
}

func MustFloat64(in interface{}) float64 {
	val, _ := ToFloat64(in)
	return val
}
