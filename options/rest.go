package options

// REST store configuration about REST service
var REST = struct {
	Address string `long:"address" env:"REST_ADDRESS" default:"0.0.0.0"`
	Port    int    `long:"port" env:"REST_PORT" default:"8080"`
}{}

func init() {
	registry[group{"rest", "REST configuration"}] = &REST
}
