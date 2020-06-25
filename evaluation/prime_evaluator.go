package evaluation

import (
	"fmt"
	"io"
)

type FirstPrimeEvaluator struct {
	Result       SequenceResult
	N            int64
	ResultWriter io.Writer
}

func (e *FirstPrimeEvaluator) Evaluate() {
	e.Result = SequenceResult{}
	for v := int64(2); len(e.Result) < int(e.N); v++ {
		if e.isPrime(v) {
			e.Result = append(e.Result, v)
		}
	}

	fmt.Fprintf(e.ResultWriter, "Result: %s\n", FormatSequenceResult(e.Result))
}

func (e *FirstPrimeEvaluator) isPrime(v int64) bool {
	for _, prime := range e.Result {
		if v%prime == 0 {
			return false
		}
	}

	return true
}
