package main

import (
	"context"
	"errors"
	"log"
	"os"

	dk "github.com/takanoriyanagitani/go-pgzip/pgzip/dec/klauspost"
	. "github.com/takanoriyanagitani/go-pgzip/util"
)

var ErrNoFilename error = errors.New("no filename")

var stdin2stdout IO[Void] = func(_ context.Context) (Void, error) {
	_, e := dk.StdinToStdout()
	return Empty, e
}

var filename IO[string] = func(_ context.Context) (string, error) {
	var filenames []string = os.Args[1:]
	if len(filenames) < 1 {
		return "", ErrNoFilename
	}
	return filenames[0], nil
}

var input2writer IO[Void] = func(ctx context.Context) (Void, error) {
	name, e := filename(ctx)
	switch e {
	case nil:
		return dk.FilenameToStdout(name)(ctx)
	default:
		return stdin2stdout(ctx)
	}
}

var sub IO[Void] = func(ctx context.Context) (Void, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return input2writer(ctx)
}

func main() {
	_, e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
