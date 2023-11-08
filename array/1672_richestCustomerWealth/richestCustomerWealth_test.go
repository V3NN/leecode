package __richestCustomerWealth

import "testing"

func TestRichestCustomerWealth(t *testing.T) {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	t.Log(maximumWealth(accounts))

	accounts = [][]int{{1, 5}, {7, 3}, {3, 5}}
	t.Log(maximumWealth(accounts))
}
