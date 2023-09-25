# Coroutines

## Notes
* Coroutines is a functions that can run with blocking itself without blocking the main thread
* When coroutine gets created then `coro` object is returned
* Coroutines __DO NOT RUNS__ in a scheduler. You need to run them by hands
* Coroutines can exchange data between caller by `coro.yield` method
    * `coro.yield` could accept data to send to the caller
    * `coro.yield` could also return data if caller sent something to it
    * `coro.yield` blocks coroutine until someones call `coro.run` to it
    * `coro.run` could accept any parameters to send into coroutine
    * `coro.run` could return some parameters from coroutine if `coro.yield` has something to return
* Coroutine has 3 states
    * 1 - Running (When code is just running)
    * 2 - Waiting (Blocked by `coro.yield`)
    * 3 - Finished
        * To check that coroutine is finished: `(coro.done? c)`




## Create
```clojure
(set c
    ; Create new coroutine
    (coro
        (c.log "Hello")

        ; Make coroutine blocked until someone run it back
        ; Could yield some data to the caller
        ; Could also take some data from the caller
        (coro.yield)
        (set dat (coro.yield))     ; Take data from the caller
        (coro.yield 100)           ; Send some data
        (set dat (coro.yield 123)) ; Send data to the caller and take some

        (c.log "Done")
    )
)
```