package gen

type Field struct {
	NameOriginal    string
	UpperName       string
	LowerName       string
	DataType        string
	Comment         string
	SeqInIndex      int
	OrdinalPosition int
}

func genCustomized(table Table, withCache, postgreSql bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
