package db

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func MysqlInit() (*sqlx.DB, error) {
	loc, err := time.LoadLocation(os.Getenv("MYSQL_TZ"))
	if err != nil {
		return nil, err
	}

	cnf := mysql.Config{
		DBName:    os.Getenv("MYSQL_DATABASE"),
		User:      os.Getenv("MYSQL_USERNAME"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Addr:      os.Getenv("MYSQL_HOSTNAME") + ":" + os.Getenv("MYSQL_PORT"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4",
		Loc:       loc,
	}

	db, err := sqlx.Open("mysql", cnf.FormatDSN())
	if err != nil {
		return nil, err
	}
	fmt.Println("======== DB Connection Succeed ! ========")

	return db, nil
}

// type Store struct {
// 	db *sqlx.DB
// 	*generated.Queries
// }

// func NewStore(db *sqlx.DB) *Store {
// 	return &Store{
// 		db:      db,
// 		Queries: generated.New(db),
// 	}
// }

// func (store *Store) ExecTx(ctx context.Context, fn func(*generated.Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	q := generated.New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
