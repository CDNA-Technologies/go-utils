package aip

import (
	"fmt"
	"reflect"
	"testing"

	testUtils "github.com/CDNA-Technologies/go-utils/testing"
	"go.einride.tech/aip/pagination"
	pb "google.golang.org/genproto/googleapis/example/library/v1"
)

func TestParseAndValidatePageSize(t *testing.T) {
	tests := []struct {
		req         pagination.Request
		minPageSize int32
		maxPageSize int32
		want        int32
		wantErr     error
	}{
		{
			req: &pb.ListBooksRequest{
				PageSize:  1,
				PageToken: "",
			},
			minPageSize: 5,
			maxPageSize: 100,
			want:        5,
			wantErr:     nil,
		},
		{
			req: &pb.ListBooksRequest{
				PageSize:  500,
				PageToken: "",
			},
			minPageSize: 5,
			maxPageSize: 100,
			want:        100,
			wantErr:     nil,
		},
		{
			req: &pb.ListBooksRequest{
				PageSize:  50,
				PageToken: "",
			},
			minPageSize: 5,
			maxPageSize: 100,
			want:        50,
			wantErr:     nil,
		},
		{
			req: &pb.ListBooksRequest{
				PageSize:  -5,
				PageToken: "",
			},
			minPageSize: 5,
			maxPageSize: 100,
			want:        0,
			wantErr:     fmt.Errorf("invalid page size : -5"),
		},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestParseAndValidatePageSize(%v, %d, %d)", input.req.GetPageSize(), input.minPageSize, input.maxPageSize), func(t *testing.T) {
			got, err := ParseAndValidatePageSize(input.req, input.minPageSize, input.maxPageSize)
			if !testUtils.IsErrorEqual(input.wantErr, err) {
				t.Errorf("ParseAndValidatePageSize(%v, %d, %d) got error =  %#v, wantErr =  %#v",
					input.req.GetPageSize(), input.minPageSize, input.maxPageSize, err, input.wantErr)
			}
			if got != input.want {
				t.Errorf("ParseAndValidatePageSize(%v, %d, %d) =  %#v, want =  %#v",
					input.req.GetPageSize(), input.minPageSize, input.maxPageSize, got, input.want)
			}
		})
	}
}

func TestNextPageToken(t *testing.T) {
	tests := []struct {
		req         pagination.Request
		resultSize  int32
		minPageSize int32
		maxPageSize int32
		want        string
		wantErr     error
	}{
		{
			req: &pb.ListBooksRequest{
				PageSize:  10,
				PageToken: "",
			},
			resultSize:  5,
			minPageSize: 5,
			maxPageSize: 100,
			want:        "",
			wantErr:     nil,
		},
		{
			req: &pb.ListBooksRequest{
				PageSize:  10,
				PageToken: "",
			},
			resultSize:  15,
			minPageSize: 5,
			maxPageSize: 100,
			want:        "NX8DAQEJUGFnZVRva2VuAf-AAAECAQZPZmZzZXQBBAABD1JlcXVlc3RDaGVja3N1bQEGAAAAC_-AARQB_JrLBEIA",
			wantErr:     nil,
		},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestNextPageToken(%v, %d, %d, %d)", input.req, input.resultSize, input.minPageSize, input.maxPageSize), func(t *testing.T) {
			got, err := NextPageToken(input.req, input.resultSize, input.minPageSize, input.maxPageSize)
			if !testUtils.IsErrorEqual(input.wantErr, err) {
				t.Errorf("NextPageToken(%v, %d, %d, %d) got error = %#v, wantErr %#v",
					input.req, input.resultSize, input.minPageSize, input.maxPageSize, err, input.wantErr)
			}
			if got != input.want {
				t.Errorf("NextPageToken(%v, %d, %d, %d) = %#v, want %#v",
					input.req, input.resultSize, input.minPageSize, input.maxPageSize, got, input.want)
			}
		})
	}
}

func TestParseAndValidatePageToken(t *testing.T) {
	tests := []struct {
		req     pagination.Request
		want    pagination.PageToken
		wantErr error
	}{
		{
			req: &pb.ListBooksRequest{
				PageToken: "Random String",
				PageSize:  10,
			},
			want:    pagination.PageToken{},
			wantErr: fmt.Errorf("parse offset page token: decode page token struct: illegal base64 data at input byte 6"),
		},
		{
			req: &pb.ListBooksRequest{
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
			req: &pb.ListBooksRequest{
				PageToken: "NX8DAQEJUGFnZVRva2VuAf-AAAECAQZPZmZzZXQBBAABD1JlcXVlc3RDaGVja3N1bQEGAAAAC_-AARQB_JrLBEIA",
				PageSize:  10,
			},
			want: pagination.PageToken{
				Offset:          10,
				RequestChecksum: 2596996162,
			},
			wantErr: nil,
		},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestParseAndValidatePageToken(%v)", input.req), func(t *testing.T) {
			got, err := ParseAndValidatePageToken(input.req)
			if !testUtils.IsErrorEqual(input.wantErr, err) {
				t.Errorf("ParseAndValidatePageToken(%v) got error = %#v, wantErr %#v",
					input.req, err, input.wantErr)
			}
			if !reflect.DeepEqual(got, input.want) {
				t.Errorf("ParseAndValidatePageToken(%v) = %#v, want %#v", input.req, got, input.want)
			}
		})
	}
}
