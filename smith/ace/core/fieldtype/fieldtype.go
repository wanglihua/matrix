package fieldtype

const (
	Int = iota
	BigInt
	Decimal
	NVarchar
	DateTime
	Boolean
	Create
	Update
	Version
)

type FieldType int
