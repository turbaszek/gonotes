![CI](https://github.com/turbaszek/gonotes/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/turbaszek/gonotes)](https://goreportcard.com/report/github.com/turbaszek/gonotes)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


# GoNotes
Yet another minimalist kindle note reader. This time as cli tool.

![book_list](img/list.png)


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

And of course remember to use [cowsay](https://en.wikipedia.org/wiki/Cowsay)!
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
