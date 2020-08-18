package seq

import (
	"database/sql"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/fkmhrk/OpenInvoice/v1/service/model"
)

const (
	tableName    = "seq"
	sqlSelectAll = "SELECT seq_type,year,value " +
		"FROM " + tableName + " "
	sqlSelectByTypeAndYear = sqlSelectAll +
		"WHERE seq_type=? AND year=? AND deleted <> 1"
	sqlInsert = "INSERT INTO " + tableName +
		"(seq_type,year,value,created_time,modified_time,deleted) " +
		"VALUES(?,?,?,unix_timestamp(),unix_timestamp(),0)"
	sqlUpdate = "UPDATE " + tableName + " " +
		"SET seq_type=?, year=?, value=?,modified_time=unix_timestamp() " +
		"WHERE seq_type=? AND year=? AND deleted <> 1"
	sqlNext = "UPDATE " + tableName + " " +
		"SET value=value+1,modified_time=unix_timestamp() " +
		"WHERE seq_type=? AND year=? AND deleted <> 1"
	sqlSoftDelete = "UPDATE " + tableName + " " +
		"SET modified_time=unix_timestamp(),deleted=1 " +
		"WHERE seq_type=? AND year=? AND deleted <> 1"
)

type seqDAO struct {
	connection *db.Connection
}

// New creates instance
func New(connection *db.Connection) model.Seq {
	return &seqDAO{
		connection: connection,
	}
}

func (d *seqDAO) Create(seqType entity.SeqType, year, value int) (entity.Seq, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return entity.Seq{}, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlInsert)
	if err != nil {
		return entity.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year, value)
	if err != nil {
		return entity.Seq{}, err
	}

	tr.Commit()

	return entity.Seq{
		SeqType: seqType,
		Year:    year,
		Value:   value,
	}, nil
}

func (d *seqDAO) Get(seqType entity.SeqType, year int) (entity.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectByTypeAndYear)
	if err != nil {
		return entity.Seq{}, err
	}
	defer st.Close()

	rows, err := st.Query(seqType, year)
	if err != nil {
		return entity.Seq{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.Seq{}, nil
	}

	return d.scan(rows), nil
}

func (d *seqDAO) Update(seqType entity.SeqType, year, value int) (entity.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlUpdate)
	if err != nil {
		return entity.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year, value, seqType, year)
	if err != nil {
		return entity.Seq{}, err
	}

	return entity.Seq{
		SeqType: seqType,
		Year:    year,
		Value:   value,
	}, nil
}

func (d *seqDAO) Next(seqType entity.SeqType, year int) (entity.Seq, error) {
	tx, err := d.connection.Begin()
	if err != nil {
		return entity.Seq{}, err
	}
	defer tx.Rollback()

	seqEntity, err := d.doNext(tx, seqType, year)
	if err == nil {
		tx.Commit()
		return seqEntity, nil
	} else {
		return entity.Seq{}, err
	}
}

func (d *seqDAO) doNext(tx *sql.Tx, seqType entity.SeqType, year int) (entity.Seq, error) {
	// update
	st, err := tx.Prepare(sqlNext)
	if err != nil {
		return entity.Seq{}, err
	}
	defer st.Close()

	result, err := st.Exec(seqType, year)
	if err != nil {
		return entity.Seq{}, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		return entity.Seq{}, err
	}
	if rowCount == 0 {
		// create
		st2, err := tx.Prepare(sqlInsert)
		if err != nil {
			return entity.Seq{}, err
		}
		defer st2.Close()

		_, err = st2.Exec(seqType, year, 1)
		if err != nil {
			return entity.Seq{}, err
		}
		return entity.Seq{
			SeqType: seqType,
			Year:    year,
			Value:   1,
		}, nil
	}

	// get
	st2, err := tx.Prepare(sqlSelectByTypeAndYear)
	if err != nil {
		return entity.Seq{}, err
	}
	defer st2.Close()

	rows, err := st2.Query(seqType, year)
	if err != nil {
		return entity.Seq{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.Seq{}, err
	}
	return d.scan(rows), nil
}

func (d *seqDAO) Delete(seqType entity.SeqType, year int) (entity.Seq, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSoftDelete)
	if err != nil {
		return entity.Seq{}, err
	}
	defer st.Close()

	_, err = st.Exec(seqType, year)
	if err != nil {
		return entity.Seq{}, err
	}

	return entity.Seq{}, nil
}

func (d *seqDAO) scan(rows *sql.Rows) entity.Seq {
	var seqType int
	var year int
	var value int
	rows.Scan(&seqType, &year, &value)
	return entity.Seq{
		SeqType: entity.SeqType(seqType),
		Year:    year,
		Value:   value,
	}
}
