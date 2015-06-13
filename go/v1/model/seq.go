package model

type SeqDAO interface {
	Create(seqType, year, value int) (Seq, error)
	Get(seqType int, year int) (Seq, error)
	Update(seqType, year, value int) (Seq, error)
	Delete(seqType int, year int) (Seq, error)
}

type Seq struct {
	SeqType int
	Year    int
	Value   int
}

func (o Seq) IsEmpty() bool {
	return o.SeqType == 0 && o.Year == 0
}
