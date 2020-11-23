package main

import "fmt"

// Playlist ...
type Playlist struct {
	songs map[int]*Song
}

// Song ...
type Song struct {
	ID int
	Name string
}

// NewPlaylist ...
func NewPlaylist() Playlist {
	songs := make(map[int]*Song)
	return Playlist {
		songs,
	}
}

// Add ...
func (p Playlist) Add(s Song)  {
	p.songs[s.ID] = &s
}

// Print ...
func (p Playlist) Print()  {
	for _, v := range p.songs {
		fmt.Printf("[%v]\t%v\n", v.ID, v.Name)
	}
}

// FindByID ...
func (p Playlist) FindByID(ID int) *Song {
	return p.songs[ID]
}

// Delete ...
func (p Playlist) Delete(ID int) {
	delete(p.songs, ID)
}

// Update ...
func (p Playlist) Update(s Song)  {
	p.songs[s.ID] = &s
}