package protoc

var url = map[string]string{
	"linux_32":   "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_32.zip",
	"linux_64":   "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip",
	"darwin":     "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-osx-x86_64.zip",
	"windows_32": "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-win32.zip",
	"windows_64": "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-win64.zip",
}

const (
	Name        = "protoc"
	ZipFileName = Name + ".zip"
)

func Install(cacheDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Exists() bool { _ = "STUB: not implemented"; return false }

func Version() (string, error) { _ = "STUB: not implemented"; return "", nil }
