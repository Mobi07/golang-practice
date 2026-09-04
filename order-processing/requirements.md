# Concurrent Order Processing System

Go service that accepts orders and processes payment, inventory, and notification work concurrently. The project is for learning Go fundamentals and concurrency patterns, not for building a production payment system.

## Goal

An order should move through this lifecycle:

```text
POST /orders
     |
     v
Order service
     |
     v
Buffered job queue
     |
     +--> payment worker
     +--> inventory worker
     +--> notification worker
     |
     v
Order completed or failed
```

## Functional requirements

### Create an order

Provide `POST /orders`.

- Accept a customer identifier and one or more items.
- Each item must contain a product identifier and a quantity greater than zero.
- Create a unique order ID.
- Set the initial order status to `pending`.
- Add the order to the processing queue.
- Return the order ID and its current status.

### Retrieve an order

Provide `GET /orders/{id}`.

- Return the order's customer, items, status, and processing result.
- Return `404 Not Found` when the order ID does not exist.
- This endpoint must be safe to call while workers are updating orders.

### Process an order

For every accepted order, complete these three steps:

1. **Payment:** simulate charging the customer.
2. **Inventory:** simulate reserving every requested item.
3. **Notification:** simulate notifying the customer after successful processing.

Payment and inventory may run concurrently. Notification must run only after payment and inventory both succeed.

### Order states

Use the following states:

| Status | Meaning |
| --- | --- |
| `pending` | The order was accepted but has not started processing. |
| `processing` | At least one worker is processing the order. |
| `completed` | Payment, inventory, and notification all succeeded. |
| `failed` | A required step failed or timed out. |
| `cancelled` | The request or order context was cancelled. |

Once an order reaches `completed`, `failed`, or `cancelled`, it is final and must not be processed again.

## Concurrency requirements

- Use a buffered channel as the order job queue.
- Start a configurable number of worker goroutines when the application starts.
- Do not start a new unmanaged goroutine for each order without a way to wait for or stop it.
- Use a `sync.WaitGroup` to wait for workers during graceful shutdown.
- Use channels to return worker results to the order coordinator.
- Use `select` when waiting for a result, cancellation, or timeout.
- Protect shared mutable order state from data races.
- Run the project with the race detector while developing: `go test -race ./...`.

## Context, cancellation, and timeouts

- Pass a `context.Context` through the HTTP handler, service, and workers.
- If the client request is cancelled before processing finishes, stop work for that order when possible and mark it `cancelled`.
- Apply separate timeouts to payment and inventory operations.
- If one required operation fails or exceeds its timeout, cancel the other in-progress work and mark the order `failed`.
- Always release context resources by calling each derived context's cancel function.

## Data and synchronization requirements

- Define structs for `Order`, `OrderItem`, and a result or error record.
- Keep active orders in a concurrency-safe store.
- Start with a normal `map` protected by `sync.RWMutex` for typed, understandable access.
- Use a `sync.Mutex` when updating multiple related order fields together, such as status, error message, and timestamps.
- Use `sync.Map` only as an optional comparison exercise; it is not required for the main store.
- Use `sync.Once` for one-time setup such as configuration or logger initialization.
- Use atomic operations only for simple independent metrics, such as a completed-order counter.

## Interfaces

Create small interfaces for external behavior:

- A payment processor that charges an order.
- An inventory service that reserves order items.
- A notifier that sends an order notification.

Provide simple in-memory or simulated implementations. Do not connect to real payment, inventory, or email providers.

## Error handling

- Validate input before queuing an order.
- Return clear client errors for invalid requests.
- Return errors from worker operations instead of panicking for expected failures.
- Preserve useful context when returning an error, such as which order and operation failed.
- Ensure a failed order records a safe message that can be returned by `GET /orders/{id}`.
- Do not log the same error repeatedly at every layer.

## Graceful shutdown

When the application receives an interrupt signal:

1. Stop accepting new orders.
2. Cancel the application context.
3. Allow active workers a short, bounded time to finish or stop.
4. Wait for worker goroutines to exit.
5. Exit without leaking goroutines.

## Suggested learning milestones

1. Define the order structs, statuses, and in-memory store.
2. Implement `POST /orders` and `GET /orders/{id}` without concurrency.
3. Add a buffered job queue and one worker.
4. Add concurrent payment and inventory processing.
5. Add notification after both required steps succeed.
6. Add cancellation and per-operation timeouts.
7. Add safe shared-state access and run the race detector.
8. Add graceful shutdown and basic tests.

## Definition of done

- Valid orders are accepted and eventually become `completed`, `failed`, or `cancelled`.
- Invalid orders are rejected without entering the queue.
- `GET /orders/{id}` reports the latest safe order state.
- A payment or inventory failure prevents notification.
- A timeout or cancellation does not leave the order stuck in `processing`.
- Multiple requests can be handled concurrently without race-detector warnings.
- Application shutdown stops workers cleanly.
