package lab2

import (
	"errors"
	"fmt"
	gocheck "gopkg.in/check.v1"
	"testing"
)

type PostfixToInfixSuite struct{}

func TestPostfixToInfixSuite(t *testing.T) { gocheck.TestingT(t) }

func init() {
	gocheck.Suite(&PostfixToInfixSuite{})
}
func (s *PostfixToInfixSuite) TestPostfixToInfixSuite(c *gocheck.C) {
	type testCase struct {
		name           string
		input          string
		expectedResult string
		expectedError  error
	}

	var testTable = []testCase{

		{
			name:           "Test1",
			input:          "2 3 + 4 5 * -",
			expectedResult: "((2 + 3) - (4 * 5) !^%£$$^)",
			expectedError:  nil,
		},
		{
			name:           "Test2",
			input:          "1 2 + 3 4 * 5 / - 6 - 7 8 / - 9 +",
			expectedResult: "(((((1 + 2) - ((3 * 4) / 5)) - 6) - (7 / 8)) + 9)",
			expectedError:  nil,
		},
		{
			name:           "Test3",
			input:          "1 3 2 * + 6 2 / -",
			expectedResult: "((1 + (3 * 2)) - (6 / 2))",
			expectedError:  nil,
		},
		{
			name:           "Test4",
			input:          "7",
			expectedResult: "7",
			expectedError:  nil,
		},
		{
			name:           "Test5",
			input:          "2 @ 3 +",
			expectedResult: "",
			expectedError:  errors.New("Incorrect postfix example"),
		},
	}

	for _, test := range testTable {
		actual, err := PostfixToInfix(test.input)
		c.Assert(actual, gocheck.Equals, test.expectedResult, gocheck.Commentf("%s: actual result %s", test.name, actual))
		c.Assert(err, gocheck.DeepEquals, test.expectedError, gocheck.Commentf("%s: actual error %v", test.name, err))
	}
}

func (s *PostfixToInfixSuite) ExampleEvaluatePostfix(c *gocheck.C) {
	outputExample := func(name string, exp string) {
		result, error := PostfixToInfix(exp)
		if result != "" {
			fmt.Printf("%s: %s\n", name, result)
		}
		if error != nil {
			fmt.Printf("Error at %s: %v\n", name, error)
		}
	}
	outputExample("Test1", "2 3 + 4 5 * -")
	outputExample("Test5", "2 @ 3 +")
	outputExample("Test4", "7")
	outputExample("Test3", "1 3 2 * + 6 2 / -")
	outputExample("Test2", "1 2 + 3 4 * 5 / - 6 - 7 8 / - 9 +")

}
