package GoMybatis

import (
	"reflect"
	"time"
	"strconv"
	"bytes"
)

const Adapter_DateType  =  `time.Time`
const Adapter_StringType  =  `string`
const Adapter_FormateDate  =  `2006-01-02 15:04:05`

var DefaultExpressionTypeConvertFunc = func(arg interface{}) interface{} {
	if reflect.TypeOf(arg).String() == Adapter_DateType {
		return arg.(time.Time).Nanosecond()
	}
	return arg
}

var DefaultSqlTypeConvertFunc = func(arg interface{}) string {
	var t = reflect.TypeOf(arg)
	if t.String() == Adapter_DateType {
		arg = arg.(time.Time).Format(Adapter_FormateDate)
	}
	if t.String() == Adapter_DateType || t.String() == Adapter_StringType {
		var argStr bytes.Buffer
		argStr.WriteString(`'`)
		argStr.WriteString(toString(arg))
		argStr.WriteString(`'`)
		return argStr.String()
	}
	return toString(arg)
}

func toString(value interface{}) string {
	if value == nil {
		return ""
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Int {
		string := strconv.Itoa(value.(int))
		return string
	} else if v.Kind() == reflect.Int32 {
		string := strconv.FormatInt(int64(value.(int32)), 10)
		return string
	} else if v.Kind() == reflect.Int64 {
		string := strconv.FormatInt(value.(int64), 10)
		return string
	} else if v.Kind() == reflect.Float32 {
		string := strconv.FormatFloat(float64(value.(float32)), 'f', 6, 64)
		return string
	} else if v.Kind() == reflect.Float64 {
		string := strconv.FormatFloat(value.(float64), 'f', 6, 64)
		return string
	} else if v.Kind() == reflect.String {
		return value.(string)
	} else {
		return ""
	}
}
