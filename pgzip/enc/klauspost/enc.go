package enc

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"io"
	"os"

	kp "github.com/klauspost/pgzip"
	pg "github.com/takanoriyanagitani/go-pgzip"
	. "github.com/takanoriyanagitani/go-pgzip/util"
)

type Level pg.EncodeLevel

type LevelConversionMap map[pg.EncodeLevel]int

var LevelConvMap LevelConversionMap = map[pg.EncodeLevel]int{
	pg.EncodeLevelStore:       kp.NoCompression,
	pg.EncodeLevelFast:        kp.BestSpeed,
	pg.EncodeLevelBest:        kp.BestCompression,
	pg.EncodeLevelDefault:     kp.DefaultCompression,
	pg.EncodeLevelHuffmanOnly: kp.HuffmanOnly,
	pg.EncodeLevelConstant:    kp.ConstantCompression,
}

func (l Level) ToLevel() (ret sql.Null[int]) {
	val, found := LevelConvMap[pg.EncodeLevel(l)]
	ret.Valid = found
	ret.V = val
	return ret
}

type Config pg.EncodeConfig

func ReaderToPgzipWriter(rdr io.Reader, pgw *kp.Writer) (int64, error) {
	return io.Copy(pgw, rdr)
}

func ReaderToWriterWithLevel(
	rdr io.Reader,
	wtr io.Writer,
	level sql.Null[int],
) (int64, error) {
	var pw *kp.Writer
	var e error

	switch level.Valid {
	case true:
		pw, e = kp.NewWriterLevel(wtr, level.V)
	default:
		pw = kp.NewWriter(wtr)
	}
	if nil != e {
		return 0, e
	}

	i, e := ReaderToPgzipWriter(rdr, pw)
	return i, errors.Join(e, pw.Close())
}

func (c Config) ReaderToWriter(rdr io.Reader, wtr io.Writer) (int64, error) {
	var l Level = Level(c.EncodeLevel)
	return ReaderToWriterWithLevel(
		rdr,
		wtr,
		l.ToLevel(),
	)
}

func (c Config) ToStdinToStdout() IO[int64] {
	return func(_ context.Context) (int64, error) {
		var br io.Reader = bufio.NewReader(os.Stdin)

		var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
		defer bw.Flush()

		return c.ReaderToWriter(br, bw)
	}
}
