package gatherpermutations

// Data una permutazione di 1..N, vogliamo raccogliere i numeri in ordine crescente cominciando
// ad analizzarli da sinistra. Scrivete un programma che stabilisce quante volte avremo bisogno di
// tornare verso sinistra.
// Esempio: Nella permutazione 4 5 1 3 6 2 l’output è 2, poiché 1 si trova andando sempre
// verso destra, poi si prosegue verso destra per raccogliere 2, ma per raccogliere 3 bisogna tornare
// indietro verso sinistra; bisogna tornare ancora indietro per raccogliere 4, dopodiché 5 e 6 si
// trovano in ordine proseguendo verso destra.

func GatherPermutations(a *[]int) int {
	backSteps := 0

	toFind := 1

	for i := 0; toFind <= len(*a); i++ {
		if i >= len(*a) && toFind <= len(*a) {
			i = 0
			backSteps++
		}

		if (*a)[i] == toFind {
			toFind++
		}
	}

	return backSteps
}
