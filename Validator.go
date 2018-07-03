package paramValidator

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// DataType - Enumeration type for JSON types
type DataType int

// Validator is any function that accepts a given JSON value and outputs validation status and error if validation fails
type Validator func(interface{}) (bool, error)

// NULL corresponds to golang nil
// INTEGER would be parsed as golang int
const (
	ANY DataType = iota
	NULL
	INTEGER
	FLOAT
	NUMBER
	STRING
	BOOLEAN
	OBJECT
	ARRAY
)

// Validation struct holds the validations to be run against the given field name
type Validation struct {
	FieldName       string
	Required        bool
	Type            DataType
	CustomValidator Validator
}

func (val *Validation) Apply(jsonValue map[string]interface{}) (bool, error) {

	if val.FieldName == "" {
		return false, errors.New("Empty Field Name")
	} else {
		value, present := jsonValue[val.FieldName]
		if !present {
			if val.Required {
				return false, errors.New("Required field " + val.FieldName + " is missing")
			} else {
				return true, nil
			}
		} else {
			if val.Type == NULL {
				if value == nil {
					return true, nil
				} else {
					return false, errors.New("Null value in " + val.FieldName)
				}
			} else {
				if val.Type == ANY {
					return checkValidation(value, val)
				} else {

					if value == nil {
						return false, errors.New("Null value in " + val.FieldName)
					}
					typeResult, err := checkType(value, val)
					if typeResult {
						return checkValidation(value, val)
					} else {
						return false, err
					}
				}
			}
		}
	}
}

func checkType(v interface{}, val *Validation) (bool, error) {
	var typeString = ""
	switch val.Type {

	case INTEGER:
		typeString = "an integer"
		if reflect.ValueOf(v).Kind() == reflect.Float64 || reflect.ValueOf(v).Kind() == reflect.Float32 {
			if v.(float64) == math.Trunc(v.(float64)) {
				return true, nil
			}
		}
	case STRING:
		typeString = "string"
		if reflect.ValueOf(v).Kind() == reflect.String {
			return true, nil

		}
	case FLOAT:
		typeString = "float"
		if reflect.ValueOf(v).Kind() == reflect.Float64 || reflect.ValueOf(v).Kind() == reflect.Float32 {
			return true, nil

		}
	case BOOLEAN:
		typeString = "bool"
		if reflect.ValueOf(v).Kind() == reflect.Bool {
			return true, nil

		}
	case ARRAY:
		typeString = "an array"

		if reflect.ValueOf(v).Kind() == reflect.Slice {
			return true, nil
		}
	case OBJECT:
		typeString = "an object"
		if reflect.ValueOf(v).Kind() == reflect.Map {
			return true, nil
		}
	case NUMBER:
		typeString = "a number"
		if reflect.ValueOf(v).Kind() == reflect.Float64 || reflect.ValueOf(v).Kind() == reflect.Float32 || reflect.ValueOf(v).Kind() == reflect.Int {
			return true, nil
		}

	}
	//print("Expected "+ typeString +" but found an ")
	//fmt.Printf("%v", reflect.TypeOf(v).Kind())
	//println("  ")
	//print("===================")
	return false, errors.New("Expected " + typeString + " but found an " + fmt.Sprintf("%v", reflect.TypeOf(v).Kind()))
}

func checkValidation(v interface{}, val *Validation) (bool, error) {
	if val.CustomValidator == nil {
		return true, nil
	} else {
		return val.CustomValidator(v)
	}
}
