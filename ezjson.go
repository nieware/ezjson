// Allows you to easily read arbitrary JSON data in non-performance-critical use cases

package ezjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Option allows changing the behaviour of the GetProperty functions
type Option int

const (
	// ErrorOnNull - if this option is specified, a NullError will be returned in case of a null value.
	ErrorOnNull Option = 1
)

// DecodeBytes decodes JSON data from a byte array (using json.Number)
func DecodeBytes(cont []byte) (data interface{}, err error) {
	d := json.NewDecoder(bytes.NewReader(cont))
	d.UseNumber()
	err = d.Decode(&data)
	return
}

// DecodeString decodes JSON data from a string (using json.Number)
func DecodeString(cont string) (data interface{}, err error) {
	return DecodeBytes([]byte(cont))
}

// KeyError is a custom error type which is returned when there is a problem with a particular
// key in a sequence of keys for nested data structures.
type KeyError struct {
	Msg string // the error message
	Idx int    // the index of the key in the sequence (the nesting level)
	Key string // the key
}

func (e *KeyError) Error() string {
	return fmt.Sprintf("%s for key %s (@ index %d)", e.Msg, e.Key, e.Idx)
}

// NullError is a custom error type which is returned if a value is "null" and the option ErrorOnNull is passed
type NullError struct {
	Key string
}

func (e *NullError) Error() string {
	return fmt.Sprintf("Value is null for key %s", e.Key)
}

// GetPropertyWithType returns a property from "JSON" object intf using the given (nested) keys.
// Each key can be either a string (for an object property), an integer (for an array index) or an option
// which changes the behaviour of the function, e.g. ErrorOnNull. Options must be specified before the actual keys.
// The function optionally checks if the property is of type dType and returns an error if it's not.
// Supported types are "array", "number", "string" and "bool". If another value is given, no checks are made.
// If ErrorOnNull is passed, a NullError is returned if the value is null.
func GetPropertyWithType(intf interface{}, dType string, keys ...interface{}) (res interface{}, err error) {
	skey, idx, ok, returnErrorOnNull, expectOptions := "", 0, true, false, true
	for idx, key := range keys {
		switch k := key.(type) {
		case int:
			skey = strconv.Itoa(k)
			expectOptions = false
			a, isArray := intf.([]interface{})
			if !isArray {
				return nil, &KeyError{"No array found", idx, skey}
			}
			if k < 0 || k >= len(a) {
				return nil, &KeyError{"Array index out of bounds", idx, skey}
			}
			intf = a[k]
		case string:
			skey = k
			expectOptions = false
			m, isMap := intf.(map[string]interface{})
			if !isMap {
				return nil, &KeyError{"No object found", idx, skey}
			}
			var ok bool
			intf, ok = m[k]
			if !ok {
				return nil, &KeyError{"Object property not found", idx, skey}
			}
		case Option:
			if !expectOptions {
				return nil, &KeyError{"Options must be specified before the actual keys", idx, fmt.Sprint("#v", k)}
			}
			if k == ErrorOnNull {
				returnErrorOnNull = true
			}
		default:
			return nil, &KeyError{"Not int or string", idx, fmt.Sprint("#v", k)}
		}
	}

	switch dType {
	case "array":
		_, ok = intf.([]interface{})
	case "number":
		_, ok = intf.(json.Number)
	case "string":
		_, ok = intf.(string)
	case "bool":
		_, ok = intf.(bool)
	}

	// if the type cast fails, we return an error - except if the value is "null", because
	// any value can be "null" in JSON. In this case, intf will be nil - checking for this is left to the caller...
	if !ok && intf != nil {
		return nil, &KeyError{fmt.Sprintf("Property is not of type %s", dType), idx, skey}
	}
	// ...except when the ErrorOnNull option is specified, in which case we return a NullError.
	if intf == nil && returnErrorOnNull {
		return nil, &NullError{skey}
	}

	return intf, nil
}

// GetProperty returns a property contained in the "JSON" object intf using the given (nested) keys as an interface{}.
// Each key can be a string (for an object property) an integer (for an array index) or an option.
func GetProperty(intf interface{}, keys ...interface{}) (res interface{}, err error) {
	return GetPropertyWithType(intf, "", keys...)
}

// GetArray returns an array contained in the "JSON" object intf using the given (nested) keys.
// Each key can be a string (for an object property) an integer (for an array index) or an option.
func GetArray(intf interface{}, keys ...interface{}) (res []interface{}, err error) {
	value, err := GetPropertyWithType(intf, "array", keys...)
	if err != nil {
		return
	}
	// cast the value to an array (we know it will work because GetProperty already tried it)
	res, _ = value.([]interface{})
	return
}

// GetNumber returns a Number contained in the "JSON" object intf using the given (nested) keys.
// Each key can be a string (for an object property) an integer (for an array index) or an option.
func GetNumber(intf interface{}, keys ...interface{}) (res json.Number, err error) {
	value, err := GetPropertyWithType(intf, "number", keys...)
	if err != nil {
		return
	}
	// get the value as a "Number" (we know it will work because GetProperty already tried it)
	res, _ = value.(json.Number)
	return
}

// GetInt returns an int64 contained in the "JSON" object intf using the given (nested) keys.
// Each key can be a string (for an object property) an integer (for an array index) or an option.
func GetInt(intf interface{}, keys ...interface{}) (res int64, err error) {
	num, err := GetNumber(intf, keys...)
	if err != nil {
		return
	}
	// return as Int64
	return num.Int64()
}

// GetFloat returns a float64 contained in the "JSON" object intf using the given (nested) keys.
// Each key can be a string (for an object property) an integer (for an array index) or an option.
func GetFloat(intf interface{}, keys ...interface{}) (res float64, err error) {
	num, err := GetNumber(intf, keys...)
	if err != nil {
		return
	}
	// return as Float64
	return num.Float64()
}

// GetString returns a string contained in the "JSON" object intf using the given (nested) keys.
// Each key can be a string (for an object property) an integer (for an array index) or an option.
func GetString(intf interface{}, keys ...interface{}) (res string, err error) {
	value, err := GetPropertyWithType(intf, "string", keys...)
	if err != nil {
		return
	}
	// cast the value to a string (we know it will work because GetProperty already tried it)
	res, _ = value.(string)
	return
}

// GetBool returns a bool contained in the "JSON" object intf using the given (nested) keys.
// Each key can be a string (for an object property) an integer (for an array index) or an option.
func GetBool(intf interface{}, keys ...interface{}) (res bool, err error) {
	value, err := GetPropertyWithType(intf, "bool", keys...)
	if err != nil {
		return
	}
	// cast the value to a string (we know it will work because GetProperty already tried it)
	res, _ = value.(bool)
	return
}
