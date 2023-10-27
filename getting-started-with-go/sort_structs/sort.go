package sortstructs

type Person struct {
	id   int
	name string
}

func QuickSort(a *[]Person) {
	quick_sort(a, 0, len(*a))
}

func quick_sort(a *[]Person, start int, end int) {
	for end-start > 1 {
		m := partition(a, start, end)
		if m-start < end-m {
			quick_sort(a, start, m)
			start = m + 1
		} else {
			quick_sort(a, m+1, end)
			end = m
		}
	}
}

func partition(a *[]Person, start int, end int) int {
	pivot := (*a)[start].id
	sx := start
	dx := end

	for sx < dx {

		for ok := true; ok; ok = (*a)[dx].id < pivot {
			dx -= 1
		}

		for ok := true; ok; ok = sx < dx && (*a)[sx].id >= pivot {
			sx += 1
		}

		if sx < dx {
			swap(a, sx, dx)
		}
	}

	swap(a, start, dx)
	return dx
}

func swap(a *[]Person, from int, to int) {
	tmp := (*a)[from]
	(*a)[from] = (*a)[to]
	(*a)[to] = tmp
}
