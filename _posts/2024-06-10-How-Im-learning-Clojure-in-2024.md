---
layout: post
title: How I'm learning Clojure in 2024
published: true
---

I've recently been learning a bit of Clojure and it's been a lot of fun!
I thought I would note down what has been useful for me, so that it might help others as well.

## Jumping right in
[https://tryclojure.org/](https://tryclojure.org/) is a great intro to the language. It provides a
REPL and a tutorial that takes you through the basic features of Clojure.

Importantly, it forces you to get used to seeing lots of `(` and `)`!

## Doing exercises
Exercism provides small coding challenges for a bunch of languages, including [Clojure](https://exercism.org/tracks/clojure).
Unlike other platforms (*cough* leetcode *cough*), Exercism is focused on learning and I found it a
great way to practice writing Clojure.

It provides a code editor and evaluates each challenge when you submit. There's also a way to submit
answers locally from your computer, but I found it quicker just to use the website.

## Editor setup
I ended up setting up Neovim for developing locally. This guide was a great inspiration:
[https://endot.org/2023/05/27/vim-clojure-dev-2023/](https://endot.org/2023/05/27/vim-clojure-dev-2023/), although I did end up going with something a little bit simpler.

My .vimrc can be seen
[here](https://github.com/AussieGuy0/dotfiles/blob/e25b8da7c01ea1358723a19ca1319cab4888beff/.vimrc)
but probably the most important plugin is [Conjure](https://github.com/Olical/conjure), which
provides REPL integration in Neovim.

The REPL is one of the big differences compared to programming in other
languages. Basically, you start a REPL in the project directory and then you can evaluate code in
your editor in that REPL.

This basically gives you really short iteration cycles, you can 'play' with your code, run
tests, reload code in a running app and all without leaving your editor!

To understand REPL driven development, I really liked this [video with teej_dv and
lispyclouds](https://www.youtube.com/watch?v=uBTRLBU-83A). One key thing I learnt was using the
`comment` function to be able to evaluate code without affecting the rest of my program.

```clojure
; my super cool function
; given a number, adds 2 to it!
(defn add-2
  [n]
  (+ n 2)
)

; This tells Clojure to ignore what comes next
; but it still has to be syntactically correct!
(comment
  (add-2 3) ; <-- I can evaluate this to check my add-2 function :)
)
```

By opening a REPL and using the Conjure plugin mentioned before I can:

- `,eb`: Evaluate the buffer I am in. Kinda like loading up the file I have opened into the REPL.
- `,ee`: Evaluate the expression my cursor is under.
- `,tc`: Run the test my cursor is over.
- `,tn`: Run all tests in current file.

I use the following alias in my `.bash_aliases` to easily spin up a REPL:

```sh
# From https://github.com/Olical/conjure/wiki/Quick-start:-Clojure
# Don't ask me questions about how this works, but it does!

alias cljrepl='clj -Sdeps '\''{:deps {nrepl/nrepl {:mvn/version "1.0.0"} cider/cider-nrepl {:mvn/version "0.42.1"}}}\'\'' \
    --main nrepl.cmdline \
    --middleware '\''["cider.nrepl/cider-middleware"]'\'' \
    --interactive'
```

## Docs
For docs, I really like [https://clojuredocs.org/](https://clojuredocs.org/), which has the docs for
the core library. I like the fact that users can submit code examples, which provides better
information for each function.

## Projects
I've currently have 2 projects in Clojure to further my learning.

1. A bad terminal-based clone(ish) of
[Balatro](https://store.steampowered.com/app/2379780/Balatro/). Balatro is a very addictive deck
builder rougelike game. Doing this has been fun and it feels very natural to 'build up' over time.
The source code can be seen [here](https://github.com/AussieGuy0/balatro-clj).
2. A application that converts a subreddit into an RSS feed. The idea that this can be a webapp that
   produces daily RSS feeds for a collection of subreddits. [Source
   code](https://github.com/AussieGuy0/reddit-to-rss)

## The End
Thanks for reading!

