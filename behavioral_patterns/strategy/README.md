### Definition
Strategy is a behavioral design pattern that turns a set of behaviors into objects and makes them interchangeable inside original context object.

The original object, called context, holds a reference to a strategy object. The context delegates executing the behavior to the linked strategy object. In order to change the way the context performs its work, other objects may replace the currently linked strategy object with another one.

### Conceptual Example
Suppose you are building an In-Memory-Cache. Since it’s in memory, it has a limited size. Whenever it reaches its maximum size, some entries have to be evicted to free-up space. This can happen via several algorithms. Some of the popular algorithms are:
- Least Recently Used (LRU): remove an entry that has been used least recently.
- First In, First Out (FIFO): remove an entry that was created first.
- Least Frequently Used (LFU): remove an entry that was least frequently used.

The problem is how to decouple our cache class from these algorithms so that we can change the algorithm at run time. Also, the cache class should not change when a new algorithm is being added.

This is were Strategy pattern comes into the picture. It suggests creating a family of the algorithm with each algorithm having its own class. Each of these classes follows the same interface, and this makes the algorithm interchangeable within the family. Let’s say the common interface name is evictionAlgo.

Now our main cache class will embed the evictionAlgo interface. Instead of implementing all types of eviction algorithms in itself, our cache class will delegate the execution to the evictionAlgo interface. Since evictionAlgo is an interface, we can change the algorithm in run time to either LRU, FIFO, LFU without changing the cache class.

### How to run
- from root run: `go run behavioral_patterns/strategy/*.go`