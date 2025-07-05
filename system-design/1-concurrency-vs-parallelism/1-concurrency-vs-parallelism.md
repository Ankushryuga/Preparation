## Concurrency: 
    =>
    1. It refers to multiple tasks being extecuted simultaneously.
    2. it normally uses one or more cores -  can works on 1 core using context switching.
    
## Parallelism:
    =>
    1. It refers to spliting a task into smaller sub-tasks and running them simultaneously.
    2. Requires multiple cores or processes.

## Context Switching:
    =>
    It is the process of saving and restoring the state of a thread or process so that execution can be paused and resumed.


## Why context switching?
    =>
    The process of saving and restoring the state of a thread or process so that execution can be paused and resumed.
    1. CPU Core can only run one thread at a time.
    2. When there are many tasks, the system rapidly switches b/w them to give the illusion of multitasking.
    3. Each switch involves saving the current tasks state and loading another's -- this is context switch.

## What gets switched?
    =>
    1. Registers (e.g, instruction pointer)
    2. Stack
    3. Program counter
    4. Memory mapping
    5. CPU Cache (often flushed)

## Downsides:
    =>
    1. Overhead: Takes time and CPU cycles.
    2. Cache invalidation: CPU cache might not help after a switch.
    3. Can be expensive in OS-level threads
    
