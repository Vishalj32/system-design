package models

type Ladders struct {
	Start int
	End   int
}

func CreateLadder(start, end int) Ladders {
	return Ladders{
		Start: start,
		End:   end,
	}
}
