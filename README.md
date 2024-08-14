# microtime-cli
microtime for the command line (in Golang)

## Usage

```
Usage of microtime:
    -diff float  :  Calculate duration from the given time
```

## Examples

```bash
# Get the current time in seconds + microseconds
$ microtime
1723639630.583408

# Get a duration in seconds
$ microtime --diff=1723639630.583408
53.081471

# use it in a script
STARTED_AT=$(microtime)
(...)
DURATION=$(microtime --diff="$STARTED_AT")
```
> Inspired by [blog.forret.com/2022/12/15/bash-microtime/](https://blog.forret.com/2022/12/15/bash-microtime/)