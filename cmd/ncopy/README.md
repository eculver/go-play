# ncopy

I'm calling this `ncopy` for lack of a better name. The setup is this:

Assume you have a function, `copyFile(src, dst)`, that copies files from one machine
(`src`) to another machine (`dst`). Write an algorithm to efficiently copy a file
to N machines. Hint: must be done non-sequentially.

# Solution

For the sake of making this demonstrable, we're going to write `copyFile` as a function
that sleeps for a random amount of time between 3 and 15 seconds to simulate some blocking
operation resembling a real file transfer.

Obviously, the simplest solution is to just roll through the list of destinations sequentially calling `copyFile`,
but this would take a long time and is pretty inefficient. Also, if we have thousands of machines, it
will take a prohibitive amount of time to complete.

So, instead, we're going to start by focusing on two optimizations:

1) do M `copyFile` operations concurrently from any given source
2) as soon as a machine receives the file, use it to send M operations to any remaining hosts

It's as though there are two buckets, one for sources and one for destinations:

![Whiteboard](https://github.com/eculver/go-play/blob/ncopy/cmd/ncopy/ncopy.jpg)

To achieve the concurrency in Go, I decided to use channels to coordinate the effort. There
are two that we care about: `seeds` and `sends`. The `seeds` channel is populated with M source values
to represent that there are M candidates to send from. The result doesn't amount to that much code.

# Alternative Solutions

I thought about managing a heap-based priority queue where each node in the queue represented a source with its priority
being tied to the number of in-flight `copyFile` operations but it got to the point where it seemed like overkill pretty quickly
just to get something working. Plus, by treating the number of in-flight operations as priority in the heap means that every time a `copyFile` finishes, we
would need to "Fix" the heap which is `O(log n)` that is not necessary when using channels. I sure there's a point at which having
a structure back prioritization, but for this example, it didn't make sense. For example, if each transfer takes `O(seconds)` in the
best case and `O(minutes)` in the worst case, it would probably make sense to implement some intelligence around how to avoid that worst case.

I'm sure there are other solutions too that I didn't consider. I'm sure that in a real-world scenario, we might want
to consider evaluating and isolating certain nodes in the case where they may be impacted by reduced I/O meaning, if a node is 
subject to relative network congestion, it would probably make sense to avoid using it as a seed node.

There's also the situation in the real-world where not every `copyFile` operation is going to complete successfully so retries should be considered.
In the face of needing to retry, we'd probably want to know if it was related to the source node and potentially isolate that node
from future copes or maybe even put it in some sort of "penalty box" to account for transient failures.

# Results

Generally, the more nodes there are, the greater paralellism can be achived for basically just the cost of keeping
integers in a channel. There aren't any expensive seeks or traversals at play since there is no prioritization. So, we're
looking at `O(n*m)` space complexity since the implementation eventually puts `m` copies of sources into the seed channel
and an `O(n)` time complexity to iterate over all destinations.
