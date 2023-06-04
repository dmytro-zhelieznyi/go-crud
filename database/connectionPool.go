package database

import (
	"database/sql"
	"sync"
)

const (
	maxConnections = 10
)

type ConnectionPool struct {
	Connections chan *sql.DB
}

var (
	pool *ConnectionPool
	once sync.Once
)

func GetConnectionPool(connectionString string) (*ConnectionPool, error) {
	once.Do(func() {
		pool = &ConnectionPool{
			Connections: make(chan *sql.DB, maxConnections),
		}

		for i := 0; i < maxConnections; i++ {
			db, err := sql.Open("postgres", connectionString)
			if err != nil {
				panic(err)
			}

			err = db.Ping()
			if err != nil {
				panic(err)
			}

			pool.Connections <- db
		}
	})

	return pool, nil
}

func (p *ConnectionPool) GetConnection() *sql.DB {
	return <-p.Connections
}

func (p *ConnectionPool) ReleaseConnection(db *sql.DB) {
	p.Connections <- db
}
