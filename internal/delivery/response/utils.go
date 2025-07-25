package response

import "reflect"

func toSafeData(data interface{}) interface{} {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice && val.IsNil() {
		return []interface{}{}
	}
	return data
}
