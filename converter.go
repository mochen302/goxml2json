package xml2json

import (
	"bytes"
	"io"
)

// Convert converts the given XML document to JSON
func Convert(r io.Reader, ps ...plugin) (*bytes.Buffer, error) {
	return ConvertWithParam(r, 0, false, ps...)
}

// Convert converts the given XML document to JSON
func ConvertWithParam(r io.Reader, maxSkipLvl int, childrenAlwaysAsArray bool, ps ...plugin) (*bytes.Buffer, error) {
	// Decode XML document
	root := &Node{}
	err := NewDecoder(r, ps...).Decode(root)
	if err != nil {
		return nil, err
	}

	// Then encode it in JSON
	buf := new(bytes.Buffer)
	root.ChildrenAlwaysAsArray = childrenAlwaysAsArray

	e := NewEncoderWithSkipLv(buf, maxSkipLvl, ps...)
	err = e.Encode(root)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
