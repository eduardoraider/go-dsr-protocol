# DSR Implementation in Golang

## Overview

This is a simplified implementation of the Dynamic Source Routing (DSR) protocol in Golang. The DSR protocol is a reactive routing protocol that allows packets to contain the complete path to the destination within their headers.

## Features

- Sends Route Request (RREQ) messages.
- Propagates RREQ messages through intermediate nodes, building the path.
- Sends Route Reply (RREP) messages with the complete path.
- Updates route cache based on received RREP messages.

## Usage

### Prerequisites

- Golang installed on your machine.

### Running the Code

1. Clone the repository or copy the code to your local machine.
2. Save the code in a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the code using the following command:

```sh
go run main.go
```

### Code Explanation

#### `RouteRequest` Struct

Defines the structure for Route Request (RREQ) messages.

- `Source`: The source node of the request.
- `Destination`: The destination node of the request.
- `Path`: A slice to record the path taken by the request.
- `TTL`: Time to live for the request, to limit the number of hops.

#### `RouteReply` Struct

Defines the structure for Route Reply (RREP) messages.

- `Source`: The source node of the reply (which is the destination of the original request).
- `Destination`: The destination node of the reply (which is the source of the original request).
- `Path`: The complete path from the source to the destination.

#### Functions

- `sendRREQ(src, dest string)`: Initiates the sending of a Route Request.
- `propagateRREQ(rreq RouteRequest)`: Simulates the propagation of the RREQ message through intermediate nodes, adding each node to the path.
- `sendRREP(rreq RouteRequest)`: Sends a Route Reply with the complete path in response to a RREQ message.
- `updateRouteCache(rrep RouteReply)`: Updates the route cache based on the received RREP message.

## Example Output

When you run the code, you will see output indicating the sending and propagation of RREQ messages, as well as the sending of RREP messages and updates to the route cache.

```sh
Sending RREQ: {Source: NodeA, Destination: NodeB, Path: [NodeA], TTL: 5}
Propagating RREQ: {Source: NodeA, Destination: NodeB, Path: [NodeA, IntermediateNode], TTL: 4}
Sending RREP: {Source: NodeB, Destination: NodeA, Path: [NodeA, IntermediateNode, NodeB]}
Updated Route Cache: map[NodeA:[NodeA IntermediateNode NodeB]]
```

## Notes

This is a simplified implementation and does not handle many practical aspects of DSR, such as collision handling, neighbor table maintenance, or route error messages. It is intended to understand the basic operation of the DSR protocol.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**Eduardo Raider** - Software Engineer