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


### **1.2.1 Race Conditions**

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.
Most of the time, this shows up in what’s called a *data race*, where one concurrent operation attempts to read a variable while at some undetermined time another concurrent operation is attempting to write to the same variable.

### **1.2.2 Atomicity**


### **1.2.3 Memory Access Synchronization**


### **1.2.4 Deadlocks, Livelocks, and Starvation**


### **1.7 Determining Concurrency Safety**


### **1.8 Simplicity in the Face of Complexity**