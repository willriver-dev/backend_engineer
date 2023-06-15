## 1. How different is Functional Programming from Object Oriented Programming? Explain.

Functional Programming (FP) and Object-Oriented Programming (OOP) are two different paradigms in software development. Here are the key differences between the two:
    * Data and State:
        OOP: In OOP, data and behavior are encapsulated within objects. Objects store state and have methods to manipulate that state.
        FP: In FP, data and behavior are separated. Data is immutable, and functions operate on this data to produce new values. FP promotes the use of pure functions that have no side effects and always produce the same result for the same inputs.

    - Mutability:
        OOP: OOP allows mutable state, meaning objects can change their state during execution.
        FP: FP encourages immutability. Data is not changed after it is created, and there is no change in state. Instead, functions create new values based on the original data.
    - Control Flow:
        OOP: OOP uses control structures like loops and conditionals to control the program flow. Mechanisms such as inheritance and polymorphism are also used to manage control flow.
        FP: FP uses function composition, recursion, and higher-order functions to control program flow. FP focuses on representing computations by combining functions, rather than using traditional control structures.
    - Side Effects:
        OOP: OOP often involves side effects, such as modifying shared state or performing I/O operations. This can lead to complex interactions and make the code difficult to test and reason about.
        FP: FP encourages avoiding side effects and emphasizes the use of pure functions. Pure functions have no side effects and always produce predictable results, without affecting external state.

