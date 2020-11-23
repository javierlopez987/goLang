package main

import "fmt"

// Flight ...
type Flight struct {
	ID int
	AirlineName string
	FlightNumber string
	DepartureDateTime string
	ArrivalDateTime string
	DepartureAirport string
	ArrivalAirport string
}

func (f Flight) Print()  {
	fmt.Printf("[%v]\t%v\t%v\t%v\t%v\n", f.ID, f.FlightNumber, f.DepartureAirport, f.ArrivalAirport, f.DepartureDateTime)
}