# Arhcitecture

## Current architecture

The project starts as a modular monolith.

Planned stacks:

- Go
- PostgreSQL
- HTTP API
- WebSocket
- gRPC
- NATS

## Planned evolution

The initial version will use a modular monolith.

As the system grows, selected modules may be extracted into independent services.

Planned services:

- Identity
- Game
- Character
- Realtime