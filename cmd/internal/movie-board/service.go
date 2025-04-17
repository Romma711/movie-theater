package movie_board

import (
	pb "github.com/Romma711/movie_theater/api/proto/movie_board_service"
)

type Service struct {
	movieBoardRepo MovieBoardRepository
	pb.UnimplementedMovieBoardServiceServer
}
