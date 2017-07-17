// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor

import (
	"splendor"
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestCombination(*testing.T) {
	rand.Seed( time.Now().UTC().UnixNano())
	game := splendor.NewSplendor(2)
	broker := splendor.NewBroker(game)
	gems := splendor.Gems{"d": 1, "s": 1, "e": 1, "o": 1, "r": 1, "g": 3}
	bonus := splendor.Gems{"d": 3, "s": 2, "e": 2, "o": 3, "r": 2}
	hand := splendor.NewHand(&gems, &bonus);
	player := splendor.NewPlayer(hand, broker);
	moves := player.GetMoves()

	fmt.Print("\nPlayer:\n", player.ToString())
	fmt.Print("\nBrocker:\n")
	broker.ToString()
	fmt.Print("\nMoves:\n")
	for _, move := range moves {
		fmt.Print(move.ToString())
	}
}
