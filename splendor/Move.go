// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor

import "fmt"

type Move struct {
	gems Gems
	card Card
	nobles []Card
}

func (m *Move) GetCards() Card {
	return m.card
}

func (move *Move) getScore() int {
	return  Ite(len(move.nobles)>0,3,0).(int)+move.card.points

}

func (move *Move) ToString() string{

	return fmt.Sprintf("\t %v %v %v, %v\n",
		move.GetCards(),
		Ite(len(move.gems) > 0, move.gems, "free"),
		Ite(len(move.nobles) > 0, move.nobles, ""),
		move.getScore())
}


