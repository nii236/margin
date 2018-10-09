package positions

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"syreclabs.com/go/faker"

	_ "github.com/mattn/go-sqlite3"
)

type Type string

const ShortPosition Type = "Short"
const LongPosition Type = "Long"

type Pair string

const XBTUSD Pair = "XBTUSD"

// ListResponse is the response from the API
type ListResponse struct {
	Positions []*Position
}

// OpenResponse is the response from the API
type OpenResponse struct {
}

// CloseResponse is the response from the API
type CloseResponse struct {
}

// Repo abstracts away the database underneath positions
type Repo interface {
	List() ([]*Position, error)
	Create(payload *Position) error
	Read(id int) (*Position, error)
	Update(id int, payload *Position) error
	Delete(id int) error
}

// SQLite is the struct containing the sqlite connection
type SQLite struct {
	*sqlx.DB
}

// NewSQLite returns a new persister
func NewSQLite(filename string, migrate bool, seed bool) (*SQLite, error) {
	DB, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	result := &SQLite{
		DB: DB,
	}

	if migrate {
		result.Migrate()
	}

	if seed {
		result.Seed()
	}

	return result, nil

}

// Migrate will seed the DB (with panic)
func (db *SQLite) Seed() {
	for i := 0; i < 50; i++ {
		pos := &Position{
			User: faker.Internet().UserName(),
			Pair: XBTUSD,
		}

		err := db.Create(pos)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// Migrate will migrate the DB (with panic)
func (db *SQLite) Migrate() {
	schema := `
CREATE TABLE positions (
	user TEXT,
	type TEXT,
	pair TEXT,
	open_value INT,
	avg_value INT,
	leverage INT,
	stack INT,
	pnl REAL,
	duration INT
);`
	db.MustExec(schema)
}

// Position is a single short or long position
type Position struct {
	User      string  `db:"user"`
	Type      string  `db:"type"`
	Pair      Pair    `db:"pair"`
	OpenValue int     `db:"open_value"`
	AvgValue  int     `db:"avg_value"`
	Leverage  int     `db:"leverage"`
	Stack     int     `db:"stack"`
	Pnl       float64 `db:"pnl"`
	Duration  int     `db:"duration"`
}

// List is standard CRUD methods
func (db *SQLite) List() ([]*Position, error) {
	result := []*Position{}
	return result, db.Select(&result, "SELECT * FROM positions;")
}

// Create is standard CRUD methods
func (db *SQLite) Create(payload *Position) error {
	q := `
INSERT INTO positions (
	user,
	type,
	pair,
	open_value,
	avg_value,
	leverage,
	stack,
	pnl,
	duration
) VALUES (
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?
)`
	_, err := db.Exec(q,
		payload.User,
		payload.Type,
		payload.Pair,
		payload.OpenValue,
		payload.AvgValue,
		payload.Leverage,
		payload.Stack,
		payload.Pnl,
		payload.Duration,
	)

	return err
}

// Read is standard CRUD methods
func (db *SQLite) Read(id int) (*Position, error) {
	result := &Position{}
	return result, db.Get(result, "SELECT * FROM positions WHERE ID = ?;")
}

// Update is standard CRUD methods
func (db *SQLite) Update(id int, payload *Position) error {
	q := `
UPDATE positions
SET 
	user = ?,
	type = ?,
	pair = ?,
	open_value = ?,
	avg_value = ?,
	leverage = ?,
	stack = ?,
	pnl = ?,
	duration = ?
WHERE
	id = ?
`

	_, err := db.Exec(q,
		payload.User,
		payload.Type,
		payload.Pair,
		payload.OpenValue,
		payload.AvgValue,
		payload.Leverage,
		payload.Stack,
		payload.Pnl,
		payload.Duration,
		id,
	)
	return err
}

// Delete is standard CRUD methods
func (db *SQLite) Delete(id int) error {
	q := `
DELETE
FROM
	positions
WHERE
	id = ?
`
	_, err := db.Exec(q, id)
	return err
}
