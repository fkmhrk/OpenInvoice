package impl

import (
	"database/sql"

	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
)

type seqDAO struct {
	connection *db.Connection
}

func NewSeqDAO(connection *db.Connection) *seqDAO {
	return &seqDAO{
		connection: connection,
	}
}

func (d *seqDAO) Create(seqType m.SeqType, year, value int) (m.Seq, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return m.Seq{}, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO seq(seq_type,year,value,created_time,modified_time,deleted) VALUES(?,?,?,unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return m.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year, value)
	if err != nil {
		return m.Seq{}, err
	}

	tr.Commit()

	return m.Seq{
		SeqType: seqType,
		Year:    year,
		Value:   value,
	}, nil
}

func (d *seqDAO) Get(seqType m.SeqType, year int) (m.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT seq_type,year,value FROM seq WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return m.Seq{}, err
	}
	defer st.Close()

	rows, err := st.Query(seqType, year)
	if err != nil {
		return m.Seq{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return m.Seq{}, nil
	}

	return d.scan(rows), nil
}

func (d *seqDAO) Update(seqType m.SeqType, year, value int) (m.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE seq SET seq_type=?, year=?, value=?,modified_time=unix_timestamp(now()) WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return m.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year, value, seqType, year)
	if err != nil {
		return m.Seq{}, err
	}

	return m.Seq{
		SeqType: seqType,
		Year:    year,
		Value:   value,
	}, nil
}

func (d *seqDAO) Next(seqType m.SeqType, year int) (m.Seq, error) {
	tx, err := d.connection.Begin()
	if err != nil {
		return m.Seq{}, err
	}
	defer tx.Rollback()

	seq, err := d.doNext(tx, seqType, year)
	if err == nil {
		tx.Commit()
		return seq, nil
	} else {
		return m.Seq{}, err
	}
}

func (d *seqDAO) doNext(tx *sql.Tx, seqType m.SeqType, year int) (m.Seq, error) {
	// update
	st, err := tx.Prepare("UPDATE seq SET value=value+1,modified_time=unix_timestamp(now()) WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return m.Seq{}, err
	}
	defer st.Close()

	result, err := st.Exec(seqType, year)
	if err != nil {
		return m.Seq{}, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		return m.Seq{}, err
	}
	if rowCount == 0 {
		// create
		st2, err := tx.Prepare("INSERT INTO seq(seq_type,year,value,created_time,modified_time,deleted) VALUES(?,?,?,unix_timestamp(now()),unix_timestamp(now()),0)")
		if err != nil {
			return m.Seq{}, err
		}
		defer st2.Close()

		_, err = st2.Exec(seqType, year, 1)
		if err != nil {
			return m.Seq{}, err
		}
		return m.Seq{
			SeqType: seqType,
			Year:    year,
			Value:   1,
		}, nil
	}

	// get
	st2, err := tx.Prepare("SELECT seq_type,year,value FROM seq WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return m.Seq{}, err
	}
	defer st2.Close()

	rows, err := st2.Query(seqType, year)
	if err != nil {
		return m.Seq{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return m.Seq{}, err
	}
	return d.scan(rows), nil
}

func (d *seqDAO) Delete(seqType m.SeqType, year int) (m.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE seq SET modified_time=unix_timestamp(now()),deleted=1 WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return m.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year)
	if err != nil {
		return m.Seq{}, err
	}

	return m.Seq{}, nil
}

func (d *seqDAO) scan(rows *sql.Rows) m.Seq {
	var seqType int
	var year int
	var value int
	rows.Scan(&seqType, &year, &value)
	return m.Seq{
		SeqType: m.SeqType(seqType),
		Year:    year,
		Value:   value,
	}
}
