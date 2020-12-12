package utils

import (
	"reflect"
	"testing"
)

func TDeepEqual(t *testing.T, got interface{}, expected interface{}) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got = %v, want %v", got, expected)
	}

	valueOfGot := reflect.ValueOf(got)
	valueOfExpect := reflect.ValueOf(expected)
	if valueOfGot.IsNil() || valueOfExpect.IsNil() {
		return
	}

	gotVal := valueOfGot.Elem()
	expectedVal := valueOfExpect.Elem()
	for i := 0; i < gotVal.NumField(); i++ {
		valueField := gotVal.Field(i)
		typeField := gotVal.Type().Field(i)
		fieldName := typeField.Name

		if fieldName[0] < 'A' || fieldName[0] > 'Z' {
			continue
		}

		gotValue := valueField.Interface()
		expectedValue := expectedVal.FieldByName(fieldName).Interface()
		if !reflect.DeepEqual(gotValue, expectedValue) {
			t.Errorf("%v: (%v/%v) expected: (%v/%v)", fieldName, gotValue, reflect.TypeOf(gotValue), expectedValue, reflect.TypeOf(expectedValue))
		}
	}
}
