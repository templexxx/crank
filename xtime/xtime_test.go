package xtime

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/templexxx/tsc"
)

func TestGetTimerEventNop(t *testing.T) {
	m := new(sync.Map)
	wg := new(sync.WaitGroup)
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(ii int) {
			defer wg.Done()
			t := time.NewTimer(-1)
			var tChan <-chan time.Time

			for j := 0; j < n; j++ {

				select {

				case <-tChan:
					m.Store(ii*n+j, tsc.UnixNano())
					tChan = nil
					continue
				default:

				}

				if tChan == nil {
					tChan = GetTimerEvent(t, -1)
				}
			}
		}(i)
	}
	wg.Wait()

	cnt := 0
	m.Range(func(key, value interface{}) bool {
		cnt++
		return true
	})
	assert.Greater(t, cnt, 1)
}
