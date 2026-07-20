package kube

import (
	_ "embed"

	"github.com/spf13/cobra"
)

const (
	category           = "kube"
	deployTemplateFile = "deployment.tpl"
	jobTemplateFile    = "job.tpl"
	basePort           = 30000
	portLimit          = 32767
)

var (
	//go:embed deployment.tpl
	deploymentTemplate string
	//go:embed job.tpl
	jobTemplate string
)

type Deployment struct {
	Name            string
	Namespace       string
	Image           string
	Secret          string
	Replicas        int
	Revisions       int
	Port            int
	TargetPort      int
	NodePort        int
	UseNodePort     bool
	RequestCpu      int
	RequestMem      int
	LimitCpu        int
	LimitMem        int
	MinReplicas     int
	MaxReplicas     int
	ServiceAccount  string
	ImagePullPolicy string
}

func deploymentCommand(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func Category() string { _ = "STUB: not implemented"; return "" }

func Clean() error { _ = "STUB: not implemented"; return nil }

func GenTemplates() error { _ = "STUB: not implemented"; return nil }

func RevertTemplate(name string) error { _ = "STUB: not implemented"; return nil }

func Update() error { _ = "STUB: not implemented"; return nil }
