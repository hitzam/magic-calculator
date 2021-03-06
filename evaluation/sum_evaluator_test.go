package evaluation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumEvaluator(t *testing.T) {

	testCases := []struct {
		name                  string
		x, y                  int64
		expectedResult        SingleResult
		expectedWrittenResult string
	}{
		{
			name:                  "x=0 and y=0 result should be 0",
			x:                     0,
			y:                     0,
			expectedResult:        SingleResult(0),
			expectedWrittenResult: "0\n",
		},
		{
			name:                  "x=2 and y=0 result should be 2",
			x:                     2,
			y:                     0,
			expectedResult:        SingleResult(2),
			expectedWrittenResult: "2\n",
		},
		{
			name:                  "x=2 and y=3 result should be 5",
			x:                     2,
			y:                     3,
			expectedResult:        SingleResult(5),
			expectedWrittenResult: "5\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := &strings.Builder{}
			se := &SumEvaluator{X: tc.x, Y: tc.y, ResultWriter: sb}
			var evaluator Evaluator = se
			err := evaluator.Evaluate()

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedResult, se.Result)
			assert.Equal(t, tc.expectedWrittenResult, sb.String())
		})
	}
}
