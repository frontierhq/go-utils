package slice

import (
	"fmt"
	"reflect"
)

func ConvertAll[K interface{}, T interface{}](source []T) ([]K, error) {
	converted := []K{}
	for _, i := range source {
		intf, ok := any(i).(K)
		if !ok {
			var zeroValue K
			return nil, fmt.Errorf("failed to convert source item (sourceType=%+v, targetType=%+v)", reflect.TypeOf(i), reflect.TypeOf(zeroValue))
		}
		converted = append(converted, intf)
	}
	return converted, nil
}
