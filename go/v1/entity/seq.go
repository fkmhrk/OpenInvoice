package entity

// Seq is entity
type Seq struct {
	SeqType SeqType
	Year    int
	Value   int
}

// IsEmpty determines this is empty
func (o Seq) IsEmpty() bool {
	return o.SeqType == 0 && o.Year == 0
}

// SeqType is type
type SeqType int

const (
	SeqType_Null SeqType = iota
	SeqType_Quotation
	SeqType_Delivery
	SeqType_Bill
)
