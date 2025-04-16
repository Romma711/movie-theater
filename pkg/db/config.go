package db

type Config struct {
	Host string,
	Port string,
	User string,
	Password string,
	DBName string,
	SSLMode string
}

func NewConn(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)
	
	log.Println("Conexion a la base de datos realizada correctamente")
	return db, nil
}