package main

import "time"

type Movie struct {
	ID          int
	Title       string
	Year        int
	Rating      float32
	Genre       string
	Plot        string
	PosterURL   string
	TicketPrice float32
	RelaseDate  time.Time
	Status      string
}

type MovieDTO struct {
	Title       string
	Year        int
	Rating      float32
	Genre       string
	Plot        string
	PosterURL   string
	TicketPrice float32
	RelaseDate  time.Time
	Status      string
}

