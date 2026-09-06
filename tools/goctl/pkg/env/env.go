package env

import (
	"log"
	"runtime"
	"strings"
	"testing"

	"github.com/zeromicro/go-zero/tools/goctl/internal/version"
	sortedmap "github.com/zeromicro/go-zero/tools/goctl/pkg/collection"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protoc"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protocgengo"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protocgengogrpc"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var goctlEnv *sortedmap.SortedMap

const (
	GoctlOS                = "GOCTL_OS"
	GoctlArch              = "GOCTL_ARCH"
	GoctlHome              = "GOCTL_HOME"
	GoctlDebug             = "GOCTL_DEBUG"
	GoctlCache             = "GOCTL_CACHE"
	GoctlVersion           = "GOCTL_VERSION"
	GoctlExperimental      = "GOCTL_EXPERIMENTAL"
	ProtocVersion          = "PROTOC_VERSION"
	ProtocGenGoVersion     = "PROTOC_GEN_GO_VERSION"
	ProtocGenGoGRPCVersion = "PROTO_GEN_GO_GRPC_VERSION"

	envFileDir      = "env"
	ExperimentalOn  = "on"
	ExperimentalOff = "off"
)

func init() {
	defaultGoctlHome, err := pathx.GetDefaultGoctlHome()
	if err != nil {
		log.Fatalln(err)
	}
	goctlEnv = sortedmap.New()
	goctlEnv.SetKV(GoctlOS, runtime.GOOS)
	goctlEnv.SetKV(GoctlArch, runtime.GOARCH)
	existsEnv := readEnv(defaultGoctlHome)
	if existsEnv != nil {
		goctlHome, ok := existsEnv.GetString(GoctlHome)
		if ok && len(goctlHome) > 0 {
			goctlEnv.SetKV(GoctlHome, goctlHome)
		}
		if debug := existsEnv.GetOr(GoctlDebug, "").(string); debug != "" {
			if strings.EqualFold(debug, "true") || strings.EqualFold(debug, "false") {
				goctlEnv.SetKV(GoctlDebug, debug)
			}
		}
		if value := existsEnv.GetStringOr(GoctlCache, ""); value != "" {
			goctlEnv.SetKV(GoctlCache, value)
		}
		experimental := existsEnv.GetOr(GoctlExperimental, ExperimentalOn)
		goctlEnv.SetKV(GoctlExperimental, experimental)
	}

	if !goctlEnv.HasKey(GoctlHome) {
		goctlEnv.SetKV(GoctlHome, defaultGoctlHome)
	}
	if !goctlEnv.HasKey(GoctlDebug) {
		goctlEnv.SetKV(GoctlDebug, "False")
	}

	if !goctlEnv.HasKey(GoctlCache) {
		cacheDir, _ := pathx.GetCacheDir()
		goctlEnv.SetKV(GoctlCache, cacheDir)
	}

	if !goctlEnv.HasKey(GoctlExperimental) {
		goctlEnv.SetKV(GoctlExperimental, ExperimentalOn)
	}

	goctlEnv.SetKV(GoctlVersion, version.BuildVersion)

	protocVer, _ := protoc.Version()
	goctlEnv.SetKV(ProtocVersion, protocVer)

	protocGenGoVer, _ := protocgengo.Version()
	goctlEnv.SetKV(ProtocGenGoVersion, protocGenGoVer)

	protocGenGoGrpcVer, _ := protocgengogrpc.Version()
	goctlEnv.SetKV(ProtocGenGoGRPCVersion, protocGenGoGrpcVer)
}

func Print(args ...string) string { _ = "STUB: not implemented"; return "" }

func Get(key string) string { _ = "STUB: not implemented"; return "" }

func Set(t *testing.T, key, value string) { _ = "STUB: not implemented"; return }

func GetOr(key, def string) string { _ = "STUB: not implemented"; return "" }

func UseExperimental() bool { _ = "STUB: not implemented"; return false }

func readEnv(goctlHome string) *sortedmap.SortedMap { _ = "STUB: not implemented"; return nil }

func WriteEnv(kv []string) error { _ = "STUB: not implemented"; return nil }
