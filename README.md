# Flower

* ⚠️ __This project is Work in Progress and not done to be used by someone__

## Overview

* Programming language that looks like lisp, but not LISP :)
* `Code is data` in mind
* ⚠️It's not the same as __LISP__ and it's __NOT__ a dialect. It's completely different language


## Idea
* [More sample here](doc/samples/README.md)
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

## Documentation

* [Documentation pages](doc/README.md)
