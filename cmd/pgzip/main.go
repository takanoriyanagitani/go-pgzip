package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pg "github.com/takanoriyanagitani/go-pgzip"
	ek "github.com/takanoriyanagitani/go-pgzip/pgzip/enc/klauspost"
	. "github.com/takanoriyanagitani/go-pgzip/util"
)

func envValByKey(key string) IO[string] {
	return func(_ context.Context) (string, error) {
		val, found := os.LookupEnv(key)
		switch found {
		case true:
			return val, nil
		default:
			return "", fmt.Errorf("env var %s missing", key)
		}
	}
}

var encLevel IO[pg.EncodeLevel] = Bind(
	envValByKey("ENV_ENCODE_LEVEL"),
	Lift(func(l string) (pg.EncodeLevel, error) {
		return pg.EncodeLevelFromStr(l), nil
	}),
).Or(Of(pg.EncodeLevelDefault))

var encCfg IO[pg.EncodeConfig] = Bind(
	encLevel,
	Lift(func(l pg.EncodeLevel) (pg.EncodeConfig, error) {
		return pg.EncodeConfig{EncodeLevel: l}, nil
	}),
)

var stdin2stdout IO[Void] = Bind(
	encCfg,
	func(cfg pg.EncodeConfig) IO[Void] {
		return Bind(
			ek.Config(cfg).ToStdinToStdout(),
			Lift(func(_ int64) (Void, error) { return Empty, nil }),
		)
	},
)

var sub IO[Void] = func(ctx context.Context) (Void, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return stdin2stdout(ctx)
}

func main() {
	_, e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
