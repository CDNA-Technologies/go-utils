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

/**
	Converts seconds to milliseconds.
**/
func SecsToMilliSecs(secs int) int {
	return secs * MilliSecsPerSec
}

/**
	Converts seconds to hours, minutes and seconds.
**/
func SecsToHMS(secs int) (int, int, int) {
	hrs := secs / SecsPerHr
	secs = secs % SecsPerHr
	mins := secs / SecsPerMin
	secs = secs % SecsPerMin
	return hrs, mins, secs
}

/**
	Converts seconds to days, hours, minutes and seconds.
**/
func SecsToDHMS(secs int) (int, int, int, int) {
	days := secs / SecsPerDay
	secs %= SecsPerDay
	hrs := secs / SecsPerHr
	secs %= MinsPerHr * 60
	mins := secs / 60
	secs %= 60
	return days, hrs, mins, secs
}

/**
	Converts minutes to seconds.
**/
func MinsToSecs(mins int) int {
	return mins * SecsPerMin
}

/**
	Converts minutes to hours and minutes.
**/
func MinsToHM(mins int) (int, int) {
	hrs := mins / MinsPerHr
	mins = mins % MinsPerHr
	return hrs, mins
}

/**
	Converts minutes to days, hours and minutes.
**/
func MinsToDHM(mins int) (int, int, int) {
	days := mins / MinsPerDay
	mins %= MinsPerDay
	hours := mins / MinsPerHr
	mins %= MinsPerHr

	return days, hours, mins
}

/**
	Converts hours to seconds.
**/
func HrsToSecs(hrs int) int {
	return hrs * SecsPerHr
}

/**
	Converts hours to minutes.
**/
func HrsToMins(hrs int) int {
	return hrs * MinsPerHr
}
