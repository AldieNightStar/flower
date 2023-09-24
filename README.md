# Flower

## Overview

* Lisp based programming language
* `Code is data` in mind
* ⚠️It's not the same as __LISP__ and it's __NOT__ a dialect. It's completely different language


## Idea
```clojure
(use sys.console c)

(def main ()
    (c.log "Hello there")
)
```
* Dependency from Internet (`net`)
```clojure
(use sys.console c)
(use net.jason.ver04381941 J) ; Let's imagine

(J.encode (map :a 1 :b 2 :c 3))

(with it (J.decode "{"a": 1, "b": 2, "c": 3}")
    (c.log it.a)
    (c.log it.b)
    (c.log it.c)
)
```
* Some logic:
```clojure
(use sys.console c)

; Show EQUAL if a same == b
(if (= a b)
    (c.log "EQUAL")
    (c.log "NOT EQUAL")
)

; Let's use 'case'
(case a
    (= b) (c.log "EQUAL TO B")
    (= c) (c.log "EQUAL TO C")
    (= d) (c.log "EQUAL TO D")
    else  (c.log "NOT EQUAL" )
)
```
* Types
```clojure
(use sys.console c)

(type User
    (name str)
    (age int)
    (def is-young () (>= age 18))
)

(interface YoungChecker
    (is-young () bool)
)
```

## Documentation

* [Documentation pages](doc/README.md)