// Code generated by generatetestcode.go; DO NOT EDIT.

package unionInOut

import (
	"testing"

	"github.com/heetch/avro/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "PrimitiveUnionTestRecord",
                "type": "record",
                "fields": [
                    {
                        "name": "UnionField",
                        "type": [
                            "int",
                            "long",
                            "float",
                            "double",
                            "string",
                            "boolean",
                            "null"
                        ],
                        "default": 1234
                    }
                ]
            }`,
	GoType: new(PrimitiveUnionTestRecord),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "withBoolean",
		InDataJSON: `{
                        "UnionField": {
                            "boolean": true
                        }
                    }`,
		OutDataJSON: `{
                        "UnionField": {
                            "boolean": true
                        }
                    }`,
	}, {
		TestName: "withInt",
		InDataJSON: `{
                        "UnionField": {
                            "int": 999
                        }
                    }`,
		OutDataJSON: `{
                        "UnionField": {
                            "int": 999
                        }
                    }`,
	}, {
		TestName: "withNull",
		InDataJSON: `{
                        "UnionField": null
                    }`,
		OutDataJSON: `{
                        "UnionField": null
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
