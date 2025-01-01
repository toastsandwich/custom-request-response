package customrequestresponsewriter

type HandlerFunc func(r *Request, w *Response) error
