package lab2

import (
	"bytes"
	"errors"
	"strings"

	. "gopkg.in/check.v1"
)

type ComputeHandlerSuite struct{}

func init() {
	Suite(&ComputeHandlerSuite{})
}

func (s *ComputeHandlerSuite) TestCompute(c *C) {
	type testCase struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}

	var testTable = []testCase{
		{
			name:           "Test1",
			input:          "4 0 4 + +",
			expectedOutput: "(4 + (0 + 4))",
			expectedError:  nil,
		},
		{
			name:           "Test2",
			input:          "4 0 4",
			expectedOutput: "",
			expectedError:  errors.New("Incorrect postfix example"),
		},
		{
			name:           "Test3",
			input:          "2 3 2 $",
			expectedOutput: "",
			expectedError:  errors.New("Incorrect postfix example"),
		},
	}

	for _, test := range testTable {
		input := strings.NewReader(test.input)
		output := new(bytes.Buffer)
		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()

		outputString := output.String()

		c.Check(outputString, Equals, test.expectedOutput, Commentf("%s: actual output %s", test.name, outputString))
		c.Check(err, DeepEquals, test.expectedError, Commentf("%s: actual error %v", test.name, err))
	}
}
