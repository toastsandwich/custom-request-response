package customrequestresponsewriter

type MetaData map[string]string

func (m *MetaData) Set(k, v string) {
	if *m == nil {
		*m = make(MetaData)
	}
	(*m)[k] = v
}

func (m *MetaData) Get(key string) string {
	if v, ok := (*m)[key]; ok {
		return v
	}
	return ""
}

func (m *MetaData) Close() {
	*m = nil
}
