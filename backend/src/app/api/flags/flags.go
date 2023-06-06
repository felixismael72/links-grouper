package flags

import "flag"

type Flags struct {
	EnvType *string
}

func Parse() *Flags {
	flags := &Flags{}
	flags.EnvType = flag.String("env-type", "development", "Specify the environment type")
	flag.Parse()
	return flags
}
