package main

func main()  {
	ta := NewTravelAgenda()
	ta.Add(Flight{0, "Iberia", "IB1234", "2020-12-28T20:00Z", "2021-12-29T08:00Z", "EZE", "MAD"})
	ta.Add(Flight{1, "Iberia", "IB5678", "2021-12-29T15:00Z", "2021-12-29T17:00Z", "MAD", "LIS"})

	ta.Print()
}