package main

// TravelAgenda ...
type TravelAgenda struct  {
	flights map[int]*Flight
}

// NewTravelAgenda ...
func NewTravelAgenda() TravelAgenda {
	flights := make(map[int]*Flight)
	return TravelAgenda {
		flights,
	}
}

// Add ...
func (ta TravelAgenda) Add(f Flight)  {
	ta.flights[f.ID] = &f
}

// Print ...
func (ta TravelAgenda) Print() {
	for _, v := range ta.flights {
		v.Print()
	}
}