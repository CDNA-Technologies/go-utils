package converter

const (
	MilliSecsPerSec = 1000
	SecsPerMin      = 60
	SecsPerHr       = 3600
	SecsPerDay      = HrsPerDay * SecsPerHr
	MinsPerHr       = 60
	MinsPerDay      = HrsPerDay * MinsPerHr
	HrsPerDay       = 24
)

func SecsToMilliSecs(secs int) int {
	return secs * MilliSecsPerSec
}

func SecsToHMS(secs int) (int, int, int) {
	hrs := secs / SecsPerHr
	secs = secs % SecsPerHr
	mins := secs / SecsPerMin
	secs = secs % SecsPerMin
	return hrs, mins, secs
}

func SecsToDHMS(secs int) (int, int, int, int) {
	days := secs / SecsPerDay
	secs %= SecsPerDay
	hrs := secs / SecsPerHr
	secs %= MinsPerHr * 60
	mins := secs / 60
	secs %= 60
	return days, hrs, mins, secs
}

func MinsToSecs(mins int) int {
	return mins * SecsPerMin
}

func MinsToHM(mins int) (int, int) {
	hrs := mins / MinsPerHr
	mins = mins % MinsPerHr
	return hrs, mins
}

func MinsToDHM(mins int) (int, int, int) {
	days := mins / MinsPerDay
	mins %= MinsPerDay
	hours := mins / MinsPerHr
	mins %= MinsPerHr

	return days, hours, mins
}

func HrsToSecs(hrs int) int {
	return hrs * SecsPerHr
}

func HrsToMins(hrs int) int {
	return hrs * MinsPerHr
}
