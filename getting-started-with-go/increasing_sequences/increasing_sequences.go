package increasingsequences

import "errors"

// Data una lunga sequenza di N interi distinti che rappresentano le altezze di una catena montuosa,
// stampare il numero di salite che vanno da sinistra a destra (una salita è una sequenza crescente
// di 2 o più interi, che partono da un punto di minimo e arrivano in un punto di massimo).
// Esempio: nella sequenza 9 1 3 5 2 0 8 6 ci sono due salite: 1 3 5 e 0 8 (1 3 e 3 5 non
// sono salite perché la prima non finisce in un punto di massimo e la seconda non inizia in un
// punto di minimo).
func IncreasingSequences(seq *[]int) ([][]int, error) {
	if len(*seq) < 2 {
		return nil, errors.New("Sequence too short")
	}

	var outSeq [][]int = nil

	var buffer []int = nil

	for _, n := range *seq {

		if buffer == nil {
			buffer = []int{n}
			continue
		}

		if n > buffer[len(buffer)-1] {
			buffer = append(buffer, n)
		} else {
			if len(buffer) >= 2 {
				outSeq = append(outSeq, buffer)
			}
			buffer = []int{n}
		}

	}

	return outSeq, nil
}
