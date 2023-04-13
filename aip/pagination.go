package repository

import (
	"fmt"

	"go.einride.tech/aip/pagination"
)

/**
	Extract page size from request and maps it between min page size and max page size if not in range.
	For more - https://google.aip.dev/158
**/
func ParseAndValidatePageSize(req pagination.Request, minps int32, maxps int32) (int32, error) {
	switch {
	case req.GetPageSize() >= 0 && req.GetPageSize() <= minps:
		return minps, nil
	case req.GetPageSize() > minps && req.GetPageSize() <= maxps:
		return req.GetPageSize(), nil
	case req.GetPageSize() > maxps:
		return maxps, nil
	default:
		return 0, fmt.Errorf("invalid page size : %d", req.GetPageSize())
	}
}

/**
	Extract page token from request and parse it.
	For more - https://google.aip.dev/158
**/
func ParseAndValidatePageToken(req pagination.Request) (pagination.PageToken, error) {
	pt, err := pagination.ParsePageToken(req)
	if err != nil {
		return pagination.PageToken{}, fmt.Errorf("invalid page token : %v", req.GetPageToken())
	}
	return pt, nil
}

/**
	Create next page token from page size and previous page token
	For more - https://google.aip.dev/158
**/
func NextPageToken(req pagination.Request, resultSize int32, minps int32, maxps int32) (string, error) {
	pageSize, err := ParseAndValidatePageSize(req, minps, maxps)
	if err != nil {
		return empty, err
	}
	pageToken, err := ParseAndValidatePageToken(req)
	if err != nil {
		return empty, err
	}
	if resultSize < pageSize {
		return empty, nil
	}
	return pageToken.Next(req).String(), nil
}
