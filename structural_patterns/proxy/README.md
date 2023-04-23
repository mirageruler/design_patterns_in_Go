### Definition
Proxy is a structural design pattern that provides an object that acts as a substitute for a real service object used by a client. A proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object.

### Conceptual Example
A web server such as Nginx can act as a proxy for your application server:

- It provides controlled access to your - application server.
- It can do rate limiting.
- It can do request caching.

### How to run
- from root run: `go run structural_pattern/proxy/*.go`