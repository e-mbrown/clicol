package base

import (
	"io"
	"os"
)

type Opt func(*CliConf) error

type CliConf struct {
	ErrStream, OutStream io.Writer
}

func NewCliConfig(opts ...Opt) (CliConf, error) {
	c := CliConf{
		ErrStream: os.Stderr,
		OutStream: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(&c); err != nil {
			return CliConf{}, err
		}
	}
	return c, nil
}

func WithErrStream(errStream io.Writer) Opt {
	return func(c *CliConf) error {
		c.ErrStream = errStream
		return nil
	}
}
func WithOutStream(outStream io.Writer) Opt {
	return func(c *CliConf) error {
		c.OutStream = outStream
		return nil
	}
}
