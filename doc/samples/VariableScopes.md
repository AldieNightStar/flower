# Variable Scopes

## Notes
* When you create function it will have it's own scope
* Calls to some functions reuse parent scope unless it sets it by hands for example `(coro)` has its own scope
* You can `set` and get some value in the scope.
* If value is not present in the scope it goes into parent scope
* If value is not exists then it will be `nil`
* Child scope __CAN'T__ assign variables above their scoe. Need to use `cell` values




## Create new scope
* To create new scope use `let` it will create new scope for your `get/set` operations
```clojure
(let ()
    (set a 32)
    (set b 44)
    (set c 99)
)
```


## DSL (Domain Specific Language)
* Flower supports `DSL` calls
```clojure
; def-dsl doesn't takes any arguments as it uses everything that you 'add' inside
(def-dsl new-user

    ; Prepare some data
    (set user (map :name "" :age 0))

    ; Set up DSL functions for the call
    (export set-age  (fn (age) (set user.age  age )))
    (export set-name (fn (age) (set user.name name)))

    ; Call the function contents
    (yield)

    ; Return user map after calls
    user
)
```
* Let's use this function
```clojure
(set user (new-user
    (set-name "Haxi")
    (set-age 30)
))
(user.name)
```


## Set variable behind the current scope
* To do it, you need to use `cell` command that creates variable that is a reference to itself
* `cell` values has `val` value that can be used to keep the variable inside
* Cells is cheaper than `maps` because they has only one single `val` field to modify or read
```clojure
(set counter (cell))

(let ()
    (set a 111)
    (set b 320)

    ; 'val' field will be changed inside cells
    (set counter.val (+ a b))

    ; You can also get that value from cell
    counter.val
)
```