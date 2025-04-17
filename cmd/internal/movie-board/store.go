package movie_board

import (
	"database/sql"
)

type StoreDB struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *StoreDB {
	return &StoreDB{db: db}
}

func (db *StoreDB) GetMovies() ([]Movie, error) {
	rows, err := db.db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies, err := scanMovieRows(rows)
	if err != nil{
		return nil, err
	}

	return movies, nil
}

func (db *StoreDB) GetMovieById (id int) (Movie ,error){
	row, err := db.db.Query("SELECT * FROM movies WHERE id = ?", id)
	if err != nil {
		return Movie{}, err
	}
	movie, err := scanMovieRows(row)
	if err != nil{
		return Movie{}, err
	}
	return movie[0], nil
}

func (db *StoreDB) InsertMovie (movie Movie) error {
	_, err := db.db.Exec("INSERT INTO movies (title, year, genre, rating, poster_url, plot, relase_date, status, ticket_price) VALUES (?,?,?,?,?,?,?,?,?)", movie.Title, movie.Year, movie.Genre, movie.Rating, movie.PosterURL, movie.Plot, movie.RelaseDate, movie.Status, movie.TicketPrice)
	if err != nil {
		return err
	}
	return nil
}

func (db *StoreDB) UpdateMovie (movie Movie) error {
	_, err := db.db.Exec("UPDATE movies SET title=?, year=?, genre=?, rating=?, poster_url=?, plot=?, relase_date=?, status=?, ticket_price=? WHERE id=?", movie.Title, movie.Year, movie.Genre, movie.Rating, movie.PosterURL, movie.Plot, movie.RelaseDate, movie.Status, movie.TicketPrice, movie.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *StoreDB) DeleteMovie (id int) error{
	_, err := db.db.Exec("DELETE FROM movies WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func scanMovieRows(rows *sql.Rows) ([]Movie ,error){
	movies := make([]Movie, 0)
	for rows.Next() {
		var m Movie
		err := rows.Scan(&m.ID, &m.Title, &m.Year, &m.Genre, &m.Rating, &m.PosterURL, &m.Plot, &m.RelaseDate, &m.Status, &m.TicketPrice)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil

}