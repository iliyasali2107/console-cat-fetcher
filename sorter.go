package main

type breedNameSorter []Breed

func (bns breedNameSorter) Len() int {
	return len(bns)
}

func (bns breedNameSorter) Swap(i, j int) {
	bns[i], bns[j] = bns[j], bns[i]
}

func (bns breedNameSorter) Less(i, j int) bool {
	return len(bns[i].Breed) > len(bns[j].Breed)
}
