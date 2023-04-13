package repository

import (
	"fmt"
	"reflect"
	"testing"

	testUtils "github.com/CDNA-Technologies/go-utils/testing"
	pb "github.com/CDNA-Technologies/proto-gen/go/gonuclei/masterdata/v2"
	"go.einride.tech/aip/pagination"
)

func TestParsePageSize(t *testing.T) {
	tests := []struct {
		req     pagination.Request
		minps   int32
		maxps   int32
		want    int32
		wantErr error
	}{
		{
			req: &pb.ListPartnersRequest{
				PageSize:  1,
				PageToken: "",
			},
			minps:   5,
			maxps:   100,
			want:    5,
			wantErr: nil,
		},
		{
			req: &pb.ListPartnersRequest{
				PageSize:  500,
				PageToken: "",
			},
			minps:   5,
			maxps:   100,
			want:    100,
			wantErr: nil,
		},
		{
			req: &pb.ListPartnersRequest{
				PageSize:  50,
				PageToken: "",
			},
			minps:   5,
			maxps:   100,
			want:    50,
			wantErr: nil,
		},
		{
			req: &pb.ListPartnersRequest{
				PageSize:  -5,
				PageToken: "",
			},
			minps:   5,
			maxps:   100,
			want:    0,
			wantErr: fmt.Errorf("invalid page size : -5"),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ParseAndValidatePageSize(%v, %v, %v)", tt.req.GetPageSize(), tt.minps, tt.maxps), func(t *testing.T) {
			got, err := ParsePageSize(tt.req, tt.minps, tt.maxps)
			if !testUtils.IsErrorEqual(tt.wantErr, err) {
				t.Errorf("expected %#v but got %#v", tt.wantErr, err)
			}
			if got != tt.want {
				t.Errorf("ParseAndValidatePageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextPageToken(t *testing.T) {
	tests := []struct {
		req        pagination.Request
		resultSize int32
		minps      int32
		maxps      int32
		want       string
		wantErr    error
	}{
		{
			req: &pb.ListPartnersRequest{
				PageSize:  10,
				PageToken: "",
			},
			resultSize: 5,
			minps:      5,
			maxps:      100,
			want:       "",
			wantErr:    nil,
		},
		{
			req: &pb.ListPartnersRequest{
				PageSize:  10,
				PageToken: "",
			},
			resultSize: 15,
			minps:      5,
			maxps:      100,
			want:       "Nv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEGT2Zmc2V0AQQAAQ9SZXF1ZXN0Q2hlY2tzdW0BBgAAAAv_ggEUAfyaywRCAA==",
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("NextPageToken(%v, %v, %v, %v)", tt.req, tt.resultSize, tt.minps, tt.maxps), func(t *testing.T) {
			got, err := NextPageToken(tt.req, tt.resultSize, tt.minps, tt.maxps)
			if !testUtils.IsErrorEqual(tt.wantErr, err) {
				t.Errorf("expected %#v but got %#v", tt.wantErr, err)
			}
			if got != tt.want {
				t.Errorf("NextPageToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePageToken(t *testing.T) {
	tests := []struct {
		req     pagination.Request
		want    pagination.PageToken
		wantErr error
	}{
		{
			req: &pb.ListPartnersRequest{
				PageToken: "Random String",
				PageSize:  10,
			},
			want:    pagination.PageToken{},
			wantErr: fmt.Errorf("invalid page token : Random String"),
		},
		{
			req: &pb.ListPartnersRequest{
				PageToken: "",
				PageSize:  10,
			},
			want: pagination.PageToken{
				Offset:          0,
				RequestChecksum: 2596996162,
			},
			wantErr: nil,
		},
		{
			req: &pb.ListPartnersRequest{
				PageToken: "Nv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEGT2Zmc2V0AQQAAQ9SZXF1ZXN0Q2hlY2tzdW0BBgAAAAv_ggEUAfyaywRCAA==",
				PageSize:  10,
			},
			want: pagination.PageToken{
				Offset:          10,
				RequestChecksum: 2596996162,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ParsePageToken(%v)", tt.req), func(t *testing.T) {
			got, err := ParsePageToken(tt.req)
			if !testUtils.IsErrorEqual(tt.wantErr, err) {
				t.Errorf("ParsePageToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePageToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
