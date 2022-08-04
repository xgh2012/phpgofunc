// Package utilsjson
// json处理
package utilsjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

// HandleAnyToString 处理任意数据
func HandleAnyToString(data interface{}) (string, error) {
	switch data.(type) {
	//region 本方法自己处理 ---- 开始
	case string:
		return data.(string), nil
	case int:
		return cast.ToString(data), nil
	case int8:
		return cast.ToString(data), nil
	case int16:
		return cast.ToString(data), nil
	case int32:
		return cast.ToString(data), nil
	case int64:
		return cast.ToString(data), nil
	case float32:
		return cast.ToString(data), nil
	case float64:
		return cast.ToString(data), nil
	case bool:
		var val = "0"
		if data.(bool) {
			val = "1"
		}
		return val, nil
	case *int:
		return cast.ToString(data), nil
	case *int32:
		return cast.ToString(data), nil
	case *int64:
		return cast.ToString(data), nil
	case *float32:
		return cast.ToString(data), nil
	case *float64:
		return cast.ToString(data), nil
	case *bool:
		var val = "0"
		if *data.(*bool) {
			val = "1"
		}
		return val, nil
	case []string:
		tmpIntS := `[`
		for _, item := range data.([]string) {
			tmpIntS += `"` + item + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []int:
		tmpIntS := `[`
		for _, item := range data.([]int) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []int8:
		tmpIntS := `[`
		for _, item := range data.([]int8) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []int16:
		tmpIntS := `[`
		for _, item := range data.([]int16) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []int32:
		tmpIntS := `[`
		for _, item := range data.([]int32) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []int64:
		tmpIntS := `[`
		for _, item := range data.([]int64) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []float32:
		tmpIntS := `[`
		for _, item := range data.([]float32) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []float64:
		tmpIntS := `[`
		for _, item := range data.([]float64) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []bool:
		tmpIntS := `[`
		for _, item := range data.([]bool) {
			var val = "0"
			if item == true {
				val = "1"
			}
			tmpIntS += `"` + val + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*string:
		tmpIntS := `[`
		for _, item := range data.([]*string) {
			tmpIntS += `"` + *item + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*int:
		tmpIntS := `[`
		for _, item := range data.([]*int) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*int8:
		tmpIntS := `[`
		for _, item := range data.([]*int8) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*int16:
		tmpIntS := `[`
		for _, item := range data.([]*int16) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*int32:
		tmpIntS := `[`
		for _, item := range data.([]*int32) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*int64:
		tmpIntS := `[`
		for _, item := range data.([]*int64) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*float32:
		tmpIntS := `[`
		for _, item := range data.([]*float32) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*float64:
		tmpIntS := `[`
		for _, item := range data.([]*float64) {
			tmpIntS += `"` + cast.ToString(item) + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	case []*bool:
		tmpIntS := `[`
		for _, item := range data.([]*bool) {
			var val = "0"
			if *interface{}(item).(*bool) == true {
				val = "1"
			}
			tmpIntS += `"` + val + `",`
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	//region 本方法自己处理 ---- 结束

	//kind := reflect.ValueOf(data).Kind()
	//if kind == reflect.Map || kind == reflect.Struct || kind == reflect.Ptr {
	case map[string]interface{}:
		return HandleMapAnyToString(data.(map[string]interface{}))
	case *map[string]interface{}:
		return HandleMapAnyToString(*data.(*map[string]interface{}))
	}
	kind := reflect.ValueOf(data).Kind()

	//fmt.Printf("keyMap %+v kind %+v \n",key,kind)
	if kind == reflect.Map || kind == reflect.Struct || kind == reflect.Ptr {
		var tp = make(map[string]interface{})
		tpMap, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("keyMap1 %+v Map %+v \n", data, err)
			return "", err
		}
		err = json.Unmarshal(tpMap, &tp)
		if err != nil {
			fmt.Printf("keyMap2 %+v Map %+v \n", data, err)
			return "", err
		}
		return HandleMapAnyToString(tp)
	} else if kind == reflect.Slice {
		var tp []interface{}
		tpMap, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("keyMap3 %+v Map %+v \n", data, err)
			return "", err
		}
		err = json.Unmarshal(tpMap, &tp)
		if err != nil {
			fmt.Printf("keyMap4 %+v Map %+v \n", data, err)
			return "", err
		}
		tmpIntS := "["
		for _, i2 := range tp {
			contents, err := HandleMapAnyToString(i2.(map[string]interface{}))
			if err != nil {
				return "", err
			}
			tmpIntS += contents
		}
		return strings.Trim(tmpIntS, ",") + "]", nil
	}

	return "", errors.New("不支持的数据类型")
}

// HandleMapAnyToString 将任意数据转换成 string
func HandleMapAnyToString(data map[string]interface{}) (string, error) {
	var (
		buff = "{"
	)

	for key, value := range data {
		switch value.(type) {
		case string:
			//判断是否是json
			var tmp map[string]interface{}
			val := value.(string)
			err := json.Unmarshal([]byte(val), &tmp)
			if err != nil { //是字符串 或者 "[]"
				//[] 特殊处理
				if strings.ReplaceAll(val, " ", "") == "[]" {
					buff += `"` + key + `":"[]",`
					continue
				}
				buff += `"` + key + `":"` + val + `",`
				continue
			}

			//{} 特殊处理
			if len(tmp) == 0 {
				buff += `"` + key + `":"{}",`
				continue
			}
			contents, err := HandleMapAnyToString(tmp)
			if err != nil {
				return "", err
			}
			buff += `"` + key + `":` + contents
		case int:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case int8:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case int16:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case int32:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case int64:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case float32:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case float64:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case bool:
			val := "0"
			if value.(bool) {
				val = "1"
			}
			buff += `"` + key + `":"` + val + `",`
		case *int:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case *int8:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case *int16:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case *int32:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case *int64:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case *float32:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case *float64:
			buff += `"` + key + `":"` + cast.ToString(value) + `",`
		case *bool:
			val := "0"
			if *value.(*bool) {
				val = "1"
			}
			buff += `"` + key + `":"` + val + `",`
		case []string:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]string) {
				tmpIntS += `"` + item + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []int:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]int) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []int8:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]int8) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []int16:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]int16) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []int32:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]int32) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []int64:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]int64) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []float32:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]float32) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []float64:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]float64) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []bool:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]bool) {
				var val = "0"
				if item == true {
					val = "1"
				}
				tmpIntS += `"` + val + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*string:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*string) {
				tmpIntS += `"` + *item + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*int:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*int) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*int8:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*int8) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*int16:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*int16) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*int32:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*int32) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*int64:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*int64) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*float32:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*float32) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*float64:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*float64) {
				tmpIntS += `"` + cast.ToString(item) + `",`
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case []*bool:
			tmpIntS := `"` + key + `":[`
			for _, item := range value.([]*bool) {
				var val = "0"
				if *interface{}(item).(*bool) == true {
					val = "1"
				}
				tmpIntS += `"` + val + `",`
				//fmt.Printf("item %+v %+v \n",reflect.TypeOf(item),*interface{}(item).(*bool))
			}
			buff += strings.Trim(tmpIntS, ",") + "],"
		case map[string]interface{}:
			contents, err := HandleMapAnyToString(value.(map[string]interface{}))
			if err != nil {
				return "", err
			}
			buff += `"` + key + `":` + contents
		case *map[string]interface{}:
			contents, err := HandleMapAnyToString(*value.(*map[string]interface{}))
			if err != nil {
				return "", err
			}
			buff += `"` + key + `":` + contents
		default:
			kind := reflect.ValueOf(value).Kind()
			//fmt.Printf("keyMap %+v kind %+v \n",key,kind)
			if kind == reflect.Map || kind == reflect.Struct || kind == reflect.Ptr {
				var tp = make(map[string]interface{})
				tpMap, err := json.Marshal(value)
				if err != nil {
					fmt.Printf("keyMap %+v Map %+v \n", key, err)
				}
				err = json.Unmarshal(tpMap, &tp)
				if err != nil {
					fmt.Printf("keyMap %+v Map %+v \n", key, err)
				}
				contents, err := HandleMapAnyToString(tp)
				if err != nil {
					return "", err
				}
				buff += `"` + key + `":` + contents
			} else if kind == reflect.Slice {
				var tp []interface{}
				tpMap, err := json.Marshal(value)
				if err != nil {
					fmt.Println(key, err, reflect.TypeOf(value))
				}
				err = json.Unmarshal(tpMap, &tp)
				if err != nil {
					fmt.Println(key, err, reflect.TypeOf(value))
				}
				tmpIntS := `"` + key + `":[`
				for _, i2 := range tp {
					contents, err := HandleMapAnyToString(i2.(map[string]interface{}))
					if err != nil {
						return "", err
					}
					tmpIntS += contents
				}
				buff += strings.Trim(tmpIntS, ",") + "],"
			}
		}
	}
	return strings.Trim(buff, ",") + "},", nil
}
