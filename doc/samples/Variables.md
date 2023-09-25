# Variables

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