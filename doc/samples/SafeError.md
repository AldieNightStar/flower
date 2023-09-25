# Errors and Safe commands


## Error command
```clojure
; Throw new error
; Error could be any type
(error "Can't handle this stuff")
```


## Catch any error
```clojure
(safe
    ; Let's assume this call will throw an error
    (call-suspicious-api 1 2 3)

    ; Adds a handler in case of any errors
    ; Also if any errors happens then this handle call result is called
    (handle error
        ; Let's check which type this error is
        ; According to their type we doing something
        (when (is-type error ?)
            str      (c.log (str "Error happened: " error))
            os.nomem (c.log "Not enought RAM")
            else     (c.log (str "Unknown error: " error))
        )
    )
)
```



## Catch error into callback
```clojure
; Will return value of the api result or nil if any error happens
; If any error then (callback error) called
(safed-by callback (call-suspicious-api 1 2 3))


; Small sample
; ===================

; Prepare some sample
(def handler (err)
    (c.log (str "Error happened: " err))
)

; Shorter example
(safe-by handler (call-suspicious-api 1 2 3))

; Longer example
(with a (safe-by handler (call-suspicious-api 1 2 3))
    ; 'handler' is called with an error before this block
    ; Only then this block will be called and 'a' will be nil

    ; Will result:
    ;   Result is Nothing
    (c.log (str "Result is " (maybe a "Nothing")))
)
```