package main

import (
	pb "github.com/Romma711/movie_threater/api/proto/movie_board"
	"context"
	"github.com/Romma711/movie_threater/internal/movie-board/types"
)

type Service struct {
	movieBoardRepo MovieBoardRepository
}
