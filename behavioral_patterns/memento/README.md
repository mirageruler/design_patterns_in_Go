### Definition
Memento is a behavioral design pattern that allows making snapshots of an object’s state and restoring it in future.

The Memento doesn’t compromise the internal structure of the object it works with, as well as data kept inside the snapshots.

### Conceptual Example
The Memento pattern lets us save snapshots for an object’s state. You can use these snapshots to revert the object to the previous state. It’s handy when you need to implement undo-redo operations on an object.

### How to run
- from root run: `go run behavioral_patterns/memento/*.go`