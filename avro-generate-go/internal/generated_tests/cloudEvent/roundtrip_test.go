// Code generated by generatetestcode.go; DO NOT EDIT.

package cloudEvent

import (
	"testing"

	"github.com/rogpeppe/avro/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "SomeEvent",
                "type": "record",
                "fields": [
                    {
                        "name": "Metadata",
                        "type": {
                            "name": "Metadata",
                            "type": "record",
                            "fields": [
                                {
                                    "name": "id",
                                    "type": "string"
                                },
                                {
                                    "name": "source",
                                    "type": "string"
                                },
                                {
                                    "name": "time",
                                    "type": "long"
                                }
                            ],
                            "namespace": "avro.apache.org"
                        }
                    },
                    {
                        "name": "other",
                        "type": "string"
                    }
                ],
                "namespace": "foo"
            }`,
	GoType: new(SomeEvent),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "Metadata": {
                            "time": 12345,
                            "id": "id1",
                            "source": "source1"
                        },
                        "other": "some other data"
                    }`,
		OutDataJSON: `{
                        "Metadata": {
                            "time": 12345,
                            "id": "id1",
                            "source": "source1"
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
