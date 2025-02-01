package enc

import (
	"bufio"
	"context"
	"io"
	"os"

	kp "github.com/klauspost/pgzip"
	. "github.com/takanoriyanagitani/go-pgzip/util"
)

func PgReaderToWriter(rdr *kp.Reader, wtr io.Writer) (int64, error) {
	return io.Copy(wtr, rdr)
}

func ReaderToWriter(rdr io.Reader, wtr io.Writer) (int64, error) {
	pr, e := kp.NewReader(rdr) // uses internal bufio
	if nil != e {
		return 0, e
	}
	defer pr.Close()

	return PgReaderToWriter(pr, wtr)
}

func FilenameToWriter(filename string, wtr io.Writer) (int64, error) {
	f, e := os.Open(filename)
	if nil != e {
		return 0, e
	}
	defer f.Close()

	return ReaderToWriter(f, wtr)
}

func FilenameToStdout(filename string) IO[Void] {
	return func(_ context.Context) (Void, error) {
		var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
		defer bw.Flush()
		_, e := FilenameToWriter(filename, bw)
		return Empty, e
	}
}

func StdinToStdout() (int64, error) {
	var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer bw.Flush()

	return ReaderToWriter(os.Stdin, bw)
}
