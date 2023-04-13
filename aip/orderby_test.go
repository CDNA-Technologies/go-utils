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
			wantErr: fmt.Errorf("invalid orderby : ?"),
		},
		{
			req: &pb.ListPartnersRequest{
				OrderBy: "name desc age asc",
			},
			want:    nil,
			wantErr: fmt.Errorf("invalid orderby : name desc age asc"),
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
			wantErr: fmt.Errorf("invalid orderby : name Desc, age"),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ParseAndValidateOrderBy(%v)", tt.req.GetOrderBy()), func(t *testing.T) {
			got, err := ParseAndValidateOrderBy(tt.req)
			if !testUtils.IsErrorEqual(tt.wantErr, err) {
				t.Errorf("expected %#v but got %#v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAndValidateOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
