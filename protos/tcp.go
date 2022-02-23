package protos

import (
	"context"
	"io"
	"net"
	"net/url"
	"time"

	"log"
)

func init() {
	RegisterProtoAdaptFunc("tcp", newTcpProc)
}

type tcpProc struct {
	name   string
	furl   *url.URL
	input  io.Reader
	output io.Writer
}

func newTcpProc(name string, r io.Reader, furl *url.URL, reqs ...[]byte) protoAdaptor {
	p := &tcpProc{
		name:   name,
		furl:   furl,
		input:  r,
		output: nil,
	}
	return p
}

func (p tcpProc) GetResp() [][]byte {
	return nil
}

func (p tcpProc) Neg(context.Context) int {
	conn, err := net.DialTimeout(p.furl.Scheme, p.furl.Host, time.Second*3)
	if err != nil {
		log.Println(err)
		return -1
	}
	p.input = conn
	return 0
}

func (p tcpProc) Handover(ctx context.Context) {
	io.Copy(p.output, p.input)
}
