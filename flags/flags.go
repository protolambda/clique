package flags

import "github.com/urfave/cli"

const envVarPrefix = "CLIQUE_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	EngineAddr = cli.StringFlag{
		Name:        "engine.addr",
		Usage:       "Engine API RPC address.",
		EnvVar:      prefixEnvVar("ENGINE_ADDR"),
		Required:    true,
		Value:       "",
		Destination: new(string),
	}
	EngineJWTSecret = cli.StringFlag{
		Name:        "engine.jwtsecret",
		Usage:       "Path to JWT secret key. Keys are 32 bytes, hex encoded in a file. A new key will be generated if left empty.",
		EnvVar:      prefixEnvVar("ENGINE_JWT_SECRET"),
		Required:    false,
		Value:       "",
		Destination: new(string),
	}
)

var Flags = []cli.Flag{
	EngineAddr,
	EngineJWTSecret,
}
