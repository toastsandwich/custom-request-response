package customrequestresponsewriter

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"strings"
)

type Request struct {
	Conn net.Conn // client addr
	// MetaData for the request like user
	MetaData   MetaData
	Service    string // what service do we want to use
	Body       []byte // content
	BodyLength int    // length of content
}

// Each time we want to write something in request
func (r *Request) WriteBody(b []byte) {
	r.Body = append(r.Body, b...)
	r.BodyLength += len(b)
}

// Return pointer to meta data don't know why using this, may be inspired by http package
func (r *Request) Meta() *MetaData {
	return &r.MetaData
}

func (r *Request) SerializeRequest() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buffer).Encode(r)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// to deserialize we will add conn of client and the incoming req.
func DeserializeRequest(from net.Conn, req []byte) (*Request, error) {
	buffer := bytes.NewBuffer(req)
	var r Request
	err := gob.NewDecoder(buffer).Decode(&r)
	if err != nil {
		return nil, err
	}
	r.Conn = from
	return &r, nil
}

func (r *Request) String() string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("from: %s", r.Conn.RemoteAddr().String()))
	metadata := *r.Meta()
	for k, v := range metadata {
		pair := fmt.Sprintf("%s: %s\n", k, v)
		builder.WriteString(pair)
	}
	builder.Write(r.Body)
	return builder.String()
}
