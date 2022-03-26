package time

var (
	demFakeTime Duration = 0
)

func SetDemFakeTime(fd Duration) {
	demFakeTime = fd
}
