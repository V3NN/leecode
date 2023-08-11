package __jewelsAndStones

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	t.Log(numJewelsInStones("aA", "aAAbbbb"))
	t.Log(numJewelsInStones("z", "ZZ"))
	t.Log(numJewelsInStones("z", "ZZABz"))
}
