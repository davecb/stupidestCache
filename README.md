# stupidestCache

## Description
This is a set of caches, created to use 
as the lowest bound when comparing the
performance of cache offerings
* stupidestCache is probably as simple as you can get, 
and  faster than any credible cache. It has lots of limitations.
* mvpCache has a time field, so it can be used in a more normal 
configuration, such as in front of an L2 cache

Both are loadable via a .csv file, and exercisable via a web service.    
The latter is used by [Play it Again, Sam](https://github.com/davecb/Play-it-Again-Sam) to 
measure their performance under increasing load

### stupid Limitations
This cache needs to be reloaded periodically by an outside actor,
and be halted occasionally to clear any unused cache slots. That's
fine for a library that's updated regularly in a program which 
does a blue-green restart nightly. That's barely credible: I have 
one program that behaves that way, out of everything I've ever written.

### mvp Limitations
This cache has a last-refreshed time, so it can sit in front of an L2
cache and poll the back end when it starts to get stale. 
It will free a cache line if it is empty and more than a certain age,
so you don't have to restart it to get it to garbage-collect.
This meets the definition of minimum viable product: I could 
use it in a few services I know of.

### Parmeterization
Both use a standard non-locked Go map, but that can be replaced with a 
faster one like Swiss-Map or a custom implementation 
if you want to see if the map is the dominating cost

# Man Pages

## NAME
cmd/stupid - exercise the cache

## SYNOPSIS
go run cmd/stupid/main.go file.csv

## DESCRIPTION
cmd/stupid is used to exercise the stupid cache. It reads a csv file and executes the command in it

## FILES
The csv files contain either "g" or "p", followed by a key and value seperated by spaces.
"g" stands for get, "p" put.
The second field is interpreted as the key, the third and subsequent field as the value.
An example is cmd/stupid/testdata/minimal.csv:
```shell
p fred flintstone and family
g fred flintstone and family
x no such thing as x
p short
```

## EXAMPLE
```sh
$ go run cmd/stupid/main.go  cmd/stupid/testdata/minimal.csv
2023/04/02 09:19:32 main.go:63: stupid read ["p" "fred" "flintstone" "and" "family"]
2023/04/02 09:19:32 main.go:73: operation = "p", key = "fred", value = "flintstone and family", err = <nil>
2023/04/02 09:19:32 main.go:96: put suceeded
2023/04/02 09:19:32 main.go:63: stupid read ["g" "fred" "flintstone" "and" "family"]
2023/04/02 09:19:32 main.go:73: operation = "g", key = "fred", value = "flintstone and family", err = <nil>
2023/04/02 09:19:32 main.go:108: get and comparison suceeded
2023/04/02 09:19:32 main.go:63: stupid read ["x" "no" "such" "thing" "as" "x"]
2023/04/02 09:19:32 main.go:73: operation = "x", key = "no", value = "such thing as x", err = <nil>
2023/04/02 09:19:32 main.go:85: ill-formed csv line. Need either "p" or "g". record = ["x" "no" "such" "thing" "as" "x"]
exit status 1
```

