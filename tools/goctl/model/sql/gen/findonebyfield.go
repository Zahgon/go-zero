package gen

type findOneCode struct {
	findOneMethod          string
	findOneInterfaceMethod string
	cacheExtra             string
}

func genFindOneByField(table Table, withCache, postgreSql bool) (*findOneCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertJoin(key Key, postgreSql bool) (in, paramJoinString, originalFieldString string) {
	_ = "STUB: not implemented"
	return "", "", ""
}
