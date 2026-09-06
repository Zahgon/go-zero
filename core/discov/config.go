package discov

import "errors"

var (
	errEmptyEtcdHosts = errors.New("empty etcd hosts")

	errEmptyEtcdKey = errors.New("empty etcd key")
)

type EtcdConf struct {
	Hosts              []string
	Key                string
	ID                 int64  `json:",optional"`
	User               string `json:",optional"`
	Pass               string `json:",optional"`
	CertFile           string `json:",optional"`
	CertKeyFile        string `json:",optional=CertFile"`
	CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool   `json:",optional"`
}

func (c EtcdConf) HasAccount() bool { _ = "STUB: not implemented"; return false }

func (c EtcdConf) HasID() bool { _ = "STUB: not implemented"; return false }

func (c EtcdConf) HasTLS() bool { _ = "STUB: not implemented"; return false }

func (c EtcdConf) Validate() error { _ = "STUB: not implemented"; return nil }
