package time

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	add   = "+"
	minus = "-"
	fat   = "2006-01-02 15:04:05"
)

var (
	fakeOnce           = sync.Once{}
	fakeTimeDiff int64 = 0
)

func fakeNow() (sec int64, nsec int32, mono int64) {
	fakeOnce.Do(func() {
		ft := os.Getenv("FAKETIME")
		fakeTimeDiff = parseFakeTime(ft)

	})
	_, _, mono = now()
	return diffNow(mono + fakeTimeDiff)
}

func diffNow(faketime int64) (sec int64, nsec int32, mono int64) {
	return faketime / 1e9, int32(faketime % 1e9), faketime
}
func parseFakeTime(ft string) (fk int64) {
	if ft == "" {
		return 0
	}
	if strings.HasPrefix(ft, add) {
		ft = ft[1:]
		return parseFakeTimeF(ft)
	}
	if strings.HasPrefix(ft, minus) {
		ft = ft[1:]
		return -parseFakeTimeF(ft)
	}
	t, err := time.Parse(fat, ft)
	if err != nil {
		panic(fmt.Errorf("parse fake_time format must is[%s] error: %s", fat, err))
	}
	_, _, nw := now()
	return t.UnixNano() - nw
}
func parseFakeTimeF(ft string) (fk int64) {
	var tmp string
	for _, r := range []rune(ft) {
		switch r {
		case 'd':
			n, err := strconv.Atoi(tmp)
			if err != nil {
				panic(err)
			}
			fk += 24 * 60 * 60 * 1e9 * int64(n)
			tmp = ""
		case 'h':
			n, err := strconv.Atoi(tmp)
			if err != nil {
				panic(err)
			}
			fk += 60 * 60 * 1e9 * int64(n)
			tmp = ""

		case 'm':
			n, err := strconv.Atoi(tmp)
			if err != nil {
				panic(err)
			}
			fk += 60 * 1e9 * int64(n)
			tmp = ""

		case 's':
			n, err := strconv.Atoi(tmp)
			if err != nil {
				panic(err)
			}
			fakeTimeDiff += 1e9 * int64(n)
			tmp = ""
		default:
			tmp += string(r)
		}
	}
	return fk
}
