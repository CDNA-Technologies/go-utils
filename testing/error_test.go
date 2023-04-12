package testing

import (
	"fmt"
	"testing"
)

func TestIsErrorEqual(t *testing.T) {
	tests := []struct {
		wantErr error
		err     error
		want    bool
	}{
		{
			wantErr: nil,
			err:     nil,
			want:    true,
		},
		{
			wantErr: nil,
			err:     fmt.Errorf("Error"),
			want:    false,
		},
		{
			wantErr: fmt.Errorf("Error"),
			err:     nil,
			want:    false,
		},
		{
			wantErr: fmt.Errorf("Error 1"),
			err:     fmt.Errorf("Error 2"),
			want:    false,
		},
		{
			wantErr: fmt.Errorf("Error"),
			err:     fmt.Errorf("Error"),
			want:    true,
		},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("IsErrorEqual(%v, %v)", input.wantErr, input.err), func(t *testing.T) {
			if got := IsErrorEqual(input.wantErr, input.err); got != input.want {
				t.Errorf("IsErrorEqual() = %v, want %v", got, input.want)
			}
		})
	}
}
