package convert

import (
	"fmt"
	"strconv"
	"strings"
)

type Rows struct {
	Ranges []Range
}

type Range struct {
	Start int
	End   int
}

// NewRows creates a new Rows which can check the row number.
//
// Values could be '30', '29-55', '1-10,20-30', '5,1-10,20-30,30-30'.
func NewRows(values []string) (*Rows, error) {
	rows := &Rows{}

	for _, v := range values {
		fileds := strings.Fields(strings.ReplaceAll(v, ",", " "))
		for _, filed := range fileds {
			if strings.Contains(filed, "-") {
				r := strings.Split(filed, "-")
				if len(r) != 2 {
					return nil, fmt.Errorf("invalid range: %s", filed)
				}

				start, err := strconv.Atoi(r[0])
				if err != nil {
					return nil, fmt.Errorf("invalid number: %s", r[0])
				}

				end, err := strconv.Atoi(r[1])
				if err != nil {
					return nil, fmt.Errorf("invalid number: %s", r[1])
				}

				rows.Ranges = append(rows.Ranges, Range{Start: start, End: end})

				continue
			}

			v, err := strconv.Atoi(filed)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", filed)
			}

			rows.Ranges = append(rows.Ranges, Range{Start: v, End: v})
		}
	}

	return rows, nil
}

func (r *Rows) IsInclude(v int) bool {
	for _, r := range r.Ranges {
		if v >= r.Start && v <= r.End {
			return true
		}
	}

	return false
}
