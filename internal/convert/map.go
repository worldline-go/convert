package convert

import "strings"

type Map struct {
	Column map[int]string
}

func (m *Map) Parse(values []string) map[string]string {
	result := make(map[string]string, len(m.Column))
	for i, v := range values {
		if columnName, ok := m.Column[i]; ok {
			result[columnName] = v
		}
	}

	return result
}

func NewMap(values map[string]string) (*Map, error) {
	v := &Map{
		Column: make(map[int]string),
	}

	for column, value := range values {
		// A -> 0, B -> 1, C -> 2, ...
		// AA -> 26, AB -> 27, ...
		v.Column[ConvertColumnToNumber(column)] = value
	}

	return v, nil
}

func ConvertColumnToNumber(column string) int {
	column = strings.ToUpper(column)
	sum := 0
	for i := 0; i < len(column); i++ {
		sum *= 26
		sum += int(column[i] - 'A' + 1)
	}

	return sum - 1
}
