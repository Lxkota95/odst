// Package protocol encodes and decodes ODST messages: an 8-byte SBE message header followed by the
// fixed-size message body, both big-endian (see schemas/odst.xml).
package protocol

import (
	"errors"
	"fmt"
	"io"

	"github.com/lxkota95/odst/schemas/generated/odst_sbe"
)

// HeaderSize is the encoded size of the SBE message header.
const HeaderSize = 8

// ErrUnknownTemplate is returned by Decode for a message type this schema doesn't define.
var ErrUnknownTemplate = errors.New("unknown template id")

// Message is a message body generated from the schema.
type Message interface {
	SbeBlockLength() uint16
	SbeTemplateId() uint16
	SbeSchemaId() uint16
	SbeSchemaVersion() uint16
	Encode(m *odst_sbe.SbeGoMarshaller, w io.Writer, doRangeCheck bool) error
}

// Encode checks msg's enum fields, then writes its header and body.
//
// The generated range check isn't used: it rejects the NUL bytes that pad short char fields
// (it only allows printable ASCII), although NUL padding is how SBE stores a shorter string.
func Encode(w io.Writer, msg Message) error {
	if err := checkEnums(msg); err != nil {
		return err
	}
	m := odst_sbe.NewSbeGoMarshaller()
	header := odst_sbe.MessageHeader{
		BlockLength: msg.SbeBlockLength(),
		TemplateId:  msg.SbeTemplateId(),
		SchemaId:    msg.SbeSchemaId(),
		Version:     msg.SbeSchemaVersion(),
	}
	if err := header.Encode(m, w); err != nil {
		return fmt.Errorf("encode header: %w", err)
	}
	if err := msg.Encode(m, w, false); err != nil {
		return fmt.Errorf("encode %T: %w", msg, err)
	}
	return nil
}

// Decode reads one message: the header says which type follows and how long its body is.
// It returns *odst_sbe.NewOrderSingle or *odst_sbe.OrderAck.
func Decode(r io.Reader) (any, error) {
	m := odst_sbe.NewSbeGoMarshaller()
	var header odst_sbe.MessageHeader
	if err := header.Decode(m, r, 0); err != nil {
		return nil, fmt.Errorf("decode header: %w", err)
	}
	switch header.TemplateId {
	case (*odst_sbe.NewOrderSingle)(nil).SbeTemplateId():
		var order odst_sbe.NewOrderSingle
		if err := order.Decode(m, r, header.Version, header.BlockLength, false); err != nil {
			return nil, fmt.Errorf("decode NewOrderSingle: %w", err)
		}
		return &order, checkEnums(&order)
	case (*odst_sbe.OrderAck)(nil).SbeTemplateId():
		var ack odst_sbe.OrderAck
		if err := ack.Decode(m, r, header.Version, header.BlockLength, false); err != nil {
			return nil, fmt.Errorf("decode OrderAck: %w", err)
		}
		return &ack, checkEnums(&ack)
	default:
		return nil, fmt.Errorf("%w: %d", ErrUnknownTemplate, header.TemplateId)
	}
}

type enum interface {
	RangeCheck(actingVersion, schemaVersion uint16) error
}

// checkEnums rejects enum fields that hold none of the schema's valid values.
func checkEnums(msg Message) error {
	var enums []enum
	switch m := msg.(type) {
	case *odst_sbe.NewOrderSingle:
		enums = []enum{m.Side, m.OrdType, m.TimeInForce}
	case *odst_sbe.OrderAck:
		enums = []enum{m.AckType, m.OrdStatus, m.Side, m.OrdType, m.TimeInForce}
	}
	v := msg.SbeSchemaVersion()
	for _, e := range enums {
		if err := e.RangeCheck(v, v); err != nil {
			return fmt.Errorf("%T: %w", msg, err)
		}
	}
	return nil
}

// SetText copies s into a fixed-size char field, padding with NUL bytes. It fails if s doesn't fit.
func SetText(dst []byte, s string) error {
	if len(s) > len(dst) {
		return fmt.Errorf("%q is longer than %d bytes", s, len(dst))
	}
	clear(dst)
	copy(dst, s)
	return nil
}

// Text returns a fixed-size char field without its NUL padding.
func Text(field []byte) string {
	for i, b := range field {
		if b == 0 {
			return string(field[:i])
		}
	}
	return string(field)
}
