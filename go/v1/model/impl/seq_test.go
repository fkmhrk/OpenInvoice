package impl

import (
	m "../"
	"database/sql"
	"testing"
)

func createSeqDAO(db *sql.DB) *seqDAO {
	return NewSeqDAO(NewConnection(db))
}

func assertSeq(t *testing.T, item *m.Seq, seqType, year, value int) {
	caller := getCaller()
	if item.SeqType != seqType {
		t.Errorf("[%s] seqType must be %d but %d", caller, seqType, item.SeqType)
	}
	if item.Year != year {
		t.Errorf("[%s] year must be %d but %d", caller, year, item.Year)
	}
	if item.Value != value {
		t.Errorf("[%s] value must be %d but %d", caller, value, item.Value)
	}
}

func hardDeleteSeq(db *sql.DB, seqType int, year int) {
	s, _ := db.Prepare("DELETE FROM seq WHERE seq_type=? AND year=?")
	defer s.Close()
	s.Exec(seqType, year)
}

func TestSeq_All(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createSeqDAO(db)

	var seqType int = 10
	var year int = 100
	var value int = 100
	hardDeleteSeq(db, seqType, year)
	item, err := dao.Create(seqType, year, value)
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertSeq(t, &item, 10, 100, 100)

	item2, err := dao.Get(seqType, year)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSeq(t, &item2, 10, 100, 100)

	value = 200
	item3, err := dao.Update(seqType, year, value)
	if err != nil {
		t.Errorf("Failed to Update : %s", err)
		return
	}
	assertSeq(t, &item3, 10, 100, 200)

	item4, err := dao.Get(seqType, year)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSeq(t, &item4, 10, 100, 200)

	item5, err := dao.Delete(seqType, year)
	if err != nil {
		t.Errorf("Failed to Delete : %s", err)
		return
	}
	if !item5.IsEmpty() {
		t.Errorf("Returned item must be empty")
		return
	}
	item6, err := dao.Get(seqType, year)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	if !item6.IsEmpty() {
		t.Errorf("Item must be empty")
		return
	}
}
