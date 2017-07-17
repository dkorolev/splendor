// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor

import (
	"math/rand"
)

func shuffle(src []Card) []Card {
	dest := make([]Card, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}

	return dest
}

func substr(from Gems, what Gems) Gems {
	res := make(map[string]int)
	for _, gem := range []string{"d", "s", "e", "o", "r"} {
		res[gem] = from[gem] - what[gem]
	}

	return res
}

func add(need Gems, have Gems) Gems {
	res := make(map[string]int)
	for _, gem := range []string{"d", "s", "e", "o", "r"} {
		if (need[gem] < 0) {
			res[gem] = -need[gem]
			if (have[gem]+need[gem] < 0) {
				res["g"] += -(have[gem] + need[gem])
				res[gem] = have[gem];
			}
		}
	}

	for k, v := range res {
		res[k] = -v
	}
	return res
}

func Ite(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func GetDifferentGems(gems []string, n int, k int, moves *[]Move, sign int) {

	comb := make([]string, k)
	combinationHelper(gems, comb, 0, n-1, 0, k, moves, sign);
}

func combinationHelper(gems []string, comb []string, start int, stop int, idx int, k int, moves *[]Move, sign int) {
	if (idx == k) {
		gems := Gems{}
		for i := 0; i < k; i++ {
			gems[comb[i]] += (1 * sign)
		}
		*moves = append(*moves, Move{gems:gems})
		return
	}

	for i := start; i <= stop && stop-i+1 >= k-idx; i++ {
		comb[idx] = gems[i];
		combinationHelper(gems, comb, i+1, stop, idx+1, k, moves, sign);
	}
}
