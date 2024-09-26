package models

type Snakes struct {
	Start int
	End   int
}

func CreateSnakes(start, end int) Snakes {
	return Snakes{
		Start: start,
		End:   end,
	}
}
