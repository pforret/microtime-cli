# microtime-cli
microtime for the command line (in Golang)

## Usage

```
Usage of microtime:
  -diff float
        Calculate duration from the given time
  -format string
        Output format (options: cmd)
```

## Examples

```bash
# Get the current time in seconds + microseconds
$ microtime
1723639630.583408

# Get a duration in seconds
$ microtime --diff=1723639630.583408
53.081471

# use it in a bash script
STARTED_AT=$(microtime)
(...)
DURATION=$(microtime --diff="$STARTED_AT")
```

```bat
# use it in a Windows cmd script
set TempScript=%TEMP%\microtime.%RANDOM%.cmd
microtime --format=cmd > %TempScript%
# set MICROTIME=1723804285.835174
@call %TempScript%
(...)
microtime --diff=%MICROTIME% --format=cmd > %TempScript%
# set ELAPSED=244.843309
@call %TempScript%
echo %ELAPSED%
```

> Inspired by [blog.forret.com/2022/12/15/bash-microtime/](https://blog.forret.com/2022/12/15/bash-microtime/)