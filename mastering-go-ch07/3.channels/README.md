**Channels**

A **channel** is a communication mechanism that, among other things, allows
goroutines to exchange data.

1. Firstly, each channel allows the exchange of a particular
data type, which is also called the element type of the channel,
2. secondly, for a channel to operate properly, you need someone to receive what is sent via the
channel.
3. You should declare a new channel using *make()* and the **chan** keyword
*(make(chan int))*,
4. and you can close a channel using the *close()* function.
5. You can declare the size of a channel by writing something like *make(chan int, 1)*.

A **pipeline** is a virtual method for connecting goroutines and channels so that the
output of one goroutine becomes the input of another goroutine using channels
to transfer your data. The benefits that you get from using pipelines are that:

1. there is a constant data flow in your program, as no goroutine or channel has to wait
for everything to be completed in order to start their execution.
2. You use fewer variables and therefore less memory space because you do not have to
save everything as a variable.
3. The use of pipelines simplifies the design of the program and improves its maintainability.