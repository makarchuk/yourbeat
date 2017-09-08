// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config


type Config struct {
	Metrics []*MetricsConfig `config:"metrics"`
}

type MetricsConfig struct{
  Test int
}

var DefaultConfig = Config{
	Metrics: []*MetricsConfig{&MetricsConfig{Test: 10}},
}
