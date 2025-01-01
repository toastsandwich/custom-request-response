package customrequestresponsewriter

import (
	"bytes"
	"encoding/gob"
	"errors"
	"net"
)

// To send a specific response for a request
// Response :
// StatusCode
// Key: Value
//
//	:     :
//
// Body
// BodyLength

type Response struct {
	// For sending the response
	Conn net.Conn `gob:"-"` // exclude from gob
	// Meta data will store all the required info like keys,
	// users,
	MetaData   MetaData
	StatusCode int
	// Body will have all the needful stuff inform
	// plan is to make them key value pair
	Body       []byte
	BodyLength int
}

// send the response to respective conn
func (r *Response) Send() error {
	defer r.Close() // Ensure cleanup
	if r.Conn == nil {
		return errors.New("connection not initialized")
	}
	return nil
}

// write data to body
func (r *Response) Write(data []byte) {
	r.Body = data
	r.BodyLength = len(data)
}

// get metadata
func (r *Response) Meta() *MetaData {
	return &r.MetaData
}

// status ok
func (r *Response) Status(statusCode int) {
	r.StatusCode = statusCode
}

// Once the response is sent close the Response
func (r *Response) Close() error {
	r.MetaData.Close()
	return r.Conn.Close()
}

func DeserializeResponse(b []byte) (*Response, error) {
	buffer := bytes.NewBuffer(b)
	var r Response
	err := gob.NewDecoder(buffer).Decode(&r)
	return &r, err
}

func (r *Response) SerializeResponse() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buffer).Encode(r)
	return buffer.Bytes(), err
}
