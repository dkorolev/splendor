// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor

type Splendor struct {
	players int
	setup   Setup
}

func NewSplendor(players int) *Splendor {
	splendor := &Splendor{players, rules[players]}
	return splendor
}

func (s *Splendor) getPlayers() int {
	return s.players
}

func (s *Splendor) String() string {
	return string(s.players)
}

