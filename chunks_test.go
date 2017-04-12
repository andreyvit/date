package date

import (
	"fmt"
	"testing"
)

var chunkTests = []struct {
	start, end   string
	daysPerChunk int
	expected     string
}{
	{"2016-12-01", "2016-12-07", 1, "2016-12-01..2016-12-01, next=2016-12-02, rem=6"},
	{"2016-12-01", "2016-12-07", 2, "2016-12-01..2016-12-02, next=2016-12-03, rem=3"},
	{"2016-12-01", "2016-12-07", 3, "2016-12-01..2016-12-03, next=2016-12-04, rem=2"},
	{"2016-12-01", "2016-12-07", 4, "2016-12-01..2016-12-04, next=2016-12-05, rem=1"},
	{"2016-12-01", "2016-12-07", 5, "2016-12-01..2016-12-05, next=2016-12-06, rem=1"},
	{"2016-12-01", "2016-12-07", 6, "2016-12-01..2016-12-06, next=2016-12-07, rem=1"},
	{"2016-12-01", "2016-12-07", 7, "2016-12-01..2016-12-07, next=2016-12-08, rem=0"},
	{"2016-12-01", "2016-12-07", 8, "2016-12-01..2016-12-07, next=2016-12-08, rem=0"},

	{"2016-12-01", "2016-12-01", 1, "2016-12-01..2016-12-01, next=2016-12-02, rem=0"},
	{"2016-12-01", "2016-12-01", 100, "2016-12-01..2016-12-01, next=2016-12-02, rem=0"},
}

func TestChunk(t *testing.T) {
	for _, tt := range chunkTests {
		chunk, nextDay, n := NextChunk(MustParse(tt.start), MustParse(tt.end), tt.daysPerChunk)
		actualStr := NextChunkResultToString(chunk, nextDay, n)
		if actualStr != tt.expected {
			t.Errorf("NextChunk(%q, %q, %v) == %s, expected %s", tt.start, tt.end, tt.daysPerChunk, actualStr, tt.expected)
		} else {
			t.Logf("NextChunk(%q, %q, %v) == %s", tt.start, tt.end, tt.daysPerChunk, actualStr)
		}
	}
}

func NextChunkResultToString(chunk Chunk, nextDay Date, n int) string {
	return fmt.Sprintf("%v, next=%v, rem=%d", chunk, nextDay, n)
}
