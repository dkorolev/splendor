// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor


//chips
const (
	diamond  = iota
	sapphire
	emerald
	ruby
	onyx
	gold
)

type Gems map[string]int

func (gems Gems) count() int {
	cnt := 0
	for _, v := range gems {
		cnt += v
	}
	return cnt
}

type Card struct {
	tier   int
	points int
	bonus  string
	cost   Gems
}



