package customrequestresponsewriter

import (
	"bytes"
	"encoding/gob"
)

type Request struct {
	MetaData   MetaData
	Service    string
	Body       []byte
	BodyLength int
}

// Each time we want to write something in request
func (r *Request) WriteBody(b []byte) {
	r.Body = append(r.Body, b...)
	r.BodyLength += len(b)
}

// Return pointer to meta data dontknow why using this, may be inspired by http package
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

func DeserializeRequest(req []byte) (*Request, error) {
	buffer := bytes.NewBuffer(req)
	var r Request
	err := gob.NewDecoder(buffer).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
