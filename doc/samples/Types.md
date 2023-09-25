# Types

## Basic types
* `int` - integer number
* `float` - float number
* `str` - string
* `atom` - string (but more safe)
* `list` - list of elements
* `map` - map by element




## New Type
```clojure
(use sys.console c)

; Create User type with 'name' and 'age' and method 'is-young'
(type User
    (name str)
    (age int)
    (def is-young () (<= age 18))
)

; Using 'with' to work with the variable
(with user (User.new "Haxi" 32)
    (c.log (user.is-young))
)
```



## Map Type
```clojure
(map :key "value")
(map "key" "value") ; the same
(map key "value") ; using key from the variable

; Check that variable is of 'map' type
; Returns true if variable is map
(map? val)
```


## List Type
```clojure
(list 10 20 30 40)

; Some operation
; Each operation creates new list
(list.add (list 10 20) 30) ; (list 10 20 30)

; Check that value is list
(list? val)
```