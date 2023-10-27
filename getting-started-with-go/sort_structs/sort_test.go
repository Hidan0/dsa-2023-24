package sortstructs

import "testing"

func TestSort(t *testing.T) {
	ppl := []Person{{6, "Francesco"}, {1, "Andrea"}, {5, "Elisa"}, {2, "Beatrice"}, {3, "Carlo"}, {4, "Dino"}, {7, "Giorgio"}, {9, "Irene"}, {8, "Henry"}}
	want := []Person{{9, "Irene"}, {8, "Henry"}, {7, "Giorgio"}, {6, "Francesco"}, {5, "Elisa"}, {4, "Dino"}, {3, "Carlo"}, {2, "Beatrice"}, {1, "Andrea"}}

	QuickSort(&ppl)

	for i, p := range ppl {
		if p != want[i] {
			t.Errorf("At index %d I was expecting %v but got %v", i, want[i], p)
		}
	}
}
