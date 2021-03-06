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

package trace

import (
	"encoding/gob"

	"github.com/ampf11111/namazu/nmz/signal"
	signalutil "github.com/ampf11111/namazu/nmz/util/signal"
)

type SingleTrace struct {
	ActionSequence []signal.Action
}

func (this *SingleTrace) Equals(o *SingleTrace) bool {
	return signalutil.AreActionsSliceEqual(this.ActionSequence, o.ActionSequence)
}

func init() {
	gob.Register(SingleTrace{})
}
