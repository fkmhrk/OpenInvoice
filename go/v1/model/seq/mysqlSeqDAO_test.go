package seq

import (
	"database/sql"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	testdb "github.com/fkmhrk/OpenInvoice/v1/model/db/test"
)

func createSeqDAO(sqlDB *sql.DB) model.Seq {
	return New(db.NewConnection(sqlDB))
}

func assertSeq(t *testing.T, item *entity.Seq, seqType entity.SeqType, year, value int) {
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

func hardDeleteSeq(db *sql.DB, seqType entity.SeqType, year int) {
	s, _ := db.Prepare("DELETE FROM seq WHERE seq_type=? AND year=?")
	defer s.Close()
	s.Exec(seqType, year)
}

func TestSeq_All(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createSeqDAO(db)

	var seqType entity.SeqType = 10
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

func TestSeq_Next(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createSeqDAO(db)

	var seqType entity.SeqType = 11
	var year int = 100
	hardDeleteSeq(db, seqType, year)
	item, err := dao.Next(seqType, year)
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertSeq(t, &item, 11, 100, 1)

	item2, err := dao.Next(seqType, year)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSeq(t, &item2, 11, 100, 2)

	item3, err := dao.Get(seqType, year)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSeq(t, &item3, 11, 100, 2)
}

func getCaller() string {
	_, caller, line, _ := runtime.Caller(2)
	path := strings.SplitN(caller, "/", -1)
	return fmt.Sprintf("%s:%d", path[len(path)-1], line)
}
