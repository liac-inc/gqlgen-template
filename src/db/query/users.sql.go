// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package query

import (
	"context"
)

const findAllUsers = `-- name: FindAllUsers :many
SELECT
    id, name
FROM
    users
`

func (q *Queries) FindAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}