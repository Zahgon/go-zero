package cli

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	errInvalidGrpcOutput = errors.New("ZRPC: missing --go-grpc_out")
	errInvalidGoOutput   = errors.New("ZRPC: missing --go_out")
	errInvalidZrpcOutput = errors.New("ZRPC: missing zrpc output, please use --zrpc_out to specify the output")
)

func ZRPC(_ *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func wrapProtocCmd(name string, args []string) []string { _ = "STUB: not implemented"; return nil }
