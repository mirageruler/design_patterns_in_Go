### Definition
Chain of Responsibility is behavioral design pattern that allows passing request along the chain of potential handlers until one of them handles request.

The pattern allows multiple objects to handle the request without coupling sender class to the concrete classes of the receivers. The chain can be composed dynamically at runtime with any handler that follows a standard handler interface.

### Conceptual Example
Let’s look at the Chain of Responsibility pattern with the case of a hospital app. A hospital could have multiple departments such as:
- Reception
- Doctor
- Medical examination room
- Cashier

Whenever any patient arrives, they first get to Reception, then to Doctor, then to Medicine Room, and then to Cashier (and so on). The patient is being sent through a chain of departments, where each department sends the patient further down the chain once their function is completed.

The pattern is applicable when there are multiple candidates to process the same request. It is also useful when you don’t want the client to choose the receiver as there are multiple objects can handle the request. Another useful case is when you want to decouple the client from receivers — the client will only need to know the first element in the chain.

As in the example of the hospital, a patient first goes to the reception. Then, based upon a patient’s current status, reception sends up to the next handler in the chain.

### How to run
- from root run: `go run behavioral_patterns/chain_of_responsibility/*.go`