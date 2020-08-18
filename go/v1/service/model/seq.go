package model

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type Seq interface {
	Create(seqType entity.SeqType, year, value int) (entity.Seq, error)
	Get(seqType entity.SeqType, year int) (entity.Seq, error)
	Update(seqType entity.SeqType, year, value int) (entity.Seq, error)
	Next(seqType entity.SeqType, year int) (entity.Seq, error)
	Delete(seqType entity.SeqType, year int) (entity.Seq, error)
}
