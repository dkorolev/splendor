// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor

type Hand struct {
	chips Gems
	bonus Gems
	score int
	//debug fields
	bonusCards []Card
	noble      []Card
}

func NewHand(chips *Gems, bonus *Gems) *Hand {
	return &Hand{chips: *chips, bonus: *bonus}
}
