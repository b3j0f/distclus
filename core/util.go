package core

func GetChunk(chunckNumber int, offset int, elemts []Elemt) []Elemt {
	var start= chunckNumber * offset
	var end= start + offset

	switch {
		case start > len(elemts):
			return []Elemt{}

		case end > len(elemts) :
			return elemts[start:]

		default:
			return elemts[start:end]
	}
}