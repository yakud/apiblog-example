package pg

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/pkg/errors"
)

func NewConnection(opt *pg.Options) (*pg.DB, error) {
	if opt == nil {
		opt = &pg.Options{
			User:     "pgadmin",
			Password: "pgadmin",
			Database: "apiblog",
			Addr:     "127.0.0.1:5432",
		}
	}

	db := pg.Connect(opt)

	ok, err := PingDB(db)
	if err != nil {
		return nil, fmt.Errorf("pgsql new connection: %s", err.Error())
	}

	if !ok {
		return nil, errors.New("can't ping pgdb, just is not ok")
	}

	return db, nil
}

func PingDB(db *pg.DB) (bool, error) {
	res, err := db.ExecOne("SELECT 1")
	if err != nil {
		return false, fmt.Errorf("pgsql ping: %s", err.Error())
	}

	if res.RowsReturned() == 1 {
		return true, nil
	}

	return false, nil
}
