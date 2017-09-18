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
	"strings"
	"time"

	"github.com/pkg/errors"
)

// TimeLayout helps to parse a date string of the format YYYY-MM-DD
//   Intended to be used with the following function:
// 	 time.Parse(TimeLayout, date)
var TimeLayout = "2006-01-02" //this represents YYYY-MM-DD

// ParseDateRange parses a date range string of the format start:end
//   where the start and end date are of the format YYYY-MM-DD.
//   The parsed dates are time.Time and will return the zero time for
//   unbounded dates, ex:
//   unbounded start:	:2000-12-31
//	 unbounded end: 	2000-12-31:
func ParseDateRange(dateRange string) (startDate, endDate time.Time, err error) {
	dates := strings.Split(dateRange, ":")
	if len(dates) != 2 {
		err = errors.New("bad date range, must be in format date:date")
		return
	}
	parseDate := func(date string) (out time.Time, err error) {
		if len(date) == 0 {
			return
		}
		out, err = time.Parse(TimeLayout, date)
		return
	}
	startDate, err = parseDate(dates[0])
	if err != nil {
		return
	}
	endDate, err = parseDate(dates[1])
	if err != nil {
		return
	}
	return
}
