// Copyright 2017 marina.github@gmail.com All rights reserved.
// Use of this source code is governed under the GNU General Public License v3.0
// that can be found in the LICENSE file.

package splendor

type Board struct {
	gems  Gems
	cards map[int][]Card
	noble []Card
}
