# Loops

## Iterate
* Iterate each element in the list
```clojure
(iterate (list 10 20 30 40) element
    (c.log element)
)
```


## Endless loop
* Will run __forever__ until you call `(break)`
```clojure
(loop
    ; Stop looping
    (break)
)
```



## While loop
```clojure
(while (< a 10)
    (c.log a)
    (-= a 1)

    ; Could also be stopped
    (break)
)
```



## Repeat `n` times
```clojure
(repeat 10 i

    ; Will print numbers from 0 to 9
    (c.log i)

    ; Could also be stopped
    (break)
)
```