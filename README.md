![CI](https://github.com/turbaszek/gonotes/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/turbaszek/gonotes)](https://goreportcard.com/report/github.com/turbaszek/gonotes)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

# GoNotes
Kindle note reader in cli version!

![book_list](docs/list.gif)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Usage](#usage)
- [Help](#help)
- [Autocomplete](#autocomplete)
- [Installation](#installation)
- [Development](#development)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Usage
First parse your clippings.txt from Kindle:
```
➜ gonotes parse clippings.txt
```

Then you can list notes from a book, press tab use autocomplete
(check how to [enable](#autocomplete) it):
```
➜ ./gonotes notes 8
```

You can display notes from single book and search using grep (!):
```
➜ gonotes notes 3 | grep tourism
tourism is about helping people construct stories and collect memories.
```

Of course remember to use [cowsay](https://en.wikipedia.org/wiki/Cowsay)!
```
➜ gonotes n 12 | grep "personal growth" | cowsay
 ______________________________________
/ Simply having enough money to spare  \
| converts the vicious cycle of stress |
| and poor decision making into a      |
\ virtuous cycle of personal growth.   /
 --------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

There's also possibility to get a random quote that will not be longer than specified number of words
```
➜ gonotes r -q -l 15 | cowsay -f bunny
 ________________________________________
/ Hate is just as injurious to the hater \
| as it is to the hated. - The           |
| Autobiography of Martin Luther King,   |
\ Jr. (Carson, Clayborne)                /
 ----------------------------------------
  \
   \   \
        \ /\
        ( )
      .( o ).
```

## Help
<!-- AUTO_STAR -->
```
  NAME:
     gonotes - Simple tool to manage Kindle notes

  USAGE:
     gonotes [global options] command [command options] [arguments...]

  VERSION:
     v0.1

  COMMANDS:
     parse, p   Parses provided file and creates books and notes
     book, b    Utilities to manage books
     notes, n   List notes
     random, r  Shows a random note
     help, h    Shows a list of commands or help for one command

  GLOBAL OPTIONS:
     --help, -h     show help (default: false)
     --version, -v  print the version (default: false)
```
<!-- AUTO_END -->

## Autocomplete
To set up autocomplete including book hints run one of the following:
```
# bash
source <(gonotes complete bash)

# zsh
source <(gonotes complete zsh)
```
To persist the autocomplete behaviour add this selected option to
your `.bashrc` or `.zshrc`.

## Installation
Currently, you can install GoNotes in two ways:
- clone the repo and then `go build ./cmd/gonotes` and `go install .` - this will install actual master
- download the binary https://github.com/turbaszek/gonotes/releases/download/v0.1-alpha2/gonotes

## Development
Feel free to open issues and PRs. To build the project follow are usuall go steps. Consider using
[pre-commits](https://pre-commit.com) for static checks and code formatting. On Mac this should do:
```
brew install pre-commit
pre-commit install
```
