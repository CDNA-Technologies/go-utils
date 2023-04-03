package converter

import (
	"fmt"
	"testing"
)

func TestSecsToMilliSecs(t *testing.T) {
	tests := []struct {
		secs int
		want int
	}{
		{-40, -40000},
		{0, 0},
		{1, 1000},
		{50, 50000},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestSecsToMilliSecs(%v)", input.secs), func(t *testing.T) {
			if got := SecsToMilliSecs(input.secs); got != input.want {
				t.Errorf("SecsToMilliSecs() = %v, want %v", got, input.want)
			}
		})
	}
}

func TestSecsToHMS(t *testing.T) {
	tests := []struct {
		secs     int
		wantHrs  int
		wantMins int
		wantSecs int
	}{
		{-67, 0, -1, -7},
		{-5, 0, 0, -5},
		{0, 0, 0, 0},
		{43, 0, 0, 43},
		{76, 0, 1, 16},
		{174, 0, 2, 54},
		{600, 0, 10, 0},
		{6341, 1, 45, 41},
		{82357, 22, 52, 37},
		{91952342, 25542, 19, 2},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestSecsToHMS(%v)", input.secs), func(t *testing.T) {
			h, m, s := SecsToHMS(input.secs)
			if input.wantHrs != h {
				t.Errorf("SecsToHMS() hrs = %v, want %v", h, input.wantHrs)
			}
			if input.wantMins != m {
				t.Errorf("SecsToHMS() mins = %v, want %v", m, input.wantMins)
			}
			if input.wantSecs != s {
				t.Errorf("SecsToHMS() secs = %v, want %v", s, input.wantSecs)
			}
		})
	}
}

func TestSecsToDHMS(t *testing.T) {
	tests := []struct {
		secs     int
		wantDays int
		wantHrs  int
		wantMins int
		wantSecs int
	}{
		{-67, 0, 0, -1, -7},
		{-5, 0, 0, 0, -5},
		{0, 0, 0, 0, 0},
		{43, 0, 0, 0, 43},
		{76, 0, 0, 1, 16},
		{174, 0, 0, 2, 54},
		{600, 0, 0, 10, 0},
		{6341, 0, 1, 45, 41},
		{82357, 0, 22, 52, 37},
		{86400, 1, 0, 0, 0},
		{91952342, 1064, 6, 19, 2},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestSecsToDHMS(%v)", input.secs), func(t *testing.T) {
			d, h, m, s := SecsToDHMS(input.secs)
			if input.wantDays != d {
				t.Errorf("SecsToDHMS() days = %v, want %v", d, input.wantDays)
			}
			if input.wantHrs != h {
				t.Errorf("SecsToDHMS() hrs = %v, want %v", h, input.wantHrs)
			}
			if input.wantMins != m {
				t.Errorf("SecsToDHMS() mins = %v, want %v", m, input.wantMins)
			}
			if input.wantSecs != s {
				t.Errorf("SecsToDHMS() secs = %v, want %v", s, input.wantSecs)
			}
		})
	}
}

func TestMinsToSecs(t *testing.T) {
	tests := []struct {
		mins int
		want int
	}{
		{-40, -2400},
		{0, 0},
		{1, 60},
		{41, 2460},
		{50, 3000},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestMinsToSecs(%v)", input.mins), func(t *testing.T) {
			if got := MinsToSecs(input.mins); got != input.want {
				t.Errorf("MinsToSecs() = %v, want %v", got, input.want)
			}
		})
	}
}

func TestMinsToHM(t *testing.T) {
	tests := []struct {
		mins     int
		wantHrs  int
		wantMins int
	}{
		{-67, -1, -7},
		{-5, 0, -5},
		{0, 0, 0},
		{43, 0, 43},
		{76, 1, 16},
		{174, 2, 54},
		{600, 10, 0},
		{6341, 105, 41},
		{82357, 1372, 37},
		{86400, 1440, 0},
		{91952342, 1532539, 2},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestMinsToHM(%v)", input.mins), func(t *testing.T) {
			h, m := MinsToHM(input.mins)
			if input.wantHrs != h {
				t.Errorf("MinsToHM() hrs = %v, want %v", h, input.wantHrs)
			}
			if input.wantMins != m {
				t.Errorf("MinsToHM() mins = %v, want %v", m, input.wantMins)
			}
		})
	}
}

func TestMinsToDHM(t *testing.T) {
	tests := []struct {
		mins     int
		wantDays int
		wantHrs  int
		wantMins int
	}{
		{-67, 0, -1, -7},
		{-5, 0, 0, -5},
		{0, 0, 0, 0},
		{43, 0, 0, 43},
		{76, 0, 1, 16},
		{174, 0, 2, 54},
		{600, 0, 10, 0},
		{6341, 4, 9, 41},
		{82357, 57, 4, 37},
		{86400, 60, 0, 0},
		{91952342, 63855, 19, 2},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestMinsToDHM(%v)", input.mins), func(t *testing.T) {
			d, h, m := MinsToDHM(input.mins)
			if input.wantDays != d {
				t.Errorf("MinsToDHM() days = %v, want %v", d, input.wantDays)
			}
			if input.wantHrs != h {
				t.Errorf("MinsToDHM() hrs = %v, want %v", h, input.wantHrs)
			}
			if input.wantMins != m {
				t.Errorf("MinsToDHM() mins = %v, want %v", m, input.wantMins)
			}
		})
	}
}

func TestHrsToMins(t *testing.T) {
	tests := []struct {
		hrs  int
		want int
	}{
		{-40, -2400},
		{0, 0},
		{1, 60},
		{41, 2460},
		{50, 3000},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestHrsToMins(%v)", input.hrs), func(t *testing.T) {
			if got := HrsToMins(input.hrs); got != input.want {
				t.Errorf("HrsToMins() = %v, want %v", got, input.want)
			}
		})
	}
}

func TestHrsToSecs(t *testing.T) {
	tests := []struct {
		hrs  int
		want int
	}{
		{-40, -144000},
		{0, 0},
		{1, 3600},
		{41, 147600},
		{50, 180000},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestHrsToSecs(%v)", input.hrs), func(t *testing.T) {
			if got := HrsToSecs(input.hrs); got != input.want {
				t.Errorf("HrsToSecs() = %v, want %v", got, input.want)
			}
		})
	}
}
