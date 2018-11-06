package ezjson_test

import (
	"fmt"
	"testing"

	"github.com/nieware/ezjson"
)

var testDataString = `
{
	"data":{
		"subData":{
			"array":[
				{
					"str":"a string",
					"int":42
				},
				"string in array",
				12.34,
				true
			],
			"bool":false
		},
		"int":123,
		"str":"string in data"
	},
	"moreData":{
		"str":"string in moreData"
	},
	"array":[
		1,
		2,
		3
	]
}`

/*
Decodes the JSON data from testData and prints it.
*/
func ExampleDecodeString() {
	testData, err := ezjson.DecodeString(testDataString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", testData)
	// this can't be used as a test function (yet?) because the sort order when printing maps is undefined
}

/*
Reads and prints the property data.subData from testData
*/
func ExampleGetProperty() {
	testData, _ := ezjson.DecodeString(testDataString)
	res, skey, err := ezjson.GetProperty(testData, "data", "subData")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("'%s': %#v\n", skey, res)
	// this can't be used as a test function (yet?) because the sort order when printing maps is undefined
}

/*
Reads the array property array from testData
*/
func ExampleGetArray() {
	testData, _ := ezjson.DecodeString(testDataString)
	res, err := ezjson.GetArray(testData, "array")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	// Output: [1 2 3]
}

/*
Reads bool property data.subData.bool from testData
*/
func ExampleGetBool() {
	testData, _ := ezjson.DecodeString(testDataString)
	res, err := ezjson.GetBool(testData, "data", "subData", "bool")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	// Output: false
}

/*
Reads the deeply nested string property data.subData.array[0].str from testData
*/
func ExampleGetString() {
	testData, _ := ezjson.DecodeString(testDataString)
	res, err := ezjson.GetString(testData, "data", "subData", "array", 0, "str")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	// Output: a string
}

/*
Reads property data.subData.array[0].int from testData as a json.Number
*/
func ExampleGetNumber() {
	testData, _ := ezjson.DecodeString(testDataString)
	res, err := ezjson.GetNumber(testData, "data", "int")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	// Output: 123
}

/*
Reads int property array[1] from testData
*/
func ExampleGetInt() {
	testData, _ := ezjson.DecodeString(testDataString)
	res, err := ezjson.GetInt(testData, "array", 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	// Output: 2
}

/*
Reads float property data.subData.array[2] from testData
*/
func ExampleGetFloat() {
	testData, _ := ezjson.DecodeString(testDataString)
	res, err := ezjson.GetFloat(testData, "data", "subData", "array", 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	// Output: 12.34
}

/*
TestGetIntWrongType checks trying to read a string property as an int
*/
func TestGetIntWrongType(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)
	_, err := ezjson.GetInt(testData, "data", "subData", "array", 0, "str")
	if err == nil {
		t.FailNow()
	}
}

/*
TestGetFloatWrongType checks trying to read a string property as an float
*/
func TestGetFloatWrongType(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)
	_, err := ezjson.GetFloat(testData, "data", "subData", "array", 0, "str")
	if err == nil {
		t.FailNow()
	}
}

/*
TestGetNumberWrongType checks trying to read a string property as a json.Number
*/
func TestGetNumberWrongType(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)
	_, err := ezjson.GetNumber(testData, "data", "subData", "array", 0, "str")
	if err == nil {
		t.FailNow()
	}
}

/*
TestGetBoolWrongType checks trying to read a string property as a json.Number
*/
func TestGetStringWrongType(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)
	_, err := ezjson.GetString(testData, "data", "subData", "array", 0, "int")
	if err == nil {
		t.FailNow()
	}
}

/*
TestGetBoolWrongType checks trying to read a string property as a json.Number
*/
func TestGetBoolWrongType(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)
	_, err := ezjson.GetBool(testData, "data", "subData", "array", 0, "str")
	if err == nil {
		t.FailNow()
	}
}

/*
TestGetArrayWrongType checks trying to read a string property as a json.Number
*/
func TestGetArrayWrongType(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)
	_, err := ezjson.GetArray(testData, "data", "subData", "array", 0, "str")
	if err == nil {
		t.FailNow()
	}
}

/*
TestPropertyNotFound tests if the various functions return an error for non-existent properties
*/
func TestPropertyNotFound(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)

	_, err := ezjson.GetArray(testData, "data", "inexistentArray")
	if err == nil {
		t.Fatal("Property not found array")
	}

	_, err = ezjson.GetBool(testData, "data", "inexistentBool")
	if err == nil {
		t.Fatal("Property not found bool")
	}

	_, err = ezjson.GetString(testData, "data", "inexistentString")
	if err == nil {
		t.Fatal("Property not found string")
	}

	_, err = ezjson.GetNumber(testData, "data", "inexistentNumber")
	if err == nil {
		t.Fatal("Property not found Number")
	}

	_, err = ezjson.GetInt(testData, "data", "inexistentInt")
	if err == nil {
		t.Fatal("Property not found int")
	}

	_, err = ezjson.GetFloat(testData, "data", "inexistentFloat")
	if err == nil {
		t.Fatal("Property not found float")
	}

	return
}

/*
TestIncorrectPath tests incorrect property paths (e.g. containing int where a string is expected)
*/
func TestIncorrectPath(t *testing.T) {
	testData, _ := ezjson.DecodeString(testDataString)

	_, err := ezjson.GetArray(testData, "array", "inexistentIndex")
	if err == nil {
		t.Fatal("Failed string instead of int key")
	}

	_, err = ezjson.GetBool(testData, "data", 0)
	if err == nil {
		t.Fatal("Failed int instead of string key")
	}

	_, err = ezjson.GetBool(testData, "data", true)
	if err == nil {
		t.Fatal("Failed key of wrong type")
	}

	return
}
