// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getCommentsByProduct = `-- name: GetCommentsByProduct :many
SELECT id, user_id, tx, ts
FROM comments
WHERE product_id = $1
`

type GetCommentsByProductRow struct {
	ID     int64
	UserID int64
	Tx     string
	Ts     pgtype.Timestamp
}

func (q *Queries) GetCommentsByProduct(ctx context.Context, productID int64) ([]*GetCommentsByProductRow, error) {
	rows, err := q.db.Query(ctx, getCommentsByProduct, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCommentsByProductRow
	for rows.Next() {
		var i GetCommentsByProductRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Tx,
			&i.Ts,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnSendNotification = `-- name: GetUnSendNotification :many
SELECT id, owner_id, comment_id, ts, status
FROM outbox_notification
WHERE status = 'new'
ORDER BY ts
    LIMIT $1
`

func (q *Queries) GetUnSendNotification(ctx context.Context, limit int32) ([]*OutboxNotification, error) {
	rows, err := q.db.Query(ctx, getUnSendNotification, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*OutboxNotification
	for rows.Next() {
		var i OutboxNotification
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.CommentID,
			&i.Ts,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const maskNotificationAsSend = `-- name: MaskNotificationAsSend :exec
UPDATE outbox_notification
SET status = 'send'
WHERE id = $1
`

func (q *Queries) MaskNotificationAsSend(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, maskNotificationAsSend, id)
	return err
}

const saveComment = `-- name: SaveComment :one
INSERT INTO comments (user_id, product_id, tx, ts)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type SaveCommentParams struct {
	UserID    int64
	ProductID int64
	Tx        string
	Ts        pgtype.Timestamp
}

func (q *Queries) SaveComment(ctx context.Context, arg *SaveCommentParams) (int64, error) {
	row := q.db.QueryRow(ctx, saveComment,
		arg.UserID,
		arg.ProductID,
		arg.Tx,
		arg.Ts,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const saveNotification = `-- name: SaveNotification :exec
INSERT INTO outbox_notification (owner_id, comment_id, ts)
VALUES ($1, $2, $3)
`

type SaveNotificationParams struct {
	OwnerID   int64
	CommentID int64
	Ts        pgtype.Timestamp
}

func (q *Queries) SaveNotification(ctx context.Context, arg *SaveNotificationParams) error {
	_, err := q.db.Exec(ctx, saveNotification, arg.OwnerID, arg.CommentID, arg.Ts)
	return err
}
