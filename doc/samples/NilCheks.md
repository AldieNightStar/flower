# Nil Checks

## Maybe
```clojure
; Return first non-nil value
; Last value SHOULD not be nil but default value or variable
; If last value is nil anyway then error is thrown
(maybe a b 0)
```




## Hand checks
```clojure
; Returns true if value is nil
(is-nil user)

; Returns true if value is NOT nil
(is-not-nil user)
```

