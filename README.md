### 1. Distributed, WAL-Backed Message Broker

Build a fault-tolerant message queue from scratch (a lightweight Kafka or RabbitMQ clone). This forces you to handle low-level disk I/O, consensus, and custom network protocols.

* **Architecture:** Implement an append-only Write-Ahead Log (WAL) using an embedded key-value store for durable storage, managing cluster state with the Raft consensus algorithm.
* **Design Challenges:** Ensuring zero-data-loss during node crashes, managing consumer offsets, achieving high-throughput zero-copy reads, and handling split-brain network partitions.
* **Tech Stack:** Go (custom TCP server), gRPC/Protobufs for inter-node communication, and embedded KV stores.

```
                GopherLog
                   │
       ┌───────────┴───────────┐
       │                       │
   Client TCP              Inter-node
       │                    gRPC/Proto
       ▼                       │
   Request Router              ▼
       │                      Raft
   ┌───┴────┐                  │
   ▼        ▼                  ▼
Produce   Fetch             Replicated
   │        │                 State
   └───┬────┘
       ▼
   Partitions
       │
       ▼
      WAL
       │
       ▼
 Embedded KV
```

## arch

```
  ┌───────────────────────────────────────────────────────────────────
                           │              CLIENT TIER                │
                           │   Producers & Consumers (Native TCP)    │
                           └────────────────────┬────────────────────┘
                                                │ Port 9092 (Data Plane)
════════════════════════════════════════════════╪═════════════════════════════════════════════
BROKER NODE                                     │
                                                ▼
┌────────────────────────────────────────────────────────────────────────────────────────────┐
│ DATA PLANE (Client Ingress / Egress)                                                      │
│                                                                                            │
│   ┌────────────────────────────────────────────────────────────────────────────────────┐   │
│   │ TCP Multiplexer & Framer (`net.Listener`)                                          │   │
│   │   ├── Goroutine-per-Connection Handler                                             │   │
│   │   ├── Packet Frame Decoder: Length-prefixed binary frames                          │   │
│   │   └── Request Router: PRODUCE, FETCH, OFFSET_COMMIT, JOIN_GROUP                    │   │
│   └───────────────┬────────────────────────────────────────────────────┬───────────────┘   │
│                   │                                                    │                   │
│                   ▼                                                    ▼                   │
│   ┌────────────────────────────────┐                  ┌────────────────────────────────┐   │
│   │ Partition Log Appender         │                  │ Zero-Copy Fetch Engine         │   │
│   │ (Write Path)                   │                  │ (Read Path)                    │   │
│   │   ├── Active Segment Writer    │                  │   ├── Sparse Index Binary-Srch │   │
│   │   ├── In-Memory Index Builder  │                  │   ├── File Segment Reader      │   │
│   │   └── Cond Broadcast Notifier  │                  │   └── `net.TCPConn.ReadFrom()` │   │
│   └───────────────┬────────────────┘                  └────────────────▲───────────────┘   │
│                   │                                                    │                   │
│                   │ Dispatches Writes                                  │ Reads Data        │
│                   ▼                                                    │ Directly          │
├───────────────────┼────────────────────────────────────────────────────┼───────────────────┤
│ STORAGE SUBSYSTEM │                                                    │                   │
│                   │                                                    │                   │
│   ┌───────────────▼────────────────┐                  ┌────────────────┴───────────────┐   │
│   │ Embedded KV Store (Pebble DB)  │                  │ Raw Append-Only Segments       │   │
│   │   ├── Raft Stable/Log Store    │                  │   ├── orders-0/000000.log      │   │
│   │   ├── Consumer Offsets         │                  │   ├── orders-0/000000.index    │   │
│   │   └── Metadata (HighWatermark) │                  │   └── (Direct Linux PageCache) │   │
│   └───────────────▲────────────────┘                  └────────────────────────────────┘   │
│                   │                                                                        │
├───────────────────┼────────────────────────────────────────────────────────────────────────┤
│ CONTROL PLANE     │                                                                        │
│                   ▼                                                                        │
│   ┌────────────────────────────────────────────────────────────────────────────────────┐   │
│   │ Consensus Engine (`hashicorp/raft`)                                                │   │
│   │   ├── Finite State Machine (FSM) Implementation                                    │   │
│   │   ├── Raft Log Replication Coordinator                                             │   │
│   │   └── Cluster Membership & Heartbeat Tracker                                       │   │
│   └───────────────────────────────────────────┬────────────────────────────────────────┘   │
│                                               │                                            │
└───────────────────────────────────────────────┼────────────────────────────────────────────┘
                                                │ Port 9093 (Control Plane)
════════════════════════════════════════════════╪═════════════════════════════════════════════
                                                │ gRPC Inter-Node Gossip / AppendEntries
                                                ▼
                           ┌─────────────────────────────────────────┐
                           │          PEER BROKER NODES              │
                           │   Node 2 (:9093)  │  Node 3 (:9093)     │
                           └───────────────────┴─────────────────────┘
```

# further optimizations :

- per-partition sparse offset indexes as a later storage optimization.