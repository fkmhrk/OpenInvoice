package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/entity"
)

type SeqDAO struct {
	CreateResult entity.Seq
	GetResult    entity.Seq
	UpdateResult entity.Seq
	NextResult   entity.Seq
	NextSeqType  entity.SeqType
	DeleteResult entity.Seq
}

func (d *SeqDAO) Create(seqType entity.SeqType, year, value int) (entity.Seq, error) {
	return d.CreateResult, nil
}

func (d *SeqDAO) Get(seqType entity.SeqType, year int) (entity.Seq, error) {
	return d.GetResult, nil
}

func (d *SeqDAO) Update(seqType entity.SeqType, year, value int) (entity.Seq, error) {
	return d.UpdateResult, nil
}

func (d *SeqDAO) Next(seqType entity.SeqType, year int) (entity.Seq, error) {
	// set args
	d.NextSeqType = seqType
	return d.NextResult, nil
}

func (d *SeqDAO) Delete(seqType entity.SeqType, year int) (entity.Seq, error) {
	return d.DeleteResult, nil
}
