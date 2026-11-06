package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"testing"

	"github.com/lxkota95/odst/schemas/generated/odst_sbe"
)

// The order from the README's "NewOrderSingle Wire Format" table.
func readmeOrder(t testing.TB) *odst_sbe.NewOrderSingle {
	t.Helper()
	order := &odst_sbe.NewOrderSingle{
		Side:         odst_sbe.Side.BUY,
		OrderQty:     100,
		OrdType:      odst_sbe.OrdType.LIMIT,
		Price:        15000,
		TimeInForce:  odst_sbe.TimeInForce.DAY,
		TransactTime: 0x17B2E8F6A2C00000,
	}
	for _, f := range []struct {
		dst []byte
		s   string
	}{
		{order.ClOrdID[:], "CLIENT001"},
		{order.Account[:], "ACCT123"},
		{order.Symbol[:], "AAPL"},
		{order.Text[:], "Test order"},
	} {
		if err := SetText(f.dst, f.s); err != nil {
			t.Fatal(err)
		}
	}
	return order
}

func encode(t testing.TB, msg Message) []byte {
	t.Helper()
	var buf bytes.Buffer
	if err := Encode(&buf, msg); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestNewOrderSingleWireFormat(t *testing.T) {
	wire := encode(t, readmeOrder(t))

	if got, want := len(wire), HeaderSize+195; got != want {
		t.Fatalf("encoded %d bytes, want %d (8-byte header + 195-byte body)", got, want)
	}
	header := wire[:HeaderSize]
	if got, want := header, []byte{0, 195, 0, 100, 0, 1, 0, 0}; !bytes.Equal(got, want) {
		t.Errorf("header = % x, want % x (blockLength, templateId, schemaId, version)", got, want)
	}

	body := wire[HeaderSize:]
	for _, f := range []struct {
		name   string
		offset int
		want   []byte
	}{
		{"clOrdID", 0, []byte("CLIENT001\x00")},
		{"account", 20, []byte("ACCT123\x00")},
		{"symbol", 32, []byte("AAPL\x00")},
		{"side", 40, []byte{'1'}},
		{"orderQty", 41, binary.BigEndian.AppendUint64(nil, 100)},
		{"ordType", 49, []byte{'2'}},
		{"price", 50, binary.BigEndian.AppendUint64(nil, 15000)},
		{"timeInForce", 58, []byte{'0'}},
		{"transactTime", 59, binary.BigEndian.AppendUint64(nil, 0x17B2E8F6A2C00000)},
		{"text", 67, []byte("Test order\x00")},
	} {
		got := body[f.offset : f.offset+len(f.want)]
		if !bytes.Equal(got, f.want) {
			t.Errorf("%s at offset %d = % x, want % x", f.name, f.offset, got, f.want)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	ack := &odst_sbe.OrderAck{
		AckType:     odst_sbe.AckType.ORDER_ACCEPTED,
		OrdStatus:   odst_sbe.OrdStatus.NEW,
		Side:        odst_sbe.Side.SELL,
		OrderQty:    5,
		OrdType:     odst_sbe.OrdType.MARKET,
		TimeInForce: odst_sbe.TimeInForce.IMMEDIATE_OR_CANCEL,
	}
	if err := SetText(ack.OrderID[:], "ORD-1"); err != nil {
		t.Fatal(err)
	}

	for _, msg := range []Message{readmeOrder(t), ack} {
		got, err := Decode(bytes.NewReader(encode(t, msg)))
		if err != nil {
			t.Fatalf("decode %T: %v", msg, err)
		}
		if g, w := json.RawMessage(mustJSON(t, got)), json.RawMessage(mustJSON(t, msg)); !bytes.Equal(g, w) {
			t.Errorf("round trip of %T:\n got %s\nwant %s", msg, g, w)
		}
	}
}

func TestDecodeErrors(t *testing.T) {
	wire := encode(t, readmeOrder(t))

	unknown := bytes.Clone(wire)
	binary.BigEndian.PutUint16(unknown[2:], 999)
	if _, err := Decode(bytes.NewReader(unknown)); !errors.Is(err, ErrUnknownTemplate) {
		t.Errorf("unknown template: got %v, want ErrUnknownTemplate", err)
	}
	if _, err := Decode(bytes.NewReader(wire[:HeaderSize+10])); err == nil {
		t.Error("truncated body: got no error")
	}
	bad := bytes.Clone(wire)
	bad[HeaderSize+40] = 'X' // side: not BUY or SELL
	if _, err := Decode(bytes.NewReader(bad)); err == nil {
		t.Error("invalid side: got no error")
	}
}

func TestSetText(t *testing.T) {
	var symbol [8]byte
	if err := SetText(symbol[:], "TOOLONGSYM"); err == nil {
		t.Error("10 bytes into an 8-byte field: got no error")
	}
	if err := SetText(symbol[:], "MSFT"); err != nil || Text(symbol[:]) != "MSFT" {
		t.Errorf("SetText/Text = %q, %v", Text(symbol[:]), err)
	}
}

func mustJSON(t testing.TB, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// jsonOrder is the same order as plain JSON, for the size and speed comparison in the README.
type jsonOrder struct {
	ClOrdID      string `json:"clOrdID"`
	Account      string `json:"account"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	OrderQty     uint64 `json:"orderQty"`
	OrdType      string `json:"ordType"`
	Price        int64  `json:"price"`
	TimeInForce  string `json:"timeInForce"`
	TransactTime uint64 `json:"transactTime"`
	Text         string `json:"text"`
}

var readmeJSON = jsonOrder{"CLIENT001", "ACCT123", "AAPL", "1", 100, "2", 15000, "0",
	0x17B2E8F6A2C00000, "Test order"}

func TestSizes(t *testing.T) {
	t.Logf("SBE: %d bytes, JSON: %d bytes", len(encode(t, readmeOrder(t))), len(mustJSON(t, readmeJSON)))
}

func BenchmarkEncodeSBE(b *testing.B) {
	order := readmeOrder(b)
	var buf bytes.Buffer
	for b.Loop() {
		buf.Reset()
		if err := Encode(&buf, order); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecodeSBE(b *testing.B) {
	wire := encode(b, readmeOrder(b))
	r := bytes.NewReader(wire)
	for b.Loop() {
		r.Reset(wire)
		if _, err := Decode(r); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodeJSON(b *testing.B) {
	for b.Loop() {
		if _, err := json.Marshal(readmeJSON); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecodeJSON(b *testing.B) {
	data := mustJSON(b, readmeJSON)
	for b.Loop() {
		var o jsonOrder
		if err := json.Unmarshal(data, &o); err != nil {
			b.Fatal(err)
		}
	}
}
