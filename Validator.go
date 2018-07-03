package param_validator

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

// Apply function executes the validator against the parsed JSON map
func (*Validation) Apply(jsonValue map[string]interface{}) (bool, error) {
	panic("Yet to be implemented")
}