package roundtrip

import "math/bits"

tests: primitive: {
	inSchema: {
		type: "record"
		name: "R"
		fields: [{
			name: "intField"
			type: "int"
		}, {
			name: "longField"
			type: "long"
		}, {
			name: "floatField"
			type: "float"
		}, {
			name: "doubleField"
			type: "double"
		}, {
			name: "boolField"
			type: "boolean"
		}, {
			name: "bytesField"
			type: "bytes"
		}, {
			name: "stringField"
			type: "string"
		}]
	}
	outSchema: inSchema
}

tests: primitive: subtests: highValues: {
	inData: {
		intField:    bits.Lsh(1, 31) - 1
		longField:   bits.Lsh(1, 63) - 1
		floatField:  2e-10
		doubleField: 2e-50
		boolField:   true
		// We'd include some binary data in the bytes field except for
		// https://github.com/linkedin/goavro/issues/192
		bytesField:  "stuff"
		stringField: "hello world"
	}
	outData: inData
}

tests: primitive: subtests: lowValues: {
	inData: {
		intField:    -bits.Lsh(1, 31)
		longField:   -bits.Lsh(1, 63)
		floatField:  -2e-10
		doubleField: -2e-50
		boolField:   false
		// We'd include some binary data in the bytes field except for
		// https://github.com/linkedin/goavro/issues/192
		bytesField:  ""
		stringField: ""
	}
	outData: inData
}
