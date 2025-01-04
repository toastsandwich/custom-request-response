package customrequestresponsewriter

// Will store all the meta data request or response
type MetaData map[string]string

// Sets MetaData
func (m *MetaData) Set(k, v string) {
	(*m)[k] = v
}

// Gets a specific metadata
func (m *MetaData) Get(key string) string {
	if v, ok := (*m)[key]; ok {
		return v
	}
	return ""
}

// Closer
func (m *MetaData) Close() {
	*m = nil
}
