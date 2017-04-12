package date

import (
	"fmt"
)

type Chunk struct {
	Start Date
	End   Date
}

func (c Chunk) String() string {
	return fmt.Sprintf("%v..%v", c.Start, c.End)
}

func Chunks(start, end Date, chunkSize int) []Chunk {
	var chunks []Chunk

	first := start
	for !first.After(end) {
		last := Min(end, first.AddDays(chunkSize-1))
		chunks = append(chunks, Chunk{first, last})

		first = last.Next()
	}

	return chunks
}

func NextChunk(start, end Date, chunkSize int) (Chunk, Date, int) {
	last := Min(end, start.AddDays(chunkSize-1))
	nextStart := last.Next()
	var remDays int
	if nextStart.After(end) {
		remDays = 0
	} else {
		remDays = 1 + int(end.tm.Sub(nextStart.tm).Hours()/24)
	}

	return Chunk{start, last}, nextStart, (remDays + chunkSize - 1) / chunkSize
}
