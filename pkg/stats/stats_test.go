package stats

import (
	"testing"
)

func TestNewTable(t *testing.T) {
	headers := []string{"Header1", "Header2"}
	table := NewTable(headers)
	if table == nil {
		t.Errorf("NewTable was incorrect, got: nil, want: not nil.")
	}
	if len(table.headers) != len(headers) {
		t.Errorf("NewTable headers length was incorrect, got: %d, want: %d.", len(table.headers), len(headers))
	}
	for i, header := range table.headers {
		if header != headers[i] {
			t.Errorf("NewTable header was incorrect, got: %s, want: %s.", header, headers[i])
		}
	}
}

func TestTable_AddRow(t *testing.T) {
	table := NewTable([]string{"Header1", "Header2"})
	row := NewRow("Repo1", "Org1", "Project1", true)
	table.AddRow(*row)
	if len(table.rows) != 1 {
		t.Errorf("AddRow was incorrect, got: %d rows, want: 1 row.", len(table.rows))
	}
	if table.rows[0].GitHubRepo != "Repo1" {
		t.Errorf("AddRow GitHubRepo was incorrect, got: %s, want: Repo1.", table.rows[0].GitHubRepo)
	}
}

func TestTable_String(t *testing.T) {
	table := NewTable([]string{"Header1", "Header2"})
	row := NewRow("Repo1", "Org1", "Project1", true)
	table.AddRow(*row)
	expectedString := "| Header1 | Header2 |\n| :---: | :---: |\n| Repo1 | Org1 | Project1 | :white_check_mark: | :white_check_mark: | :ok_hand: |\n"
	if table.String() != expectedString {
		t.Errorf("String was incorrect, got: %s, want: %s.", table.String(), expectedString)
	}
}
