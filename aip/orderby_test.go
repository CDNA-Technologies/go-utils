package aip

import (
	"fmt"
	"reflect"
	"testing"

	testUtils "github.com/CDNA-Technologies/go-utils/testing"
	pb "github.com/CDNA-Technologies/proto-gen/go/gonuclei/masterdata/v2"
	"go.einride.tech/aip/ordering"
)

func TestParseAndValidateOrderBy(t *testing.T) {
	tests := []struct {
		req     ordering.Request
		want    map[string]string
		wantErr error
	}{
		{
			req: &pb.ListPartnersRequest{
				OrderBy: "",
			},
			want:    make(map[string]string),
			wantErr: nil,
		},
		{
			req: &pb.ListPartnersRequest{
				OrderBy: "?",
			},
			want:    nil,
			wantErr: fmt.Errorf("unmarshal order by '?': invalid character '?'"),
		},
		{
			req: &pb.ListPartnersRequest{
				OrderBy: "name desc age asc",
			},
			want:    nil,
			wantErr: fmt.Errorf("unmarshal order by 'name desc age asc': invalid format"),
		},
		{
			req: &pb.ListPartnersRequest{
				OrderBy: "name desc",
			},
			want: map[string]string{
				"name": "DESC",
			},
			wantErr: nil,
		},
		{
			req: &pb.ListPartnersRequest{
				OrderBy: "name desc, age",
			},
			want: map[string]string{
				"name": "DESC",
				"age":  "",
			},
			wantErr: nil,
		},
		{
			req: &pb.ListPartnersRequest{
				OrderBy: "name Desc, age",
			},
			want:    nil,
			wantErr: fmt.Errorf("unmarshal order by 'name Desc, age': invalid format"),
		},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestParseAndValidateOrderBy(%v)", input.req), func(t *testing.T) {
			got, err := ParseAndValidateOrderBy(input.req)
			if !testUtils.IsErrorEqual(input.wantErr, err) {
				t.Errorf("ParseAndValidateOrderBy(%v) got error = %#v, wantErr %#v",
					input.req, err, input.wantErr)
			}
			if !reflect.DeepEqual(got, input.want) {
				t.Errorf("ParseAndValidateOrderBy(%v) = %#v, want %#v",
					input.req, got, input.want)
			}
		})
	}
}
