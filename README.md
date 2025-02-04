This repository contains basic simple programs, pattern displays, go rountines, Rest Apis and LinkedList samples.
Contents would be updated as it goes.
Anyone who needs to prepare for interviews can first go through study material on below repos then can practice the above basic problem statements.

More interview prep content.
* Credits goes to owner of below repos. 

https://github.com/pedrobertao/golang-interview-prep/tree/main

https://github.com/Devinterview-io/golang-interview-questions?form=MG0AV3

1st Feb 2025 Deutsche Bank interview questions

    A) 
	B) 
	C) 
	D) 
    Ans: 

1) Which function to used to create map
    A) append
	B) cap
	C) make
	D) None
    Ans: C

2) Which functions can be used to copy array ?
    A) append
	B) cap
	C) copy
	D) None
    Ans: I tick A and C
        dest = append(dest,src)
        copy(dest,src)

3) What is output of below code.
	A) Address of x
	B) Error
	C) 0
	D) 5
    
        func myfunc(a *int) {
            *a = 0
        }

        x := 5
        myfunc(&x)
        fmt.Println("x is :", x)
    Ans: C

4) How to increase size of slice
    A) append
	B) resize
	C) cap
	D) len
    Ans: A

5) What is I/O package 
    