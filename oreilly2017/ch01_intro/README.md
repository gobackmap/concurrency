# **1. An Introduction to Concurrency**

Concurrency is an interesting word because it means different things to different people in our field. In addition to “concurrency,” you may have heard the words, “asynchronous,” “parallel,” or “threaded” bandied about. Some people take these words to mean the same thing, and other people very specifically delineate between each of
those words.

- A practical definition:

>
> When most people use the word “concurrent,” they’re usually referring to a process that occurs simultaneously with one or more processes. It is also usually implied that all of these processes are making progress at about the same time. Under this definition, an easy way to think about this are people. You are currently reading this sentence while others in the world are simultaneously living their lives. They are existing concurrently to you.
>

- Why concurrency became an important topic in computer science?
- Why concurrency is difficult and warrants careful study?

## **1.1 Moore’s Law, Web Scale, and the Mess We’re In**

- Gordon **Moore's Law** (1965-1975): the number of components on an integrated circuit would double every two years. This prediction more or less held true until just recently—around 2012 --> born of **multicore processors**: a clever way to solve the bounding problems of Moore’s law | bounded by the limits of the Amdahl's Law. 

- Gene **Amdahl’s law** describes a way in which to model the potential performance gains from implementing the solution to a problem in a parallel manner. Simply put, it states that the gains are bounded by how much of the program must be written in a sequential manner.


For example, imagine you were writing a program that was largely *GUI based*: a user is presented with an interface, clicks on some buttons, and stuff happens. This type of program is bounded by one very large sequential portion of the pipeline: human interaction. No matter how many cores you make available to this program, it will always be bounded by how quickly the user can interact with the interface.


Now consider a different example, *calculating digits of pi*. Thanks to a class of algorithms called spigot algorithms, this problem is called *embarrassingly parallel*, which —despite sounding made up—is a technical term which means that it can easily be divided into parallel tasks. In this case, significant gains can be made by making more cores available to your program, and your new problem becomes how to combine and store the results.

Amdahl’s law helps us understand the difference between these two problems, and can help us decide whether parallelization is the right way to address performance concerns in our system.

For problems that are *embarrassingly parallel*, it is recommended that you write your application so that it can *scale horizontally*. This means that you can take instances of your program, run it on more CPUs, or machines, and this will cause the runtime of the system to improve. Embarrassingly parallel problems fit this model so well because it’s very easy to structure your program in such a way that you can send chunks of a problem to different instances of your application.

- **Cloud Computing (1970-2000)**: a new kind of scale and approach to application deployments and horizontal 
scaling. --> A new kind of thinking: 
  - cheap access to vast amounts of computing power to solve large problems.
  - Solutions could now trivially span many machines and even global regions.
  - made solved problems that were previously only solvable by tech giants.
  - Challenges:
    - Provisioning these resources
    - Communicating between machine instances
    - Aggregating and storing the results
    - How to model code concurrently: The fact that pieces of your solution could be running on disparate machines exacerbated some of the issues commonly faced when modeling a problem concurrently. Successfully solving these issues soon led to a new type of brand for software, *web scale*.

- **Web Scale**: If software was web scale, among other things, you could expect that it would be *embarrassingly parallel*; that is, web scale software is usually expected to be able to handle hundreds of thousands (or more) of simultaneous workloads by adding more instances of the application. This enabled all kinds of properties like rolling upgrades, elastic horizontally scalable architecture, and geographic distribution. It also introduced new levels of complexity both in comprehension and fault tolerance.

- Modern developers my be overwhelmed in this world of
  - Multiple cores
  - Cloud computing
  - Web scale
  - Problems that may or may not be parallelizabl 

- **Herb Sutter** (2015):
>
> - "The free lunch is over: A fundamental turn toward concurrency in software."
> - “We desperately need a higher-level programming model for concurrency than languages offer today.”
>


## **1.2 Why Is Concurrency Hard?**
Following are some of the most common issues that make working with concurrent code both frustrating and interesting.

### **1.2.1 Race Conditions**

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.
Most of the time, this shows up in what’s called a *data race*, where one concurrent operation attempts to read a variable while at some undetermined time another concurrent operation is attempting to write to the same variable.

### **1.2.2 Atomicity**

When something is considered atomic, or to have the property of atomicity, this means that within the *context* that it is operating, it is *indivisible*, or *uninterruptible*.

- Something may be atomic in one **context**, but not another. In other words, the atomicity of an operation can change depending on the currently defined **scope**. This fact can work both for and against you!
- The terms **indivisible** and **uninterruptible** mean that within the context you’ve defined, something that is atomic will happen in its entirety without anything happening in that context simultaneously. Let's look at the example of i++. It may look atomic, but a brief analysis reveals several operations: 
  - Retrieve the value of i.
  - Increment the value of i.
  - Store the value of i.
  
While each of these operations alone is atomic, the combination of the three may not be, depending on your context: If your context is a program with no concurrent processes, then this code is atomic within that context. If your context is a goroutine that doesn’t expose i to other goroutines, then this code is atomic.
- Atomicity is important because if something is atomic, implicitly it is safe within concurrent contexts. This allows us to
  - compose logically correct programs,
  - optimize concurrent programs.
- Most statements are not atomic, let alone functions, methods, and programs. If atomicity is the key to composing logically correct programs, and most statements aren’t atomic, how do we reconcile these two statements? In short we can force atomicity by employing various techniques. The art then becomes determining which areas of your code need to be atomic, and at what level of granularity.

### **1.2.3 Memory Access Synchronization**


### **1.2.4 Deadlocks, Livelocks, and Starvation**

Even if you successfully handle program correctness classes of issues, there is another class of issues to contend with: **deadlocks**, **livelocks**, and **starvation**. These issues all concern ensuring your program has something useful to do at all times. If not handled properly, your program could enter a state in which it will stop functioning altogether.
#### **1.2.4.1 Deadlocls**
It turns out there are a few conditions that must be present for deadlocks to arise, and in 1971, Edgar Coffman enumerated these conditions in a paper. The conditions are now known as the **Coffman Conditions** and are the basis for techniques that help detect, prevent, and correct deadlocks. The Coffman Conditions are as follows:

- **Mutual Exclusion**: A concurrent process holds exclusive rights to a resource at any one time.
- **Wait For Condition**: A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
- **No Preemption**: A resource held by a concurrent process can only be released by that process, so it fulfills this condition.
- **Circular Wait**: A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1), so it fulfills this final condition too.


These laws allow us to prevent deadlocks too. If we ensure that at least one of these conditions is not true, we can prevent deadlocks from occurring. Unfortunately, in practice these conditions can be hard to reason about, and therefore difficult to prevent.
#### **1.2.4.2 Livelocls**

Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.

Livelocks are more difficult to spot than deadlocks simply because it can appear as if the program is doing work. If a livelocked program were running on your machine and you took a look at the CPU utilization to determine if it was doing anything, you might think it was. Depending on the livelock, it might even be emitting other signals that would make you think it was doing work. And yet all the while, your program would be playing an eternal game of hallway-shuffle.


#### **1.2.4.3 Starvations**

### **1.7 Determining Concurrency Safety**


### **1.8 Simplicity in the Face of Complexity**