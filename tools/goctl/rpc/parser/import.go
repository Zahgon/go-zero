package parser

import (
	"github.com/emicklei/proto"
)

type Import struct {
	*proto.Import
}

type ImportedProto struct {
	Src string

	ProtoPackage string

	GoPackage string

	PbPackage string
}

func BuildProtoPackageMap(importedProtos []ImportedProto) map[string]ImportedProto {
	_ = "STUB: not implemented"
	return nil
}

func ResolveImports(src string, protoPaths []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseImportedProtos(src string, protoPaths []string) ([]ImportedProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectImports(src string, protoPaths []string, visited map[string]bool, result *[]string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseImportFilenames(src string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func parseGoPackage(src string) (goPackage, pbPackage, protoPackage string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func lookupProtoFile(filename string, protoPaths []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isWellKnownProto(filename string) bool { _ = "STUB: not implemented"; return false }
