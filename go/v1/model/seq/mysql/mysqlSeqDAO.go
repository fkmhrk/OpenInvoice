package mysql

import (
	"database/sql"

	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/fkmhrk/OpenInvoice/v1/model/seq"
)

type seqDAO struct {
	connection *db.Connection
}

// New creates instance
func New(connection *db.Connection) seq.DAO {
	return &seqDAO{
		connection: connection,
	}
}

func (d *seqDAO) Create(seqType seq.SeqType, year, value int) (seq.Seq, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return seq.Seq{}, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO seq(seq_type,year,value,created_time,modified_time,deleted) VALUES(?,?,?,unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return seq.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year, value)
	if err != nil {
		return seq.Seq{}, err
	}

	tr.Commit()

	return seq.Seq{
		SeqType: seqType,
		Year:    year,
		Value:   value,
	}, nil
}

func (d *seqDAO) Get(seqType seq.SeqType, year int) (seq.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT seq_type,year,value FROM seq WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return seq.Seq{}, err
	}
	defer st.Close()

	rows, err := st.Query(seqType, year)
	if err != nil {
		return seq.Seq{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return seq.Seq{}, nil
	}

	return d.scan(rows), nil
}

func (d *seqDAO) Update(seqType seq.SeqType, year, value int) (seq.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE seq SET seq_type=?, year=?, value=?,modified_time=unix_timestamp(now()) WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return seq.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year, value, seqType, year)
	if err != nil {
		return seq.Seq{}, err
	}

	return seq.Seq{
		SeqType: seqType,
		Year:    year,
		Value:   value,
	}, nil
}

func (d *seqDAO) Next(seqType seq.SeqType, year int) (seq.Seq, error) {
	tx, err := d.connection.Begin()
	if err != nil {
		return seq.Seq{}, err
	}
	defer tx.Rollback()

	seqEntity, err := d.doNext(tx, seqType, year)
	if err == nil {
		tx.Commit()
		return seqEntity, nil
	} else {
		return seq.Seq{}, err
	}
}

func (d *seqDAO) doNext(tx *sql.Tx, seqType seq.SeqType, year int) (seq.Seq, error) {
	// update
	st, err := tx.Prepare("UPDATE seq SET value=value+1,modified_time=unix_timestamp(now()) WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return seq.Seq{}, err
	}
	defer st.Close()

	result, err := st.Exec(seqType, year)
	if err != nil {
		return seq.Seq{}, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		return seq.Seq{}, err
	}
	if rowCount == 0 {
		// create
		st2, err := tx.Prepare("INSERT INTO seq(seq_type,year,value,created_time,modified_time,deleted) VALUES(?,?,?,unix_timestamp(now()),unix_timestamp(now()),0)")
		if err != nil {
			return seq.Seq{}, err
		}
		defer st2.Close()

		_, err = st2.Exec(seqType, year, 1)
		if err != nil {
			return seq.Seq{}, err
		}
		return seq.Seq{
			SeqType: seqType,
			Year:    year,
			Value:   1,
		}, nil
	}

	// get
	st2, err := tx.Prepare("SELECT seq_type,year,value FROM seq WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return seq.Seq{}, err
	}
	defer st2.Close()

	rows, err := st2.Query(seqType, year)
	if err != nil {
		return seq.Seq{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return seq.Seq{}, err
	}
	return d.scan(rows), nil
}

func (d *seqDAO) Delete(seqType seq.SeqType, year int) (seq.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE seq SET modified_time=unix_timestamp(now()),deleted=1 WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return seq.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year)
	if err != nil {
		return seq.Seq{}, err
	}

	return seq.Seq{}, nil
}

func (d *seqDAO) scan(rows *sql.Rows) seq.Seq {
	var seqType int
	var year int
	var value int
	rows.Scan(&seqType, &year, &value)
	return seq.Seq{
		SeqType: seq.SeqType(seqType),
		Year:    year,
		Value:   value,
	}
}
