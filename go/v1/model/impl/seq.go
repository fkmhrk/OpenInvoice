package impl

import (
	m "../"
	"database/sql"
)

type seqDAO struct {
	connection *Connection
}

func NewSeqDAO(connection *Connection) *seqDAO {
	return &seqDAO{
		connection: connection,
	}
}

func (d *seqDAO) Create(seqType, year, value int) (m.Seq, error) {
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

func (d *seqDAO) Get(seqType int, year int) (m.Seq, error) {
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

func (d *seqDAO) Update(seqType, year, value int) (m.Seq, error) {
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

func (d *seqDAO) Next(seqType, year int) (m.Seq, error) {
	tx, err := d.connection.Begin()
	if err != nil {
		return m.Seq{}, err
	}
	defer tx.Rollback()

	// get
	st, err := tx.Prepare("SELECT seq_type,year,value FROM seq WHERE seq_type=? AND year=? AND deleted <> 1 FOR UPDATE")
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
		return d.Create(seqType, year, 1)
	}
	seq := d.scan(rows)

	// update
	seq.Value += 1
	st, err = tx.Prepare("UPDATE seq SET value=?,modified_time=unix_timestamp(now()) WHERE seq_type=? AND year=? AND deleted <> 1")
	if err != nil {
		return m.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seq.Value, seqType, year)
	if err != nil {
		return m.Seq{}, err
	}

	tx.Commit()

	return seq, nil

}

func (d *seqDAO) Delete(seqType int, year int) (m.Seq, error) {
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
		SeqType: seqType,
		Year:    year,
		Value:   value,
	}
}
