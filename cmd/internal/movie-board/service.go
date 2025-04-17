package main

import (
	pb "github.com/Romma711/movie_theater/api/proto/movie_board_service"
)

type Service struct {
	movieBoardRepo MovieBoardRepository
}
