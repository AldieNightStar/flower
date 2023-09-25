# Variables

## Scopes
* When you create function it will have it's own scope
* Calls to some functions reuse parent scope unless it sets it by hands for example `(coro)` has its own scope
* You can `set` and get some value in the scope.
* If value is not present in the scope it goes into parent scope
* If value is not exists then it will be `nil`
[More about Scopes](VariableScopes.md)


## `with` command
```clojure
; First argument is a constant of returned value from (download-data) function call
; After block is executed then it calls (data.close) if this method is present
(with data (download-data)
    (c.log data)
)

; (file.close) will be called at the end
(with file (open-file "test.txt" :read)
    (c.log (file.readline))
    (c.log (file.readline))
    (c.log (file.readline))
    (c.log (file.readline))
)
```



## `set` command
```clojure
; Set variable in the current scope
; Scope lives until code is ended
(set a 10)
```