package trace

const TraceName = "go-zero"

type Config struct {
	Name     string  `json:",optional"`
	Endpoint string  `json:",optional"`
	Sampler  float64 `json:",default=1.0"`
	Batcher  string  `json:",default=otlpgrpc,options=zipkin|otlpgrpc|otlphttp|file"`

	OtlpHeaders map[string]string `json:",optional"`

	OtlpHttpPath string `json:",optional"`

	OtlpHttpSecure bool `json:",optional"`

	Disabled bool `json:",optional"`
}
