# ODST - Order Distribution System Tool

An equity order message format in Go with [Simple Binary Encoding](https://github.com/aeron-io/simple-binary-encoding)
(SBE), built to learn how binary encodings work.

> **Status: concluded study (November 2026).** The goal was to understand binary encoding: schema
> design, fixed-size layouts, byte order, and how generated codecs read and write them. The schema,
> the generated codecs and a tested encode/decode layer are done; the TCP client and server below
> were never built and stay as ideas (see [What's next](#whats-next)).

## Overview

ODST demonstrates:
- **Simple Binary Encoding (SBE)**: a schema (`schemas/odst.xml`) compiled into Go codecs
- **FIX semantics**: field ids are FIX tag numbers (`ClOrdID` = 11, `Side` = 54, ...)
- **Fixed-size messages**: every field at a known offset, no parsing
- **Big-endian byte ordering** (network order), chosen in the schema

Try it:

```bash
go test -v ./...                                  # wire format, round trip, error cases
go test -run '^$' -bench . -benchmem ./pkg/protocol  # SBE vs JSON
```

## Architecture (planned, not built)

```
┌─────────────────┐         ┌─────────────────┐
│   ODST Client   │◄────────┤   ODST Server   │
│                 │   TCP   │                 │
│ ┌─────────────┐ │◄────────┤ ┌─────────────┐ │
│ │SBE Encoder/ │ │         │ │SBE Decoder/ │ │
│ │Decoder      │ │         │ │Encoder      │ │
│ └─────────────┘ │         │ └─────────────┘ │
│                 │         │                 │
│ ┌─────────────┐ │         │ ┌─────────────┐ │
│ │Order Manager│ │         │ │Order Engine │ │
│ └─────────────┘ │         │ └─────────────┘ │
└─────────────────┘         └─────────────────┘
```

## Message Flow (planned)

1. **Order Submission**: Client → Server (`NewOrderSingle`)
2. **Order Acknowledgment**: Server → Client (`OrderAck`)

Future iterations will add execution reports, market data, and order matching.

## Prerequisites

- **Go 1.24+**
- Only to regenerate the codecs after changing the schema: **Java 11+** and the **SBE Tool**
  (the generated code is committed)

## Regenerating the Codecs

### 1. Clone and Setup

```bash
git clone https://github.com/Lxkota95/odst.git
cd odst
mkdir -p schemas/generated
```

### 2. Install SBE Tool

```bash
# Option A: Build from source (recommended)
git clone https://github.com/aeron-io/simple-binary-encoding.git
cd simple-binary-encoding
./gradlew
cd ../

# Option B: Download from Maven Central
VERSION="1.32.0"  # Check for latest version
wget "https://repo1.maven.org/maven2/uk/co/real-logic/sbe-all/${VERSION}/sbe-all-${VERSION}.jar"
```

### 3. Generate Go Codecs

```bash
# Find the SBE JAR
SBE_JAR=$(find simple-binary-encoding/sbe-all/build/libs/ -name "sbe-all-*.jar" | head -1)

# Generate codecs
java --add-opens java.base/jdk.internal.misc=ALL-UNNAMED \
    -Dsbe.generate.ir=true \
    -Dsbe.target.language=golang \
    -Dsbe.output.dir=schemas/generated \
    -jar $SBE_JAR \
    schemas/odst.xml
```

### 4. Verify Generation

```bash
ls schemas/generated/odst_sbe/
# Should show: MessageHeader.go, NewOrderSingle.go, OrderAck.go, Side.go, etc.
```

## SBE Schema

### Message Types

| Message | ID | Description | Size |
|---------|----|-----------| -----|
| `NewOrderSingle` | 100 | Order submission | 195-byte body |
| `OrderAck` | 101 | Order acknowledgment | 213-byte body |

### Supported Order Types

- **Market Orders**: Execute immediately at current market price
- **Limit Orders**: Execute only at specified price or better

### Supported Order Sides

- **Buy** (1): Purchase securities
- **Sell** (2): Sell securities

### Time in Force Options

- **Day** (0): Valid for current trading session
- **Good Till Cancel** (1): Valid until explicitly canceled
- **Immediate or Cancel** (3): Execute immediately or cancel

## Wire Protocol

### Message Framing

```
[4-byte length prefix][8-byte SBE header][Message body]
```

The length prefix is for the planned TCP transport; `pkg/protocol` encodes the header and body.

### SBE Header (8 bytes)

- `blockLength` (2 bytes): Size of message body
- `templateId` (2 bytes): Message type (100=NewOrderSingle, 101=OrderAck)
- `schemaId` (2 bytes): Schema identifier (1)
- `version` (2 bytes): Schema version (0)

### Example: NewOrderSingle Wire Format

Offsets are within the body, after the 8-byte header; `TestNewOrderSingleWireFormat` checks
every row.

```
Offset  Field           Size    Example Value
------  -----           ----    -------------
0       clOrdID         20      "CLIENT001\0\0\0\0\0\0\0\0\0\0\0"
20      account         12      "ACCT123\0\0\0\0\0"
32      symbol          8       "AAPL\0\0\0\0"
40      side            1       0x31 ('1' = BUY)
41      orderQty        8       0x0000000000000064 (100)
49      ordType         1       0x32 ('2' = LIMIT)
50      price           8       0x0000000000003A98 (15000 ticks)
58      timeInForce     1       0x30 ('0' = DAY)
59      transactTime    8       0x17B2E8F6A2C00000 (nanoseconds)
67      text            128     "Test order\0\0\0..."
```

**Total Message Size**: 195 bytes (fixed)

## Code Examples

`pkg/protocol` wraps the generated codecs: `Encode` writes the 8-byte header and the body,
`Decode` reads the header and dispatches on its template id.

### Encoding a New Order

```go
order := &odst_sbe.NewOrderSingle{
    Side:         odst_sbe.Side.BUY,
    OrderQty:     100,
    OrdType:      odst_sbe.OrdType.LIMIT,
    Price:        15000, // $150.00 in ticks
    TimeInForce:  odst_sbe.TimeInForce.DAY,
    TransactTime: uint64(time.Now().UnixNano()),
}
protocol.SetText(order.ClOrdID[:], "CLIENT001") // NUL-padded fixed-size fields
protocol.SetText(order.Symbol[:], "AAPL")

var buf bytes.Buffer
err := protocol.Encode(&buf, order) // 8-byte header + 195-byte body
```

### Decoding a Message

```go
msg, err := protocol.Decode(conn)
switch m := msg.(type) {
case *odst_sbe.OrderAck:
    fmt.Println(protocol.Text(m.OrderID[:]), m.OrdStatus)
case *odst_sbe.NewOrderSingle:
    fmt.Println(protocol.Text(m.Symbol[:]), m.OrderQty)
}
```

## What I learned

Measured with `go test -bench . -benchmem ./pkg/protocol` (Go 1.27, Intel Core Ultra X7 358H),
for the order in the wire-format table above:

| | Size | Encode | Decode |
|---|---|---|---|
| SBE (generated Go codec) | 203 bytes (8 + 195) | ~330–420 ns, 5 allocs | ~450–540 ns, 5 allocs |
| JSON (`encoding/json`) | 186 bytes | ~1.1 µs, 3 allocs | ~1.0 µs, 1 alloc |

- **Fixed size isn't automatically small.** The 128-byte `text` field is mostly padding, so this
  SBE message is *bigger* than the JSON. SBE wins on predictability (every field at a known
  offset, no parsing), not on bytes; variable-length data belongs in SBE's `<data>` fields.
- **The codec style matters more than the format.** SBE's headline numbers (tens of nanoseconds)
  come from flyweight codecs that read and write fields in place in a buffer, as in Java and C++.
  The Go generator streams field by field through an `io.Writer` and checks enums with
  reflection, so it's only 2–3× faster than `encoding/json` here.
- **Byte order is a schema decision.** `byteOrder="bigEndian"` in the schema changes every
  multi-byte field; the test checks `orderQty`, `price` and `transactTime` byte by byte.
- **Generated range checks can disagree with the spec.** SBE pads a short `char` array with NUL
  bytes, but the Go generator's range check only allows printable ASCII, so it rejects every
  padded string. `pkg/protocol` validates enum fields itself instead.

## Project Structure

```
odst/
├── cmd/
│   ├── odst-server/          # placeholder (planned)
│   └── odst-client/          # placeholder (planned)
├── pkg/
│   ├── protocol/             # Encode / Decode over the generated codecs, with tests
│   └── network/              # placeholder (planned TCP transport)
└── schemas/
    ├── odst.xml              # SBE schema definition
    └── generated/
        └── odst_sbe/         # Generated Go codecs (do not edit)
```

## What's next

The study is concluded. If it's picked up again:
- [ ] TCP client and server, with the 4-byte length prefix framing described above
- [ ] Order validation and end-to-end tests
- [ ] Execution reports, cancel/replace, session logon/logout and heartbeats
- [ ] A flyweight-style codec (read fields in place from a `[]byte`) to compare with the generated one

## Acknowledgments

- [Simple Binary Encoding](https://github.com/aeron-io/simple-binary-encoding) - High-performance message encoding
- [FIX Protocol](https://www.fixtrading.org/) - Financial messaging standard
- Go team for excellent networking and concurrency primitives

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
