# Go Concurrency Patterns

Rob Pike talks about go concurrency patterns on [2012 Google I/O talk](https://go.dev/talks/2012/concurrency.slide#1). Some useful discussed concurrency patterns are as following:  


- Generator Pattern: function that runs goroutine and returns channel.
- Fan-in (Multiplexing) Pattern: function that takes multiple channels and pipes to one channel, so that the returned channel receives both outputs.
- Daisychain Pattern: functions whose I/O are daisy-chained with channels together.
