package fieldtype

const (
    Int = iota
    Decimal
    NVarchar
    DateTime
    Boolean
    Create
    Update
    Version
)

type FieldType int

