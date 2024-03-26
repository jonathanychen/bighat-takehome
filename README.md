# BigHat Takehome Assignment

## General Overview

Simple program written in Go v1.21.6, which utilizes concurrency to allow for running multiple timers at once. 

To run the code, you can either use the provided binary executable ```bighat``` or install Go and then run the command ```go run .``` in the root directory. Alternatively, you can also re-build the executable with ```go build``` and use the binary the same way. 

To modify the input for testing, no changes to the code are necessary. All you need to do is change the contents of ```input.json```. This works for either method of running the code.

One major assumption being made: the input is well-formed. I wanted to spend more time on the solution itself rather than input validation, so this is the one assumption I made for my code.

## Thought Process

Concurrency is my preferred way of tackling this problem, because then using a blocking call like ```time.Sleep``` is viable. This led me to use Go because not only am I currently learning it, but it also provides a really easy interface for setting up concurrency through its <a href="https://gobyexample.com/goroutines">Goroutines</a>. 

These Goroutines, when used with <a href="https://gobyexample.com/waitgroups">wait groups</a>, allow for a relatively clean recursive implementation.


### Sample Test:

Input (given in assignment): 

```{"A": {"start": true, "edges": {"B": 5, "C": 7}}, "B": {"edges": {}}, "C": {"edges": {}}}```

Output:

```
Start: 2024-03-26 14:46:09.855803 -0400 EDT m=+0.000085710
0.000513958 A
5.001675041 B
7.001829166 C
End: 7.002065166s
```

*Note that timestamps are given before each node label is printed, just to illustrate the timing of each print. Timing is also not perfectly accurate, but this is somewhat expected and (hopefully) acceptable.

More sample test cases are given in the ```sample_tests``` folder, along with diagrams illustrating each sample and a text file representing the expected output from that test. Timestamps are approximated in the expected output files.