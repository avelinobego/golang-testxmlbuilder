package element

import "bytes"

func New(root Element) *Body {
	return &Body{root: root, buffer: &bytes.Buffer{}}
}

type Body struct {
	root   Element
	buffer *bytes.Buffer
}
