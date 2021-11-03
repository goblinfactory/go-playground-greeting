# backpressure - quotes and extracts

## references

[Backpressure explained — the resisted flow of data through software](https://medium.com/@jayphelps/backpressure-explained-the-flow-of-data-through-software-2350b3e77ce7) : _Jay Phelps (Feb 1, 2019)_

In his article Jay states that we have 4 possible options;

1. Control the producer : (slow down/speed up is decided by consumer)
2. Buffer : (accumulate incoming data spikes temporarily)
3. Drop : (sample a percentage of the incoming data)
4. Ignore : (Ignore the backpressure)

[5 proven patterns for resilient software architecture design](https://searchapparchitecture.techtarget.com/tip/5-proven-patterns-for-resilient-software-architecture-design) : _Pryank Gupta (10 Jun 2021)_

Key points

-   Back pressured mechanisms ought to cascade across multiple nodes.
-   Components should be able to push workload back through upstream producers.
-   As services recover, they should allow "backpressured" calls to reach it again.

[Applying Bulkheads And Backpressure using MicroProfile (Video)](https://blog.sebastian-daschner.com/entries/bulkheads-backpressure-microprofile) _Sebastian Daschner (3 Mar 2019)_

Extract

_"The whole purpose of defining and constraining multiple resource pools is to being predictable to whether our service is able to respond within a given time. If that is not the case, we want the service to immediately deny the request without unnecessarily consuming too much of the client’s time."_
