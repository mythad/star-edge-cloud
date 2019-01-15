package logger

import (
	"fmt"
	"sort"
	"time"
	"sync"
)

var tickLog = GetLogger("stats")

// The interval between logging the number of actions/sec
var ReportInterval = time.Second * 5

var ticks = make(map[string]int)
var mu = sync.Mutex{}

func init() {
	go func() {
		for {
			time.Sleep(ReportInterval)

			if len(ticks) == 0 {
				continue
			}

			ticks2 := make(map[string]int)

			mu.Lock()
			keys := make([]string, 0, len(ticks))
			for key := range ticks {
				keys = append(keys, key)
			}
			mu.Unlock()
			sort.Strings(keys)

			var msg string
			mu.Lock()
			for _, name := range keys {
				ticks2[name] = 0
				count := ticks[name]
				msg += fmt.Sprintf("[%s: %.f/sec] ", name, float64(count)/(float64(ReportInterval)/float64(time.Second)))
			}
			mu.Unlock()

			ticks = ticks2

			tickLog.Infof(msg)
		}

	}()
}

func Tick(name string) {
	TickN(name, 1)
}

func TickN(name string, number int) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := ticks[name]; !ok {
		ticks[name] = 0
	}
	ticks[name] += number
}
