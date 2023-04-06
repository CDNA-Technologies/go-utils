package currency

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLatestRates(t *testing.T) {
	type args struct {
		req LatestRatesRequest
	}
	tests := []struct {
		args    args
		want    LatestRatesResponse
		wantErr bool
	}{
		{
			args: args{
				req: LatestRatesRequest{
					WantCurrencies: []string{"INR"},
					Amount:         1,
					RoundPlaces:    2,
				},
			},
			want:    LatestRatesResponse{},
			wantErr: false,
		},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestGetLatestRates(%+v)", input.args.req), func(t *testing.T) {
			got, err := GetLatestRates(input.args.req)
			if (err != nil) != input.wantErr {
				t.Errorf("GetLatestRates() error = %v, wantErr %v", err, input.wantErr)
				return
			}
			if !reflect.DeepEqual(got, input.want) {
				t.Errorf("GetLatestRates() = %v, want %v", got, input.want)
			}
		})
	}
}
