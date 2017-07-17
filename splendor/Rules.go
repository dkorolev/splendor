// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.s

package splendor

type Setup struct {
	gems     int
	gold     int
	nobles   int
	numCards int
}

type Deck map[int][]Card

var CHIPS_LIMIT = 10

var rules = map[int]Setup{
	2: {4,5,3, 5},
	3: {5,5,4, 5},
	4: {7,5,5, 5},
}

var deckLevel3 = []Card{
	{3, 4, "d", Gems{"o": 7}},
	{3, 4, "s", Gems{"d": 7}},
	{3, 4, "e", Gems{"s": 7}},
	{3, 4, "r", Gems{"e": 7}},
	{3, 4, "o", Gems{"r": 7}},

	{3, 5, "d", Gems{"o": 7, "d": 3}},
	{3, 5, "s", Gems{"d": 7, "s": 3}},
	{3, 5, "e", Gems{"s": 7, "e": 3}},
	{3, 5, "r", Gems{"e": 7, "r": 3}},
	{3, 5, "o", Gems{"r": 7, "o": 3}},

	{3, 4, "d", Gems{"d": 3, "r": 3, "s": 6}},
	{3, 4, "s", Gems{"s": 3, "o": 3, "d": 6}},
	{3, 4, "e", Gems{"d": 3, "e": 3, "s": 6}},
	{3, 4, "r", Gems{"s": 3, "r": 3, "e": 6}},
	{3, 4, "o", Gems{"e": 3, "o": 3, "r": 6}},

	{3, 3, "d", Gems{"s": 3, "e": 3, "o": 3, "r": 5}},
	{3, 3, "e", Gems{"s": 3, "r": 3, "o": 3, "d": 5}},
	{3, 3, "o", Gems{"s": 3, "d": 3, "r": 3, "e": 5}},
	{3, 3, "s", Gems{"d": 3, "e": 3, "r": 3, "o": 6}},
	{3, 3, "r", Gems{"d": 3, "e": 3, "o": 3, "s": 6}},
}

var deckLevel2 = []Card{
	{2, 3, "d", Gems{"d": 6}},
	{2, 3, "s", Gems{"s": 6}},
	{2, 3, "e", Gems{"e": 6}},
	{2, 3, "r", Gems{"r": 6}},
	{2, 3, "o", Gems{"o": 6}},

	{2, 2, "d", Gems{"r": 5}},
	{2, 2, "s", Gems{"s": 5}},
	{2, 2, "e", Gems{"e": 5}},
	{2, 2, "r", Gems{"o": 5}},
	{2, 2, "o", Gems{"d": 5}},

	{2, 2, "d", Gems{"r": 5, "o": 3}},
	{2, 2, "o", Gems{"e": 5, "r": 3}},
	{2, 2, "e", Gems{"s": 5, "e": 3}},
	{2, 2, "s", Gems{"d": 5, "s": 3}},
	{2, 2, "r", Gems{"o": 5, "d": 3}},

	{2, 2, "e", Gems{"d": 4, "s": 2, "o": 1}},
	{2, 2, "r", Gems{"s": 4, "e": 2, "d": 1}},
	{2, 2, "o", Gems{"e": 4, "r": 2, "s": 1}},
	{2, 2, "d", Gems{"r": 4, "o": 2, "e": 1}},
	{2, 2, "s", Gems{"o": 4, "d": 2, "r": 1}},

	{2, 1, "s", Gems{"r": 3, "s": 2, "e": 2}},
	{2, 1, "r", Gems{"o": 3, "d": 2, "r": 2}},
	{2, 1, "o", Gems{"d": 3, "s": 2, "e": 2}},
	{2, 1, "e", Gems{"s": 3, "d": 2, "o": 2}},
	{2, 1, "d", Gems{"e": 3, "r": 2, "o": 2}},

	{2, 1, "s", Gems{"e": 3, "o": 3, "s": 2}},
	{2, 1, "o", Gems{"e": 3, "d": 3, "o": 2}},
	{2, 1, "e", Gems{"d": 3, "r": 3, "e": 2}},
	{2, 1, "d", Gems{"s": 3, "r": 3, "d": 2}},
	{2, 1, "r", Gems{"s": 3, "o": 3, "d": 2}},
}

var deckLevel1 = []Card{

	{1, 1, "d", Gems{"e": 4}},
	{1, 1, "s", Gems{"r": 4}},
	{1, 1, "e", Gems{"o": 4}},
	{1, 1, "r", Gems{"d": 4}},
	{1, 1, "o", Gems{"s": 4}},

	{1, 0, "s", Gems{"o": 3}},
	{1, 0, "d", Gems{"d": 3}},
	{1, 0, "r", Gems{"r": 3}},
	{1, 0, "e", Gems{"e": 3}},
	{1, 0, "o", Gems{"o": 3}},

	{1, 0, "e", Gems{"s": 3, "d": 1, "e": 1}},
	{1, 0, "s", Gems{"e": 3, "s": 1, "r": 1}},
	{1, 0, "o", Gems{"r": 3, "e": 1, "o": 1}},
	{1, 0, "r", Gems{"o": 3, "r": 1, "d": 1}},
	{1, 0, "d", Gems{"d": 3, "s": 1, "o": 1}},

	{1, 0, "e", Gems{"r": 2, "o": 2, "s": 1}},
	{1, 0, "s", Gems{"r": 2, "e": 2, "d": 1}},
	{1, 0, "d", Gems{"s": 2, "e": 2, "o": 1}},
	{1, 0, "o", Gems{"s": 2, "d": 2, "r": 1}},
	{1, 0, "r", Gems{"o": 2, "d": 2, "e": 1}},

	{1, 0, "d", Gems{"s": 2, "o": 2}},
	{1, 0, "e", Gems{"s": 2, "r": 2}},
	{1, 0, "r", Gems{"d": 2, "r": 2}},
	{1, 0, "o", Gems{"d": 2, "e": 2}},
	{1, 0, "s", Gems{"e": 2, "o": 2}},

	{1, 0, "d", Gems{"e": 2, "s": 1, "r": 1, "o": 1}},
	{1, 0, "r", Gems{"d": 2, "s": 1, "e": 1, "o": 1}},
	{1, 0, "s", Gems{"r": 2, "d": 1, "e": 1, "o": 1}},
	{1, 0, "o", Gems{"s": 2, "d": 1, "e": 1, "r": 1}},
	{1, 0, "e", Gems{"o": 2, "d": 1, "s": 1, "r": 1}},

	{1, 0, "o", Gems{"e": 2, "r": 1}},
	{1, 0, "s", Gems{"o": 2, "d": 1}},
	{1, 0, "r", Gems{"s": 2, "e": 1}},
	{1, 0, "w", Gems{"r": 2, "o": 1}},
	{1, 0, "e", Gems{"d": 2, "s": 1}},

	{1, 0, "o", Gems{"d": 1, "s": 1, "e": 1, "r": 1}},
	{1, 0, "r", Gems{"d": 1, "s": 1, "e": 1, "o": 1}},
	{1, 0, "s", Gems{"d": 1, "e": 1, "r": 1, "o": 1}},
	{1, 0, "e", Gems{"d": 1, "s": 1, "r": 1, "o": 1}},
	{1, 0, "d", Gems{"s": 1, "e": 1, "r": 1, "o": 1}},
}

var deckNoble = []Card{
	{points:3, cost:Gems{"d": 3, "s": 3, "e": 3}},
	{points:3, cost:Gems{"d": 3, "s": 3, "o": 3}},
	{points:3, cost:Gems{"d": 3, "r": 3, "o": 3}},
	{points:3, cost:Gems{"o": 3, "r": 3, "e": 3}},
	{points:3, cost:Gems{"e": 3, "s": 3, "r": 3}},
	{points:3, cost:Gems{"s": 4, "e": 4}},
	{points:3, cost:Gems{"s": 4, "d": 4}},
	{points:3, cost:Gems{"d": 4, "o": 4}},
	{points:3, cost:Gems{"r": 4, "o": 4}},
	{points:3, cost:Gems{"r": 4, "e": 4}},
}

var deck = Deck{
	3: deckLevel3,
	2: deckLevel2,
	1: deckLevel1,
}
