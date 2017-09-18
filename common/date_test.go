// Copyright 2017 Tendermint. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	date  = time.Date(2015, time.Month(12), 31, 0, 0, 0, 0, time.UTC)
	date2 = time.Date(2016, time.Month(12), 31, 0, 0, 0, 0, time.UTC)
	zero  time.Time
)

func TestParseDateRange(t *testing.T) {
	assert := assert.New(t)

	var testDates = []struct {
		dateStr string
		start   time.Time
		end     time.Time
		errNil  bool
	}{
		{"2015-12-31:2016-12-31", date, date2, true},
		{"2015-12-31:", date, zero, true},
		{":2016-12-31", zero, date2, true},
		{"2016-12-31", zero, zero, false},
		{"2016-31-12:", zero, zero, false},
		{":2016-31-12", zero, zero, false},
	}

	for _, test := range testDates {
		start, end, err := ParseDateRange(test.dateStr)
		if test.errNil {
			assert.Nil(err)
			testPtr := func(want, have time.Time) {
				assert.True(have.Equal(want))
			}
			testPtr(test.start, start)
			testPtr(test.end, end)
		} else {
			assert.NotNil(err)
		}
	}
}
