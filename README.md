![CI](https://github.com/turbaszek/gonotes/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/turbaszek/gonotes)](https://goreportcard.com/report/github.com/turbaszek/gonotes)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


# GoNotes
Yet another minimalist kindle note reader. This time as cli tool.

![book_list](docs/list.gif)


## Usage
First parse your clippings.txt from Kindle:
```
➜ gonotes parse clippings.txt
```

Then you can list your books and select one to show related notes.
```
➜ gonotes book ls
Use the arrow keys to navigate: ↓ ↑ → ←
? Your books::
  1 | How to use Knotes
  2 | The Autobiography of Martin Luther King, Jr. (Carson, Clayborne)
  3 | Thinking, Fast and Slow (Kahneman, Daniel)
  4 | Interventions: A Life in War and Peace (Annan, Kofi)
  5 | Psychologia jogi. Wprowadzenie do "Jogasutr" Patańdźalego (Maciej Wielobób)
  6 | Diary of a Professional Commodity Trader: Lessons from 21 Weeks of Real Trading (Brandt, Peter L.)
  7 | The Checklist Manifesto: How to Get Things Right (Gawande, Atul)
↓ 8 | Superforecasting (Philip E. Tetlock)
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

You can also select a random quote that will not be longer than specified number of words
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

## Autocomplete
To set up nice autocomplete (including book hints) run one of the following:
```
# bash
source <(gonotes complete bash)

# zash
source <(gonotes complete zsh)
```
To persist the autocomplete behaviour add this to your `.bashrc` or `.zshrc`.

## Installation
Currently, you can install GoNotes in two ways:
- clone the repo and then `go build ./cmd/gonotes` and `go install .` - this will install actual master
- download the binary https://github.com/turbaszek/gonotes/releases/download/v0.1-alpha.1/gonotes

## Development
Feel free to open issues and PRs. To build the project follow are usuall go steps. Consider using
[pre-commits](https://pre-commit.com) for static checks and code formatting. On Mac this should do:
```
brew install pre-commit
pre-commit install
```
