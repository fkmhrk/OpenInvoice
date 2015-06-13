package mock

import (
	m "../"
)

type SeqDAO struct {
	CreateResult m.Seq
	GetResult    m.Seq
	UpdateResult m.Seq
	NextResult   m.Seq
	NextSeqType  m.SeqType
	DeleteResult m.Seq
}

func (d *SeqDAO) Create(seqType m.SeqType, year, value int) (m.Seq, error) {
	return d.CreateResult, nil
}

func (d *SeqDAO) Get(seqType m.SeqType, year int) (m.Seq, error) {
	return d.GetResult, nil
}

func (d *SeqDAO) Update(seqType m.SeqType, year, value int) (m.Seq, error) {
	return d.UpdateResult, nil
}

func (d *SeqDAO) Next(seqType m.SeqType, year int) (m.Seq, error) {
	// set args
	d.NextSeqType = seqType
	return d.NextResult, nil
}

func (d *SeqDAO) Delete(seqType m.SeqType, year int) (m.Seq, error) {
	return d.DeleteResult, nil
}
