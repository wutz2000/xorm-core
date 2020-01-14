package core

const (
	AssociateNone = iota
	AssociateHasOne
	AssociateHasMany
)

type Associate struct {
	AssociateFieldName string
	AssociateTableName string
	Condition          interface{}
	Params             []interface{}
	AssociateType      int64
}
