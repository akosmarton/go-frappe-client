package frappe

// Document represents a Frappe document
type Document map[string]interface{}

// NewDocument returns an initialized Document
func NewDocument() Document {
	return make(Document)
}

// Get returns the value of a field
func (d Document) Get(field string) interface{} {
	return d[field]
}

// Set sets the value of a field
func (d Document) Set(field string, value interface{}) {
	d[field] = value
}

// Delete deletes a field
func (d Document) Delete(field string) {
	delete(d, field)
}
