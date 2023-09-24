# Logic


## Logic conditions
```clojure
(use sys.console c)

; Show EQUAL if a same == b
(if (= a b)
    (c.log "EQUAL")
    (c.log "NOT EQUAL")
)

; Let's use 'case'
; Each condition call uses 'a' as a first argument internaly
(case a
    (= b) (c.log "EQUAL TO B")
    (= c) (c.log "EQUAL TO C")
    (= d) (c.log "EQUAL TO D")
    else  (c.log "NOT EQUAL" )
)


; Let's use 'when'
; Same as 'case' but this time you use function at start
; '?' will be replaced by your variant
(when (= a ?)
    b    (c.log "EQUAL TO B")
    c    (c.log "EQUAL TO C")
    d    (c.log "EQUAL TO D")
    else (c.log "NOT EQUAL" )
)
```

* Let's check the type

```clojure
; Let's check the type
; Using method (is-type a User)
(case a
    (is-type User) (c.log "User")
    (is-type int)  (c.log "Int" )
)

; Let's use 'when'
(when (is-type a ?)
    User (c.log "User")
    int  (c.log "Int")
)
```