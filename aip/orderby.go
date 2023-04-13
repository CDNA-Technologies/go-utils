package aip

import (
    "fmt"

    "go.einride.tech/aip/ordering"
)

const (
    desc  = "DESC"
    empty = ""
)

/**
    Extract order by from request and parse it and converts into a map of field and order.
    Eg: Name desc
        Age
        Salary desc
    For more - https://google.aip.dev/132
**/
func ParseAndValidateOrderBy(req ordering.Request) (map[string]string, error) {
    o, err := ordering.ParseOrderBy(req)
    if err != nil {
        return nil, fmt.Errorf("invalid orderby : %v", req.GetOrderBy())
    }

    m := make(map[string]string, len(o.Fields))
    for _, v := range o.Fields {
        if v.Desc {
            m[v.Path] = desc
        } else {
            m[v.Path] = empty
        }
    }
    return m, nil
}
