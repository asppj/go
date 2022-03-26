package time

var (
	demFakeTime int64 = 0
)

func fakeNow() (sec int64, nsec int32, mono int64) {
	_, _, mono = now()
	return diffNow(mono + demFakeTime)
}

func fakeRuntimeNano() int64 {
	return runtimeNano() + demFakeTime
}

func diffNow(fk int64) (sec int64, nsec int32, mono int64) {
	return fk / 1e9, int32(fk % 1e9), fk
}
