# The Go Programming Language Exercises

I'm working through [_The Go Programming Language_](https://www.gopl.io/) by Donovan and Kernighan, and here are the exercises.

I may not work through every exercise, and I've modified the examples based on changes in the language since the book was written, as well as a little personal flair here and there.

I included `Makefile`s in each chapter's directory that allow you to build and in some cases execute each example. I also included a root `Makefile` that includes the `all` and `clean` rules, along with rules to build all the examples in each chapter independently. You can run the chapter-based `make` rules either by changing to the chapter directory and calling make in that directory, or by staying in the project directory and using `make -C {chapterdir}`.

All binaries are built to the `bin` directory in the project root, and can be executed with `./bin/{binaryname}`

## Attribution

The official repository is [here](https://github.com/adonovan/gopl.io/).

The original code (and therefore this repository) is distributed with the Creative Commons BY-NC-SA 4.0 (Attribution-NonCommercial-ShareAlike) license. The following snippet is included in all their example code and is repeated here.

> Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.  
> License: <https://creativecommons.org/licenses/by-nc-sa/4.0/>
