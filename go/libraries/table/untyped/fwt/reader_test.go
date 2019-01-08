package fwt

import (
	"github.com/liquidata-inc/ld/dolt/go/libraries/filesys"
	"github.com/liquidata-inc/ld/dolt/go/libraries/table"
	"github.com/liquidata-inc/ld/dolt/go/libraries/table/untyped"
	"io"
	"strings"
	"testing"
)

var PersonDB1 = `Bill Billerson32Senior Dufus
Rob Robertson 25Dufus       
John Johnson  21Intern Dufus`

var PersonDB2 = PersonDB1 + "\n"
var PersonDB3 = strings.Replace(PersonDB2, "\n", "\n\n", -1)

var PersonDBWithBadRow = `Bill Billerson | 32 | Senior Dufus
Rob Robertson  | 25 | Dufus       
John Johnson   | 21 |`

var PersonDBWithBadRow2 = PersonDBWithBadRow + "\n"
var PersonDBWithBadRow3 = strings.Replace(PersonDBWithBadRow2, "\n", "\n\n", -1)

func TestReader(t *testing.T) {
	colNames := []string{"name", "age", "title"}
	sch := untyped.NewUntypedSchema(colNames)

	goodExpectedRows := []*table.Row{
		untyped.NewRowFromStrings(sch, []string{"Bill Billerson", "32", "Senior Dufus"}),
		untyped.NewRowFromStrings(sch, []string{"Rob Robertson", "25", "Dufus"}),
		untyped.NewRowFromStrings(sch, []string{"John Johnson", "21", "Intern Dufus"}),
	}
	badExpectedRows := []*table.Row{
		untyped.NewRowFromStrings(sch, []string{"Bill Billerson", "32", "Senior Dufus"}),
		untyped.NewRowFromStrings(sch, []string{"Rob Robertson", "25", "Dufus"}),
	}

	widths := map[string]int{
		colNames[0]: 14,
		colNames[1]: 2,
		colNames[2]: 12,
	}

	fwtSch, _ := NewFWTSchema(sch, widths)

	i := []struct {
		inputStr     string
		expectedRows []*table.Row
		sep          string
	}{
		{PersonDB1, goodExpectedRows, ""},
		{PersonDB2, goodExpectedRows, ""},
		{PersonDB3, goodExpectedRows, ""},

		{PersonDBWithBadRow, badExpectedRows, " | "},
		{PersonDBWithBadRow2, badExpectedRows, " | "},
		{PersonDBWithBadRow3, badExpectedRows, " | "},
	}
	tests := i

	for _, test := range tests {
		rows, numBad, err := readTestRows(t, test.inputStr, fwtSch, test.sep)

		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		expectedBad := len(goodExpectedRows) - len(test.expectedRows)
		if numBad != expectedBad {
			t.Error("Unexpected bad rows count. expected:", expectedBad, "actual:", numBad)
		}

		if !rows[0].GetSchema().Equals(test.expectedRows[0].GetSchema()) {
			t.Fatal("Unexpected schema")
		} else if len(rows) != len(test.expectedRows) {
			t.Error("Did not receive the correct number of rows. expected: ", len(test.expectedRows), "actual:", len(rows))
		} else {
			for i, row := range rows {
				expectedRow := test.expectedRows[i]
				if !table.RowsEqualIgnoringSchema(row, expectedRow) {
					t.Error(table.RowFmt(row), "!=", table.RowFmt(expectedRow))
				}
			}
		}
	}
}

func readTestRows(t *testing.T, inputStr string, fwtSch *FWTSchema, sep string) ([]*table.Row, int, error) {
	const root = "/"
	const path = "/file.csv"

	fs := filesys.NewInMemFS(nil, map[string][]byte{path: []byte(inputStr)}, root)
	fwtRd, err := OpenFWTReader(path, fs, fwtSch, sep)
	defer fwtRd.Close()

	if err != nil {
		t.Fatal("Could not open reader", err)
	}

	badRows := 0
	var rows []*table.Row
	for {
		row, err := fwtRd.ReadRow()

		if err != io.EOF && err != nil && err != table.ErrBadRow {
			return nil, -1, err
		} else if err == table.ErrBadRow {
			badRows++
			continue
		} else if err == io.EOF && row == nil {
			break
		}

		rows = append(rows, row)
	}

	return rows, badRows, err
}