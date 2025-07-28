# ODST - Order Distribution System Tool

A high-performance equity order management system built with Go and Simple Binary Encoding (SBE) for low-latency trading applications.

## Overview

ODST (Order Distribution System Tool) is a proof-of-concept trading system that demonstrates:
- **Simple Binary Encoding (SBE)** for ultra-fast message serialization
- **FIX Protocol compliance** for industry standardization  
- **Client-server architecture** with TCP transport
- **Big-endian byte ordering** for network compatibility

## Architecture

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

## Message Flow (Proof of Concept)

1. **Order Submission**: Client → Server (`NewOrderSingle`)
2. **Order Acknowledgment**: Server → Client (`OrderAck`)

Future iterations will add execution reports, market data, and order matching.

## Prerequisites

- **Go 1.19+** (for development)
- **Java 11+** (for SBE code generation)
- **SBE Tool** (Simple Binary Encoding compiler)

## Quick Start

### 1. Clone and Setup

```bash
git clone <your-repo-url>
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
    schemas/odst-messages.xml
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
| `NewOrderSingle` | 100 | Order submission | ~195 bytes |
| `OrderAck` | 101 | Order acknowledgment | ~195 bytes |

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

### SBE Header (8 bytes)

- `blockLength` (2 bytes): Size of message body
- `templateId` (2 bytes): Message type (100=NewOrderSingle, 101=OrderAck)
- `schemaId` (2 bytes): Schema identifier (1)
- `version` (2 bytes): Schema version (0)

### Example: NewOrderSingle Wire Format

```
Offset  Field           Size    Example Value
------  -----           ----    -------------
0       clOrdID         20      "CLIENT001\0\0\0\0\0\0\0\0\0\0"
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

### Encoding a New Order

```go
package main

import (
    "encoding/binary"
    "time"
    "your-module/schemas/generated/odst_sbe"
)

func encodeNewOrder() []byte {
    // Create buffer
    buffer := make([]byte, 1024)
    offset := 0
    
    // Encode SBE header
    header := odst_sbe.MessageHeader{}
    header.WrapForEncode(buffer, uint64(offset), uint64(len(buffer)))
    header.SetBlockLength(195) // NewOrderSingle size
    header.SetTemplateId(100)  // NewOrderSingle template
    header.SetSchemaId(1)
    header.SetVersion(0)
    offset += int(header.Size())
    
    // Encode message body
    order := odst_sbe.NewOrderSingle{}
    order.WrapForEncode(buffer, uint64(offset), uint64(len(buffer)))
    order.SetClOrdID("CLIENT001")
    order.SetAccount("ACCT123")
    order.SetSymbol("AAPL")
    order.SetSide(odst_sbe.Side.BUY)
    order.SetOrderQty(100)
    order.SetOrdType(odst_sbe.OrdType.LIMIT)
    order.SetPrice(15000) // $150.00 in ticks
    order.SetTimeInForce(odst_sbe.TimeInForce.DAY)
    order.SetTransactTime(uint64(time.Now().UnixNano()))
    order.SetText("Test limit order")
    
    totalSize := int(header.Size()) + int(order.Size())
    return buffer[:totalSize]
}
```

### Decoding an Order Acknowledgment

```go
func decodeOrderAck(data []byte) {
    offset := 0
    
    // Decode SBE header
    header := odst_sbe.MessageHeader{}
    header.WrapForDecode(data, uint64(offset), uint64(len(data)))
    
    templateId := header.TemplateId()
    blockLength := header.BlockLength()
    offset += int(header.Size())
    
    if templateId == 101 { // OrderAck
        ack := odst_sbe.OrderAck{}
        ack.WrapForDecode(data, uint64(offset), uint64(len(data)))
        
        fmt.Printf("Order %s: %s\n", 
            ack.ClOrdID(), 
            ack.AckType())
        fmt.Printf("Status: %s\n", ack.OrdStatus())
        fmt.Printf("Text: %s\n", ack.Text())
    }
}
```

## Performance Characteristics

### Message Size Comparison

| Format | NewOrderSingle Size |
|--------|-------------------|
| JSON | ~400-600 bytes |
| FIX (text) | ~300-400 bytes |
| SBE | ~195 bytes |

### Encoding Performance

| Operation | Latency |
|-----------|---------|
| SBE Encoding | ~50-100ns |
| SBE Decoding | ~30-50ns |
| JSON Encoding | ~1-2μs |
| FIX Text Encoding | ~800ns-1μs |

## Project Structure

```
odst/
├── cmd/
│   ├── odst-server/          # Server application
│   └── odst-client/          # Client application
├── pkg/
│   ├── protocol/             # Message handling
│   ├── network/              # TCP transport
│   └── validation/           # Order validation
├── schemas/
│   ├── odst-messages.xml     # SBE schema definition
│   └── generated/
│       └── odst_sbe/         # Generated Go codecs
├── examples/                 # Usage examples
├── docs/                     # Documentation
└── README.md
```

## Development Roadmap

### Phase 1: Proof of Concept ✅
- [x] SBE schema definition
- [x] Code generation pipeline
- [x] Basic message encoding/decoding
- [ ] TCP client/server implementation
- [ ] Order validation logic
- [ ] End-to-end testing

### Phase 2: Core Features
- [ ] Execution reports with fill details
- [ ] Order cancellation and replacement
- [ ] Session management (logon/logout)
- [ ] Heartbeat mechanism
- [ ] Error handling and recovery

### Phase 3: Advanced Features
- [ ] Market data distribution
- [ ] Order book state management
- [ ] Multi-symbol support
- [ ] Risk management controls
- [ ] Performance monitoring

## Configuration

### Network Settings
- **Default Port**: 8080
- **Max Connections**: 1000
- **Read Timeout**: 5s
- **Write Timeout**: 5s

### Message Settings
- **Buffer Size**: 8192 bytes
- **Max Message Size**: 4096 bytes
- **Heartbeat Interval**: 30 seconds

## Testing

```bash
# Run unit tests
go test ./...

# Run benchmarks
go test -bench=. ./pkg/protocol/

# Test codec generation
make generate-codecs

# Integration tests
make test-integration
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Simple Binary Encoding](https://github.com/aeron-io/simple-binary-encoding) - High-performance message encoding
- [FIX Protocol](https://www.fixtrading.org/) - Financial messaging standard
- Go team for excellent networking and concurrency primitives