// SPDX-License-Identifier: Apache-2.0
package how_much_is_a_goroutine

import (
	"fmt"
	"testing"
)

var fibonacci = []struct {
	iterations int
	result     int64
}{
	{
		0,
		0,
	},
	{
		1,
		1,
	},
	{
		2,
		1,
	},
	{
		3,
		2,
	},
	{
		4,
		3,
	},
	{
		5,
		5,
	},
	{
		6,
		8,
	},
	{
		7,
		13,
	},
	{
		8,
		21,
	},
	{
		9,
		34,
	},
	{
		10,
		55,
	},
	{
		11,
		89,
	},
	{
		12,
		144,
	},
}

func TestFibonacci(t *testing.T) {
	for i := 0; i < len(fibonacci); i++ {
		t.Run(fmt.Sprintf("number of iterations: %d", fibonacci[i].iterations), func(tt *testing.T) {
			result := GetFibonacciByIndex(fibonacci[i].iterations)
			if result != fibonacci[i].result {
				tt.Errorf("expected result: %d, encountered: %d", fibonacci[i].result, result)
			}
		})
	}
}
