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

package flags_test

import (
	"bytes"
	"strings"
	"testing"

	tmflags "github.com/tendermint/tmlibs/cli/flags"
	"github.com/tendermint/tmlibs/log"
)

const (
	defaultLogLevelValue = "info"
)

func TestParseLogLevel(t *testing.T) {
	var buf bytes.Buffer
	jsonLogger := log.NewTMJSONLogger(&buf)

	correctLogLevels := []struct {
		lvl              string
		expectedLogLines []string
	}{
		{"mempool:error", []string{
			``, // if no default is given, assume info
			``,
			`{"_msg":"Mesmero","level":"error","module":"mempool"}`,
			`{"_msg":"Mind","level":"info","module":"state"}`, // if no default is given, assume info
			``}},

		{"mempool:error,*:debug", []string{
			`{"_msg":"Kingpin","level":"debug","module":"wire"}`,
			``,
			`{"_msg":"Mesmero","level":"error","module":"mempool"}`,
			`{"_msg":"Mind","level":"info","module":"state"}`,
			`{"_msg":"Gideon","level":"debug"}`}},

		{"*:debug,wire:none", []string{
			``,
			`{"_msg":"Kitty Pryde","level":"info","module":"mempool"}`,
			`{"_msg":"Mesmero","level":"error","module":"mempool"}`,
			`{"_msg":"Mind","level":"info","module":"state"}`,
			`{"_msg":"Gideon","level":"debug"}`}},
	}

	for _, c := range correctLogLevels {
		logger, err := tmflags.ParseLogLevel(c.lvl, jsonLogger, defaultLogLevelValue)
		if err != nil {
			t.Fatal(err)
		}

		logger = logger

		buf.Reset()

		logger.With("module", "wire").Debug("Kingpin")
		if have := strings.TrimSpace(buf.String()); c.expectedLogLines[0] != have {
			t.Errorf("\nwant '%s'\nhave '%s'\nlevel '%s'", c.expectedLogLines[0], have, c.lvl)
		}

		buf.Reset()

		logger.With("module", "mempool").Info("Kitty Pryde")
		if have := strings.TrimSpace(buf.String()); c.expectedLogLines[1] != have {
			t.Errorf("\nwant '%s'\nhave '%s'\nlevel '%s'", c.expectedLogLines[1], have, c.lvl)
		}

		buf.Reset()

		logger.With("module", "mempool").Error("Mesmero")
		if have := strings.TrimSpace(buf.String()); c.expectedLogLines[2] != have {
			t.Errorf("\nwant '%s'\nhave '%s'\nlevel '%s'", c.expectedLogLines[2], have, c.lvl)
		}

		buf.Reset()

		logger.With("module", "state").Info("Mind")
		if have := strings.TrimSpace(buf.String()); c.expectedLogLines[3] != have {
			t.Errorf("\nwant '%s'\nhave '%s'\nlevel '%s'", c.expectedLogLines[3], have, c.lvl)
		}

		buf.Reset()

		logger.Debug("Gideon")
		if have := strings.TrimSpace(buf.String()); c.expectedLogLines[4] != have {
			t.Errorf("\nwant '%s'\nhave '%s'\nlevel '%s'", c.expectedLogLines[4], have, c.lvl)
		}
	}

	incorrectLogLevel := []string{"some", "mempool:some", "*:some,mempool:error"}
	for _, lvl := range incorrectLogLevel {
		if _, err := tmflags.ParseLogLevel(lvl, jsonLogger, defaultLogLevelValue); err == nil {
			t.Fatalf("Expected %s to produce error", lvl)
		}
	}
}
