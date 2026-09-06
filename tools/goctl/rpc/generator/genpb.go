package generator

func (g *Generator) GenPb(ctx DirContext, c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genPbDirect(ctx DirContext, c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) buildProtocCmd(c *ZRpcContext, pwd string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func relativeToProtoPath(f string, protoPaths []string, pwd string) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) setPbDir(ctx DirContext, c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	pbSuffix   = "pb.go"
	grpcSuffix = "_grpc.pb.go"
)

func findPbFile(current string, src string, grpc bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
