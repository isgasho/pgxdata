package data

// This file is automatically generated by pgxdata.

import (
	"context"
	"strings"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	errors "golang.org/x/xerrors"
)

type Blob struct {
	ID      pgtype.Int4
	Payload pgtype.Bytea
}

const countBlobSQL = `select count(*) from "blob"`

func CountBlob(ctx context.Context, db Queryer) (int64, error) {
	var n int64
	err := prepareQueryRow(ctx, db, "pgxdataCountBlob", countBlobSQL).Scan(&n)
	return n, err
}

const SelectAllBlobSQL = `select
  "id",
  "payload"
from "blob"`

func SelectAllBlob(ctx context.Context, db Queryer) ([]Blob, error) {
	var rows []Blob

	dbRows, err := prepareQuery(ctx, db, "pgxdataSelectAllBlob", SelectAllBlobSQL)
	if err != nil {
		return nil, err
	}

	for dbRows.Next() {
		var row Blob
		dbRows.Scan(
			&row.ID,
			&row.Payload,
		)
		rows = append(rows, row)
	}

	if dbRows.Err() != nil {
		return nil, dbRows.Err()
	}

	return rows, nil
}

const selectBlobByPKSQL = `select
  "id",
  "payload"
from "blob"
where "id"=$1`

func SelectBlobByPK(
	ctx context.Context,
	db Queryer,
	id int32,
) (*Blob, error) {
	var row Blob
	err := prepareQueryRow(ctx, db, "pgxdataSelectBlobByPK", selectBlobByPKSQL, id).Scan(
		&row.ID,
		&row.Payload,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &row, nil
}

func InsertBlob(ctx context.Context, db Queryer, row *Blob) error {
	args := pgx.QueryArgs(make([]interface{}, 0, 2))

	var columns, values []string

	if row.ID.Status != pgtype.Undefined {
		columns = append(columns, `id`)
		values = append(values, args.Append(&row.ID))
	}
	if row.Payload.Status != pgtype.Undefined {
		columns = append(columns, `payload`)
		values = append(values, args.Append(&row.Payload))
	}

	sql := `insert into "blob"(` + strings.Join(columns, ", ") + `)
values(` + strings.Join(values, ",") + `)
returning "id"
  `

	psName := preparedName("pgxdataInsertBlob", sql)

	return prepareQueryRow(ctx, db, psName, sql, args...).Scan(&row.ID)
}

func UpdateBlob(ctx context.Context, db Queryer,
	id int32,
	row *Blob,
) error {
	sets := make([]string, 0, 2)
	args := pgx.QueryArgs(make([]interface{}, 0, 2))

	if row.ID.Status != pgtype.Undefined {
		sets = append(sets, `id`+"="+args.Append(&row.ID))
	}
	if row.Payload.Status != pgtype.Undefined {
		sets = append(sets, `payload`+"="+args.Append(&row.Payload))
	}

	if len(sets) == 0 {
		return nil
	}

	sql := `update "blob" set ` + strings.Join(sets, ", ") + ` where ` + `"id"=` + args.Append(id)

	psName := preparedName("pgxdataUpdateBlob", sql)

	commandTag, err := prepareExec(ctx, db, psName, sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}

func DeleteBlob(ctx context.Context, db Queryer,
	id int32,
) error {
	args := pgx.QueryArgs(make([]interface{}, 0, 1))

	sql := `delete from "blob" where ` + `"id"=` + args.Append(id)

	commandTag, err := prepareExec(ctx, db, "pgxdataDeleteBlob", sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}
