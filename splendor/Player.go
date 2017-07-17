// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor

import (
	"fmt"
	"reflect"
)

type Player struct {
	hand   Hand
	broker Broker
}

func NewPlayer(hand *Hand, broker *Broker) *Player {
	return &Player{*hand, *broker};
}

func (player *Player) GetMoves() []Move {
	moves := []Move{}
	cardMoves(player, &moves)
	chipMoves(player, &moves)

	return moves
}

func cardMoves(player *Player, moves *[]Move) {
	for _, card := range player.broker.getOpenedCards() {
		need := substr(player.hand.bonus, card.cost)
		need = add(need, player.hand.chips)

		if (need["g"] > player.hand.chips["g"]) {
			continue
		}

		nobles := calculateNobles(card,player);
		*moves = append(*moves, Move{need, card, *nobles})
	}
}
func calculateNobles(card Card, p *Player) *[]Card {
	myNobles := []Card{}
	nobles := p.broker.board.noble;

	for _, n := range  nobles{
		receivesNoble:=true
		for k,v :=range n.cost {
			if (p.hand.bonus[k] + (Ite(k==card.bonus,1,0)).(int)- v < 0){
				receivesNoble=false
				break
			}
		}
		if (receivesNoble){
			myNobles=append(myNobles, n)
		}

	}

	return &myNobles

}

func chipMoves(player *Player, moves *[]Move) {
	candidateGemMoves := []Move{}
	gems := make([]string, 0)
	for k, cnt := range player.broker.getGems() {
		if (k == "g") {
			continue
		}
		//Take 2 gem tokens of the same color
		if (cnt >= 4) {
			m := Gems{k: 2}
			candidateGemMoves = append(candidateGemMoves, Move{gems:m})
		}
		//prepare for Take 3 gem tokens of different colors
		if (cnt >= 1) {
			gems = append(gems, k)
		}
	}
	//Take 3 gem tokens of different colors
	GetDifferentGems(gems, len(gems), 3 , &candidateGemMoves, 1)
	player.process(&candidateGemMoves, moves)

}

func (p *Player) process(candidateGemMoves *[]Move, moves *[]Move) {

	for _, candidate := range *candidateGemMoves {
		m := (CHIPS_LIMIT - p.hand.chips.count() - candidate.gems.count())

		if (m < 0) {
			//over the limit
			gemList := getGemListWithReturnTokens(p, candidate, -m)
			subMoves := []Move{}
			GetDifferentGems(gemList, len(gemList), -m, &subMoves, -1)
			for _, sm := range subMoves {
				for k, v := range candidate.gems {
					sm.gems[k] += v;
				}
				appendIfMissing(moves, sm)
			}
		} else {
			//good to merge
			*moves = append(*moves, candidate)
		}

	}
}
func appendIfMissing(moves *[]Move, m2 Move) {
	present := false
	for _, m := range *moves {
		eq := reflect.DeepEqual(m.gems, m2.gems)
		if eq {
			present = true
		}
	}
	if !present {
		*moves = append(*moves, m2)
	}
}

func getGemListWithReturnTokens(p *Player, candidate Move, overlimit int) []string {
	gems := make([]string, 0)
	for k, v := range p.hand.chips {
		for i := 0; i < overlimit && i < v; i++ {
			gems = append(gems, k)
		}
	}
	for k, v := range candidate.gems {
		for i := 0; i < overlimit && i < v; i++ {
			gems = append(gems, k)
		}
	}
	return gems
}

func (p *Player) ToString() string {
	return fmt.Sprintf("\tBonus: %v\n\tChips: %v", p.hand.bonus, p.hand.chips)
}
