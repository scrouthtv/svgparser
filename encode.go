package svgparser

import (
	"encoding/xml"
	"io"
)

func (e *Element) Encode(w io.Writer) error {
	enc := xml.NewEncoder(w)

	err := e.write(enc)
	if err != nil {
		return err
	}

	return enc.Flush()
}

func (e *Element) write(enc *xml.Encoder) error {
	err := enc.EncodeToken(xml.StartElement{
		Name: e.Name,
		Attr: e.Attributes,
	})
	if err != nil {
		return err
	}

	if e.Content != "" {
		err = enc.EncodeToken(xml.CharData([]byte(e.Content)))
		if err != nil {
			return err
		}
	} else {
		for _, c := range e.Children {
			err = c.write(enc)
			if err != nil {
				return err
			}
		}
	}

	return enc.EncodeToken(xml.EndElement{
		Name: e.Name,
	})
}
