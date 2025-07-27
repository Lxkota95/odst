# ODST - Order Distribution System Tool

## SBE Architecture Design

### Overview
ODST uses Simple Binary Encoding (SBE) for high-performance, low-latency message serialization between trading clients and order management servers. SBE provides deterministic encoding/decoding with minimal CPU overhead.

### Core Components

#### 1. Schema Definition Layer
```
schemas/
├── fix_messages.xml          # FIX protocol message definitions
├── odst_extensions.xml       # ODST-specific extensions
└── generated/
    ├── go/                   # Generated Go codecs
    │   ├── new_order_single.go
    │   ├── execution_report.go
    │   ├── order_cancel.go
    │   └── market_data.go
    └── sbe-tool-output/      # Raw SBE generator output
```

#### 2. Message Types (FIX-compliant)
- **NewOrderSingle (D)** - Client order submission
- **ExecutionReport (8)** - Order status updates
- **OrderCancelRequest (F)** - Cancel existing orders
- **OrderCancelReject (9)** - Cancel rejection
- **MarketDataRequest (V)** - Market data subscription
- **MarketDataSnapshot (W)** - Market data updates

#### 3. Client-Server Architecture

```
┌─────────────────┐         ┌─────────────────┐
│   ODST Client   │◄────────┤   ODST Server   │
│                 │         │                 │
│ ┌─────────────┐ │   TCP   │ ┌─────────────┐ │
│ │SBE Encoder/ │ │◄────────┤ │SBE Decoder/ │ │
│ │Decoder      │ │         │ │Encoder      │ │
│ └─────────────┘ │         │ └─────────────┘ │
│                 │         │                 │
│ ┌─────────────┐ │         │ ┌─────────────┐ │
│ │Order Manager│ │         │ │Order Engine │ │
│ └─────────────┘ │         │ └─────────────┘ │
│                 │         │                 │
│ ┌─────────────┐ │         │ ┌─────────────┐ │
│ │FIX Protocol │ │         │ │Risk Manager │ │
│ │Validator    │ │         │ │             │ │
│ └─────────────┘ │         │ └─────────────┘ │
└─────────────────┘         └─────────────────┘
```

### 4. SBE Integration Workflow

#### Schema Generation Process
```bash
# 1. Define FIX messages in XML schema
# 2. Generate Go codecs using SBE tool
java -jar sbe-all-1.x.x.jar schemas/fix_messages.xml

# 3. Generated files provide:
#    - Message encoding/decoding functions
#    - Field accessors with bounds checking  
#    - Buffer management utilities
```

#### Message Flow
```
Client Side:
1. Create order struct
2. Validate FIX compliance
3. Encode using SBE codec
4. Send binary message over TCP
5. Receive SBE-encoded response
6. Decode execution report

Server Side:
1. Receive SBE binary message
2. Decode using generated codec
3. Validate order parameters
4. Process through order engine
5. Generate execution report
6. Encode response with SBE
7. Send back to client
```

### 5. Go Package Structure

```
cmd/
├── odst-server/
│   └── main.go
└── odst-client/
    └── main.go

pkg/
├── sbe/
│   ├── codecs/              # Generated SBE codecs
│   ├── buffer/              # Buffer management
│   └── validator/           # Message validation
├── fix/
│   ├── protocol/            # FIX protocol logic
│   ├── fields/              # FIX field definitions
│   └── messages/            # FIX message builders
├── network/
│   ├── tcp/                 # TCP connection handling
│   ├── session/             # Session management
│   └── framing/             # Message framing
└── engine/
    ├── orders/              # Order management
    ├── matching/            # Order matching logic
    └── risk/                # Risk management
```

### 6. Key Benefits of SBE Integration

- **Performance**: Sub-microsecond encoding/decoding
- **Deterministic**: No garbage collection pressure
- **Compact**: Minimal wire format overhead
- **Type Safety**: Generated Go structs with compile-time validation
- **FIX Compliance**: Native support for FIX protocol semantics
- **Schema Evolution**: Backward/forward compatibility support

### 7. Configuration

```yaml
# odst-config.yaml
sbe:
  schema_path: "./schemas/fix_messages.xml"
  generated_path: "./pkg/sbe/codecs"
  buffer_size: 8192
  
network:
  listen_addr: ":8080"
  max_connections: 1000
  read_timeout: "5s"
  write_timeout: "5s"

fix:
  version: "FIX.4.4"
  sender_comp_id: "ODST_SERVER"
  heartbeat_interval: 30
```

### 8. Next Implementation Steps

1. **Schema Definition**: Create FIX message XML schemas
2. **Code Generation**: Set up SBE code generation pipeline
3. **Network Layer**: Implement TCP server/client with message framing
4. **Message Handling**: Build encoding/decoding wrapper functions
5. **FIX Protocol**: Implement FIX session management
6. **Order Engine**: Create basic order matching logic
7. **Testing**: Unit tests for codec performance and correctness