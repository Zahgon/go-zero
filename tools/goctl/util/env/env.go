package env

const (
	bin                = "bin"
	binGo              = "go"
	binProtoc          = "protoc"
	binProtocGenGo     = "protoc-gen-go"
	binProtocGenGrpcGo = "protoc-gen-go-grpc"
	cstOffset          = 60 * 60 * 8
)

func InChina() bool { _ = "STUB: not implemented"; return false }

func LookUpGo() (string, error) { _ = "STUB: not implemented"; return "", nil }

func LookUpProtoc() (string, error) { _ = "STUB: not implemented"; return "", nil }

func LookUpProtocGenGo() (string, error) { _ = "STUB: not implemented"; return "", nil }

func LookUpProtocGenGoGrpc() (string, error) { _ = "STUB: not implemented"; return "", nil }

func LookPath(xBin string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func CanExec() bool { _ = "STUB: not implemented"; return false }

func getExeSuffix() string { _ = "STUB: not implemented"; return "" }
