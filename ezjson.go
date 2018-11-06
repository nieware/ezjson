package ezjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

/*
DecodeBytes decodes JSON data from a byte array (using json.Number)
*/
func DecodeBytes(cont []byte) (data interface{}, err error) {
	d := json.NewDecoder(bytes.NewReader(cont))
	d.UseNumber()
	err = d.Decode(&data)
	return
}

/*
DecodeString decodes JSON data from a string (using json.Number)
*/
func DecodeString(cont string) (data interface{}, err error) {
	data, err = DecodeBytes([]byte(cont))
	return
}

/*
GetProperty returns a property from "JSON" object intf using the given (nested) keys.
Each key can be either a string (for an object property) or an integer (for an array index).
*/
func GetProperty(intf interface{}, keys ...interface{}) (res interface{}, skey string, err error) {
	for idx, key := range keys {
		switch k := key.(type) {
		case int:
			skey = strconv.Itoa(k)
			a, isArray := intf.([]interface{})
			if !isArray {
				err = fmt.Errorf("No array found for key #%d (%d)", idx, k)
				return
			}
			intf = a[k]
		case string:
			skey = k
			m, isMap := intf.(map[string]interface{})
			if !isMap {
				err = fmt.Errorf("No object found for key #%d (\"%s\")", idx, k)
				return
			}
			var ok bool
			intf, ok = m[k]
			if !ok {
				err = fmt.Errorf("Object property not found for key #%d (\"%s\")", idx, k)
				return
			}
		default:
			return
		}
	}

	return intf, skey, nil
}

/*
GetArray returns an array contained in the "JSON" object intf using the given (nested) keys.
Each key can be either a string (for an object property) or an integer (for an array index).
*/
func GetArray(intf interface{}, keys ...interface{}) (res []interface{}, err error) {
	value, skey, err := GetProperty(intf, keys...)
	if err != nil {
		return
	}

	// tries to cast the value to a slice of interfaces
	res, ok := value.([]interface{})
	if !ok {
		err = fmt.Errorf("Property \"%s\" is not of type array", skey)
	}
	return
}

/*
GetNumber returns a Number contained in the "JSON" object intf using the given (nested) keys.
Each key can be either a string (for an object property) or an integer (for an array index).
*/
func GetNumber(intf interface{}, keys ...interface{}) (res json.Number, err error) {
	value, skey, err := GetProperty(intf, keys...)
	if err != nil {
		return
	}

	// get the value as a "Number"
	res, ok := value.(json.Number)
	if !ok {
		err = fmt.Errorf("Property \"%s\" is not of type Number", skey)
	}
	return
}

/*
GetInt returns an int64 contained in the "JSON" object intf using the given (nested) keys.
Each key can be either a string (for an object property) or an integer (for an array index).
*/
func GetInt(intf interface{}, keys ...interface{}) (res int64, err error) {
	num, err := GetNumber(intf, keys...)
	if err != nil {
		return
	}

	// return as Int64
	res, err = num.Int64()
	return
}

/*
GetFloat returns a float64 contained in the "JSON" object intf using the given (nested) keys.
Each key can be either a string (for an object property) or an integer (for an array index).
*/
func GetFloat(intf interface{}, keys ...interface{}) (res float64, err error) {
	num, err := GetNumber(intf, keys...)
	if err != nil {
		return
	}

	// return as Float64
	res, err = num.Float64()
	return
}

/*
GetString returns a string contained in the "JSON" object intf using the given (nested) keys.
Each key can be either a string (for an object property) or an integer (for an array index).
*/
func GetString(intf interface{}, keys ...interface{}) (res string, err error) {
	value, skey, err := GetProperty(intf, keys...)
	if err != nil {
		return
	}

	// tries to cast the value to a string
	res, ok := value.(string)
	if !ok {
		err = fmt.Errorf("Property \"%s\" is not of type string", skey)
	}
	return
}

/*
GetBool returns a bool contained in the "JSON" object intf using the given (nested) keys.
Each key can be either a string (for an object property) or an integer (for an array index).
*/
func GetBool(intf interface{}, keys ...interface{}) (res bool, err error) {
	value, skey, err := GetProperty(intf, keys...)
	if err != nil {
		return
	}

	// tries to cast the value to a boolean
	res, ok := value.(bool)
	if !ok {
		err = fmt.Errorf("Property \"%s\" is not of type bool", skey)
	}
	return
}
