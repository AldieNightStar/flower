# Block of Code

## Notes
* Some time you need to add more code than single expression
* For example to work with `if`, `when` or `case` functions
* You could use `do` for that
* Block returns what is the last in the code





## Block
* Usage
```clojure
(do
    ; Here are some code block

    ; Return value
    123
)
```
* Example with `if`
```clojure
(if (= a 1) (do
    (c.log "A is equal to one")
) (do
    (c.log "A is NOT equal to one")
))
```





# Block with own scope
* Sometime you need to call block with own scope
* We using `dos` command. Which means `do scoped`
```clojure
(dos
    ; This variable will wipe when block is finished
    (set a 1)

    ; Return value
    10
)
```




