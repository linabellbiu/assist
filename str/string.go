package str

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"strconv"
	"strings"
)

//拼接字符串
//返回string类型
func Join(arg ...interface{}) (string, error) {
	if len(arg) == 0 {
		return "", nil
	}
	if len(arg) == 1 {
		str, err := toString(arg[0])
		if err != nil {
			e := fmt.Sprintf("arg %d:%s", 1, err.Error())
			return "", errors.New(e)
		}
		return str, nil
	}
	var s strings.Builder
	for index, str := range arg {
		str, err := toString(str)
		if err != nil {
			e := fmt.Sprintf("arg %d:%s", index+1, err.Error())
			return "", errors.New(e)
		}
		s.WriteString(str)
	}
	_s := s
	return _s.String(), nil
}

func toString(val interface{}) (string, error) {
	var s string
	switch val.(type) {
	case int:
		s = strconv.Itoa(val.(int))
	case int8:
		s = strconv.FormatInt(int64(val.(int8)), 10)
	case int16:
		s = strconv.FormatInt(int64(val.(int16)), 10)
	case int32:
		s = strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		s = strconv.FormatInt(val.(int64), 10)
	case float32:
		s = decimal.NewFromFloat32(val.(float32)).String()
	case float64:
		s = decimal.NewFromFloat(val.(float64)).String()
	case uint:
		s = strconv.Itoa(int(val.(uint)))
	case uint8:
		s = strconv.Itoa(int(val.(uint8)))
	case uint16:
		s = strconv.Itoa(int(val.(uint16)))
	case uint32:
		s = strconv.Itoa(int(val.(uint32)))
	case uint64:
		s = strconv.Itoa(int(val.(uint64)))
	case string:
		s = val.(string)
	case []byte:
		s = string(val.([]byte))
	default:
		s = fmt.Sprintf("%s to string:Unsupported type conversion", reflect.TypeOf(val))
		return "", errors.New(s)
	}

	return s, nil
}
