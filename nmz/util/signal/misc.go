// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// can we remove this?
package signal

import . "github.com/ampf11111/namazu/nmz/signal"

// can we remove this?
func AreActionsSliceEqual(a, b []Action) bool {
	aLen := len(a)
	bLen := len(b)
	if aLen != bLen {
		return false
	}

	for i := 0; i < aLen; i++ {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}
