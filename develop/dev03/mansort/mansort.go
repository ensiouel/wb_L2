package mansort

import (
	"sort"
	"strconv"
	"strings"
)

func Sort(lines []string, column int, isNumericColumn bool, isReverse bool, isUnique bool) []string {
	if isUnique {
		lines = set(lines)
	}

	fields := make([][]string, 0, len(lines))

	for _, line := range lines {
		if column == -1 {
			fields = append(fields, []string{line})
			continue
		}

		fields = append(fields, strings.Fields(line))
	}

	column = max(0, column-1)

	var slice = &FieldSlice{
		fields:          fields,
		column:          column,
		isNumericColumn: isNumericColumn,
	}

	if isReverse {
		sort.Sort(sort.Reverse(slice))
	} else {
		sort.Sort(slice)
	}

	lines = make([]string, 0, len(fields))
	for _, field := range slice.fields {
		lines = append(lines, strings.Join(field, " "))
	}

	return lines
}

type FieldSlice struct {
	fields [][]string

	column          int
	isNumericColumn bool
}

func (slice *FieldSlice) Len() int {
	return len(slice.fields)
}

func (slice *FieldSlice) Less(i, j int) bool {
	iLen := len(slice.fields[i])
	jLen := len(slice.fields[j])

	if iLen == 0 || jLen == 0 {
		return iLen < jLen
	}

	iColumn := max(0, min(slice.column, iLen-1))
	jColumn := max(0, min(slice.column, jLen-1))

	if iColumn != jColumn {
		return iColumn < jColumn
	}

	if slice.isNumericColumn {
		var (
			a, b float64
			err  error
		)

		a, err = strconv.ParseFloat(trimNumber(slice.fields[i][iColumn]), 64)
		if err != nil {
			return slice.fields[i][iColumn] < slice.fields[j][jColumn]
		}

		b, err = strconv.ParseFloat(trimNumber(slice.fields[j][jColumn]), 64)
		if err != nil {
			return slice.fields[i][iColumn] < slice.fields[j][jColumn]
		}

		return a < b
	}

	return slice.fields[i][iColumn] < slice.fields[j][jColumn]
}

func (slice *FieldSlice) Swap(i, j int) {
	slice.fields[i], slice.fields[j] = slice.fields[j], slice.fields[i]
}
