package paramValidator

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"testing"
)

const testJSON0 = `{
  "field_num": 0.0,
  "field_int": -1,
  "field_float": 0.12,
  "field_string": "its a string",
  "field_object": {
     "subfield_num1": 1,
     "subfield_bool": true
	},
  "field_array": ["a",2,false, {}, []],
  "field_bool": false,
  "field_null": null
}
`

func Test_Required(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	_ = json.Unmarshal([]byte(testJSON0), &jsonValue)
	requiredValidator0 := Validation{
		FieldName: "field_num",
		Required:  true,
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required value which is already present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The optional field <field_num> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "subfield_num1",
		Required:  false,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for an optional value which is not present")
		t.Fail()
	}
	if !result1 {
		t.Errorf("Validation failed. The optional field <subfield_num1> was not present but validation result was false")
		t.Fail()
	}
}

func Test_DataTypeNULL(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_null",
		Required:  true,
		Type:      NULL,
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required null value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The optional field <field_null> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_num",
		Required:  false,
		Type:      NULL,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Num field is encountered when checking for a null value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}
}

func Test_DataTypeNumber(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_num",
		Required:  true,
		Type:      NUMBER,
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required num value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The numeric field <field_num> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_null",
		Required:  false,
		Type:      NUMBER,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Null field is encountered when checking for a num value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}

	requiredValidator2 := Validation{
		FieldName: "field_int",
		Required:  true,
		Type:      NUMBER,
		CustomValidator: func(i interface{}) (bool, error) {

			if reflect.TypeOf(i).Kind() != reflect.Float64 && !(i.(float64) == math.Trunc(i.(float64))) {
				return false, fmt.Errorf("Expecting integer data type found %T", i)
			}
			value := int(math.Trunc(i.(float64)))
			if value < 0 {
				return true, nil
			}
			return false, fmt.Errorf("Expecting a negative value found %d", value)

		},
	}
	result2, err := requiredValidator2.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required num value which is present")
		t.Errorf("Got following error: %s", err.Error())
		t.Fail()
	}
	if !result2 {
		t.Errorf("Validation failed. The numeric field <field_num> was present but validation result was false")
		t.Fail()
	}
}

func Test_DataTypeInteger(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_int",
		Required:  true,
		Type:      INTEGER,
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required int value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The integer field <field_int> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_float",
		Required:  false,
		Type:      INTEGER,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Float field is encountered when checking for a num value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}
}

func Test_DataTypeFloat(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_float",
		Required:  true,
		Type:      FLOAT,
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required float value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The floating point field <field_float> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_string",
		Required:  false,
		Type:      FLOAT,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Float field is encountered when checking for a num value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}
}

func Test_DataTypeString(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_string",
		Required:  true,
		Type:      STRING,
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required string value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The string field <field_string> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_object",
		Required:  false,
		Type:      STRING,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Object field is encountered when checking for a string value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}
}

func Test_DataTypeObject(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_object",
		Required:  true,
		Type:      OBJECT,
		CustomValidator: func(i interface{}) (bool, error) {
			return true, nil
		},
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required object value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The object field <field_object> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_array",
		Required:  false,
		Type:      OBJECT,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Array field is encountered when checking for an object value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}
}

func Test_DataTypeArray(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_array",
		Required:  true,
		Type:      ARRAY,
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required array value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The array field <field_array> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_object",
		Required:  false,
		Type:      ARRAY,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Object field is encountered when checking for an array value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}
}

func Test_DataTypeBoolean(t *testing.T) {
	var jsonValue = make(map[string]interface{})
	json.Unmarshal([]byte(testJSON0), &jsonValue)

	requiredValidator0 := Validation{
		FieldName: "field_bool",
		Required:  true,
		Type:      BOOLEAN,
		CustomValidator: func(i interface{}) (bool, error) {
			return true, nil
		},
	}
	result0, err := requiredValidator0.Apply(jsonValue)
	if err != nil {
		t.Errorf("Error must be nil when checking for a required boolean value which is present")
		t.Fail()
	}
	if !result0 {
		t.Errorf("Validation failed. The boolean field <field_bool> was present but validation result was false")
		t.Fail()
	}

	requiredValidator1 := Validation{
		FieldName: "field_num",
		Required:  false,
		Type:      BOOLEAN,
	}
	result1, err := requiredValidator1.Apply(jsonValue)
	if err == nil {
		t.Errorf("Error must not be nil. Numeric field is encountered when checking for an boolean value")
		t.Fail()
	}
	if result1 {
		t.Errorf("Expected: Validation failing. Actual: Validation passed.")
	}
}
