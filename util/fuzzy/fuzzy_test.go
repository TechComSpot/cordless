package fuzzy

import (
	"testing"
)

func TestFuzzyScore(t *testing.T) {
	scoreA := Score("Marc", "Marcel#4759")
	scoreB := Score("Marc", "Maurice")
	if scoreA <= scoreB {
		t.Errorf("Incorrect score rating")
	}

	scoreC := Score("Marc", "Marchesi#4377")
	if scoreC <= scoreB {
		t.Errorf("Incorrect score rating")
	}

	arr := []string{"tests", "test", "testosterone", "atesta", "bob"}
	sorted := SortSearchResults(ScoreSearch("te", arr))
	expected := [4]string{"test", "testosterone", "tests", "atesta"}

	if len(sorted) != len(expected) {
		t.Errorf("Expected length of %d, but received %d.\n", len(expected), len(sorted))
	}

	var results [4]string
	for i, result := range sorted {
		results[i] = result.Key
	}

	if results != expected {
		t.Errorf("Expected\n%v\nbut received\n%v\n", expected, results)
	}

}