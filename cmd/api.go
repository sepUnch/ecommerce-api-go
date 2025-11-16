package main

type application struct {
	config config
}

type config struct {
	address string
	db dbConfig

}

type dbConfig struct {
	dsn string
}