# Random thoughts and observations, Goland vs VScode

-   Debugging concurrent goRoutines
    -   When I was writing/testing [donechannel.go](pkg/concurrencypatterns/donechannel.go) I noticed that setting a breakpoint on func `getStatus` did not cause breakpoint to be hit "per" thread. Pressing F5 did not allow the IDE to automatically switch to the next thread that has not yet hit the breakpoint. (This is my guess) I need to test this code in GoLand, I suspect this could be a significant area of difference between the 2 IDE's.
