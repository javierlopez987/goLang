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

// FindByID ...
func (ta TravelAgenda) FindByID(ID int) (f *Flight)  {
	return ta.flights[ID]
}

// Delete ...
func (ta TravelAgenda) Delete(ID int)  {
	delete(ta.flights, ID)
}

// Print ...
func (ta TravelAgenda) Print() {
	for _, v := range ta.flights {
		v.Print()
	}
}