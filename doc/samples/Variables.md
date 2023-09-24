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



## `let` command
```clojure
; 'let' allows you initialize couple of mutable variables
(let (a 10 b 20 c 30)
    (c.log a)
    (c.log b)
    (c.log c)
)
```