![CI](https://github.com/turbaszek/gonotes/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/turbaszek/gonotes)](https://goreportcard.com/report/github.com/turbaszek/gonotes)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


# GoNotes
Yet another minimalist kindle note reader. This time as cli tool.

```
NAME:
   GoNotes - Simple cli tool to manage Kindle notes

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
   parse    Parses provided file
   note     Notes related operations
   book     Books related operations
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

First parse your clippings.txt from Kindle:
```
➜ gonotes parse clippings.txt
```

Then you can list your books
```
➜ gonotes book ls
2 | The Autobiography of Martin Luther King, Jr. (Carson, Clayborne)
3 | Thinking, Fast and Slow (Kahneman, Daniel)
4 | Interventions: A Life in War and Peace (Annan, Kofi)
5 | Psychologia jogi. Wprowadzenie do "Jogasutr" Patańdźalego (Maciej Wielobób)
6 | Diary of a Professional Commodity Trader: Lessons from 21 Weeks of Real Trading (Brandt, Peter L.)
```

and notes from single book and search using grep (!):
```
➜ ./gonotes note cat 3 | grep tourism
536 | tourism is about helping people construct stories and collect memories.
```

And of course remember to use [cowsay](https://en.wikipedia.org/wiki/Cowsay)!
```
➜ ./gonotes note cat 3 | grep tourism | cowsay
 _________________________________________
/ 536 | tourism is about helping people   \
\ construct stories and collect memories. /
 -----------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```
