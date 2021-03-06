// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanfix_test

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/scanfix"
	"github.com/protoscan/protoscan"
)

var ScanFieldTestCases = []struct {
	line         int
	input        string
	expectedTags []string
	expectedVals []string
}{
	{
		line:         line(),
		input:        "8=FIX.4.4|9=70|35=A|49=FG|56=tgFZctx20008c|34=5|52=20210224-12:11:42.000394000|98=0|108=30|10=173|",
		expectedTags: []string{"8", "9", "35", "49", "56", "34", "52", "98", "108", "10"},
		expectedVals: []string{"FIX.4.4", "70", "A", "FG", "tgFZctx20008c", "5", "20210224-12:11:42.000394000", "0", "30", "173"},
	},
}

func TestScanField(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range ScanFieldTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			input := strings.ReplaceAll(tc.input, "|", "\x01")
			r := bytes.NewReader([]byte(input))
			split := &scanfix.Field{Format: format}
			scan := protoscan.Protoscan{
				Reader: r,
				Split:  split.Split,
			}

			var tags []string
			var vals []string

			for scan.Scan() {
				tags = append(tags, strconv.Itoa(split.Tag))
				vals = append(vals, string(scan.Token))

				expectedGap := fmt.Sprintf("%d=\x01", split.Tag)
				recievedGap := string(split.Gaps)
				if recievedGap != expectedGap {
					t.Fatalf("unexpected scan gap, expected, %q, recieved: %q %s", expectedGap, recievedGap, linkToExample)
				}
			}

			err := scan.Err()
			if err != nil {
				t.Fatalf("unexpected scan error: %v %s", err, linkToExample)
			}

			gap := split.Gaps
			if len(gap) != 0 {
				t.Fatalf("unexpected scan gap: %v %s", gap, linkToExample)
			}

			if strings.Join(tags, "") != strings.Join(tc.expectedTags, "") {
				t.Errorf("unexpected tag, expected: %#v, received: %#v %s", tc.expectedTags, tags, linkToExample)
			}

			if strings.Join(vals, "") != strings.Join(tc.expectedVals, "") {
				t.Errorf("unexpected values, expected: %q, received: %q %s", tc.expectedVals, vals, linkToExample)
			}
		})
	}
}

var format = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:    f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		MsgType35:       f0.Fld{Req, f0.ASCII, f0.StringDefault(), f0.Var{1, 2}},
		SenderCompID49:  f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 12}},
		TargetCompID56:  f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MsgSeqNum34:     f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		SendingTime52:   f0.Fld{Req, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Con{27}},
		EncryptMethod98: f0.Fld{Req, f0.ASCII, f0.IntDefault(0), f0.Con{1}},
		HeartBtInt108:   f0.Fld{Req, f0.ASCII, f0.SecondsDuration(30 * time.Second), f0.Var{1, 19}},
		Password554:     f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 8}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 19}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,
		BodyLength9,
		MsgType35,
		SenderCompID49,
		TargetCompID56,
		MsgSeqNum34,
		SendingTime52,
		EncryptMethod98,
		HeartBtInt108,
		Password554,
	},
}

const Req = true

const (
	BeginString8    = 8
	BodyLength9     = 9
	MsgType35       = 35
	SenderCompID49  = 49
	TargetCompID56  = 56
	MsgSeqNum34     = 34
	SendingTime52   = 52
	EncryptMethod98 = 98
	HeartBtInt108   = 108
	Password554     = 554
)
