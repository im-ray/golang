# golang
Golang Learnings and Tools

#### Channels and goroutines
// Go uses goroutines for concurrency (preemptive multitasking), while Python uses async/await (cooperative multitasking). its bit confusing so clarify by understanding through below articles

References: 
Concurrency is not parallelism:
https://go.dev/blog/waza-talk

https://go.dev/doc/effective_go#goroutines

https://gobyexample.com/goroutines

https://docs.python.org/3/library/asyncio.html

https://realpython.com/async-io-python/

Learn About go Scheduler:: also try to visualise function call stack in scenarios
https://medium.com/@sanilkhurana7/understanding-the-go-scheduler-and-looking-at-how-it-works-e431a6daacf



--> Learn in Depth about channels
1. https://go.dev/ref/spec#Channel_types

2. https://go.dev/doc/effective_go#channels

3. Share Memory By Communicating: https://go.dev/blog/codelab-share

4. https://gobyexample.com/channels

5. Mastering Go Channesl: https://medium.com/@sumit-s/mastering-go-channels-from-beginner-to-pro-9c1eaba0da9e

6. Concurrency in Go (Book by Katherine Cox-Buday)
📌 One of the best books to master concurrency in Go, including channels.
https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/

7. Go Concurrency Patterns (Rob Pike's Talk at GopherCon)
📌 A must-watch talk explaining Go’s concurrency model and channel patterns.
Video: https://www.youtube.com/watch?v=f6kdp27TYZs

Slides: https://go.dev/talks/2012/concurrency.slide#1

8. 🔹 Advanced Topics
📌 Covers buffered/unbuffered channels, timeouts, and select statements.
https://blog.gopheracademy.com/series/advent-2014/

9. Concurrency Patterns in Go (Go.dev Blog)
📌 Discusses common concurrency patterns with channels.
https://go.dev/blog/pipelines

10. Practical Example: Worker Pool in Go
📌 Learn how to use channels to manage a worker pool efficiently.

11. https://medium.com/goturkiye/concurrency-in-go-channels-and-waitgroups-25dd43064d1

--> Learn more about blocking channels



