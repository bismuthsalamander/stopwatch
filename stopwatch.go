package stopwatch

import (
	"fmt"
	"time"
)

var Watch Stopwatch

type Stopwatch struct {
	Buckets      map[string]int64
	BucketStarts map[string]int64
	EntryCounts  map[string]int64
}

func init() {
	Watch = Stopwatch{}
	Watch.Buckets = make(map[string]int64)
	Watch.BucketStarts = make(map[string]int64)
	Watch.EntryCounts = make(map[string]int64)
	Watch.Start("")
}

func (s *Stopwatch) Start(b string) {
	if _, ok := s.BucketStarts[b]; ok {
		return
	}
	s.EntryCounts[b]++
	s.BucketStarts[b] = time.Now().UnixNano()
}

func (s *Stopwatch) Stop(b string) {
	end := time.Now().UnixNano()
	start, ok := s.BucketStarts[b]
	if !ok {
		return
	}
	s.Buckets[b] += end - start
	delete(s.BucketStarts, b)
}

func (s *Stopwatch) Results() string {
	now := time.Now().UnixNano()
	out := ""
	for k, v := range s.Buckets {
		res := v
		if start, ok := s.BucketStarts[k]; ok {
			res += start - now
		}
		out += fmt.Sprintf("%s (%d): %.4f\n", k, s.EntryCounts[k], float64(res)/1000000000.0)
	}
	s.Stop("")
	out += fmt.Sprintf("TOTAL: %.4f\n", float64(s.Buckets[""])/1000000000.0)
	s.Start("")
	return out
}

func Start(s string) {
	Watch.Start(s)
}

func Stop(s string) {
	Watch.Stop(s)
}

func Results() string {
	return Watch.Results()
}

func Buckets() map[string]int64 {
	res := make(map[string]int64)
	for k, v := range Watch.Buckets {
		res[k] = v
	}
	return res
}
