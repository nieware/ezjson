# ezjson
A package to easily read arbitrary JSON data.

[![Go Report Card](https://goreportcard.com/badge/github.com/nieware/ezjson)](https://goreportcard.com/report/github.com/nieware/ezjson)
[![GoDoc](https://godoc.org/github.com/nieware/ezjson?status.svg)](https://godoc.org/github.com/nieware/ezjson)

## Usage

To decode JSON from a string or a slice of bytes, use the functions `DecodeString` or `DecodeBytes`, which return an `interface{}` containing the data.

To get a property from the JSON data, use the function `GetProperty`. For example, if you have the following JSON:

    {
        "resp":{
            "s":"a string",
            "i":42,
            "f":12.34,
            "a":[
                1,
                2,
                false
            ],
            "n":null
        }
    }

...you can get the property `resp` as an `interface{}` with `resp, err := ezjson.GetProperty(data, "resp")`. You can then get the string `s` with `s, err := ezjson.GetString(resp, "s")` or directly from `data` with `s, err := ezjson.GetString(data, "resp", "s")`. You can get the first element of the array a with `i, err := ezjson.GetInt(data, "resp", "a", 0)` or the third one with `b, err := ezjson.GetBool(resp, "a", 2)`.

Usually, `null` values are returned as a `nil` interface or as the zero type of the other functions - i.e. `n, err := ezjson.GetInt(data, "resp", "n") ` will return 0. You can specify the option `ezjson.ErrorOnNull` to get an `ezjson.NullError` in case of a null value: when calling `n, err := ezjson.GetInt(data, ezjson.ErrorOnNull, "resp", "n")`, err will be `ezjson.NullError`.

See the [documentation and examples](https://godoc.org/github.com/nieware/ezjson) for more details.