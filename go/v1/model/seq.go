package model

type SeqDAO interface {
	Create(seqType SeqType, year, value int) (Seq, error)
	Get(seqType SeqType, year int) (Seq, error)
	Update(seqType SeqType, year, value int) (Seq, error)
	Next(seqType SeqType, year int) (Seq, error)
	Delete(seqType SeqType, year int) (Seq, error)
}

type Seq struct {
	SeqType SeqType
	Year    int
	Value   int
}

func (o Seq) IsEmpty() bool {
	return o.SeqType == 0 && o.Year == 0
}

type SeqType int

const (
	SeqType_Null SeqType = iota
	SeqType_Quotation
	SeqType_Delivery
	SeqType_Bill
)
