package converter

import (
	"github.com/zeromicro/ddl-parser/parser"
)

var unsignedTypeMap = map[string]string{
	"int":   "uint",
	"int8":  "uint8",
	"int16": "uint16",
	"int32": "uint32",
	"int64": "uint64",
}

var commonMysqlDataTypeMapInt = map[int]string{

	parser.Bit:       "byte",
	parser.TinyInt:   "int64",
	parser.SmallInt:  "int64",
	parser.MediumInt: "int64",
	parser.Int:       "int64",
	parser.MiddleInt: "int64",
	parser.Int1:      "int64",
	parser.Int2:      "int64",
	parser.Int3:      "int64",
	parser.Int4:      "int64",
	parser.Int8:      "int64",
	parser.Integer:   "int64",
	parser.BigInt:    "int64",
	parser.Float:     "float64",
	parser.Float4:    "float64",
	parser.Float8:    "float64",
	parser.Double:    "float64",
	parser.Decimal:   "float64",
	parser.Dec:       "float64",
	parser.Fixed:     "float64",
	parser.Numeric:   "float64",
	parser.Real:      "float64",

	parser.Date:      "time.Time",
	parser.DateTime:  "time.Time",
	parser.Timestamp: "time.Time",
	parser.Time:      "string",
	parser.Year:      "int64",

	parser.Char:            "string",
	parser.VarChar:         "string",
	parser.NVarChar:        "string",
	parser.NChar:           "string",
	parser.Character:       "string",
	parser.LongVarChar:     "string",
	parser.LineString:      "string",
	parser.MultiLineString: "string",
	parser.Binary:          "string",
	parser.VarBinary:       "string",
	parser.TinyText:        "string",
	parser.Text:            "string",
	parser.MediumText:      "string",
	parser.LongText:        "string",
	parser.Enum:            "string",
	parser.Set:             "string",
	parser.Json:            "string",
	parser.Blob:            "string",
	parser.LongBlob:        "string",
	parser.MediumBlob:      "string",
	parser.TinyBlob:        "string",

	parser.Bool:    "bool",
	parser.Boolean: "bool",
}

var commonMysqlDataTypeMap = map[int]string{

	parser.Bit:       "bit",
	parser.TinyInt:   "tinyint",
	parser.SmallInt:  "smallint",
	parser.MediumInt: "mediumint",
	parser.Int:       "int",
	parser.MiddleInt: "middleint",
	parser.Int1:      "int1",
	parser.Int2:      "int2",
	parser.Int3:      "int3",
	parser.Int4:      "int4",
	parser.Int8:      "int8",
	parser.Integer:   "integer",
	parser.BigInt:    "bigint",
	parser.Float:     "float",
	parser.Float4:    "float4",
	parser.Float8:    "float8",
	parser.Double:    "double",
	parser.Decimal:   "decimal",
	parser.Dec:       "dec",
	parser.Fixed:     "fixed",
	parser.Numeric:   "numeric",
	parser.Real:      "real",

	parser.Date:      "date",
	parser.DateTime:  "datetime",
	parser.Timestamp: "timestamp",
	parser.Time:      "time",
	parser.Year:      "year",

	parser.Char:            "char",
	parser.VarChar:         "varchar",
	parser.NVarChar:        "nvarchar",
	parser.NChar:           "nchar",
	parser.Character:       "character",
	parser.LongVarChar:     "longvarchar",
	parser.LineString:      "linestring",
	parser.MultiLineString: "multilinestring",
	parser.Binary:          "binary",
	parser.VarBinary:       "varbinary",
	parser.TinyText:        "tinytext",
	parser.Text:            "text",
	parser.MediumText:      "mediumtext",
	parser.LongText:        "longtext",
	parser.Enum:            "enum",
	parser.Set:             "set",
	parser.Json:            "json",
	parser.Blob:            "blob",
	parser.LongBlob:        "longblob",
	parser.MediumBlob:      "mediumblob",
	parser.TinyBlob:        "tinyblob",

	parser.Bool:    "bool",
	parser.Boolean: "boolean",
}

var commonMysqlDataTypeMapString = map[string]string{

	"bool":    "bool",
	"_bool":   "pq.BoolArray",
	"boolean": "bool",

	"tinyint":   "int64",
	"smallint":  "int64",
	"mediumint": "int64",
	"int":       "int64",
	"int1":      "int64",
	"int2":      "int64",
	"_int2":     "pq.Int64Array",
	"int3":      "int64",
	"int4":      "int64",
	"_int4":     "pq.Int64Array",
	"int8":      "int64",
	"_int8":     "pq.Int64Array",
	"integer":   "int64",
	"_integer":  "pq.Int64Array",
	"bigint":    "int64",
	"float":     "float64",
	"float4":    "float64",
	"_float4":   "pq.Float64Array",
	"float8":    "float64",
	"_float8":   "pq.Float64Array",
	"double":    "float64",
	"decimal":   "float64",
	"dec":       "float64",
	"fixed":     "float64",
	"real":      "float64",
	"bit":       "byte",

	"date":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"time":      "string",
	"year":      "int64",

	"linestring":      "string",
	"multilinestring": "string",
	"nvarchar":        "string",
	"nchar":           "string",
	"char":            "string",
	"bpchar":          "string",
	"_char":           "pq.StringArray",
	"character":       "string",
	"varchar":         "string",
	"_varchar":        "pq.StringArray",
	"binary":          "string",
	"bytea":           "string",
	"longvarbinary":   "string",
	"varbinary":       "string",
	"tinytext":        "string",
	"text":            "string",
	"_text":           "pq.StringArray",
	"mediumtext":      "string",
	"longtext":        "string",
	"enum":            "string",
	"set":             "string",
	"json":            "string",
	"jsonb":           "string",
	"blob":            "string",
	"longblob":        "string",
	"mediumblob":      "string",
	"tinyblob":        "string",
	"ltree":           "[]byte",
}

func ConvertDataType(dataBaseType int, isDefaultNull, unsigned, strict bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func ConvertStringDataType(dataBaseType string, isDefaultNull, unsigned, strict bool) (
	goType string, thirdPkg string, isPQArray bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func convertDatatypeWithConfig(dataBaseType string, isDefaultNull, unsigned bool) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func mayConvertNullType(goDataType string, isDefaultNull, unsigned, strict bool) string {
	_ = "STUB: not implemented"
	return ""
}
