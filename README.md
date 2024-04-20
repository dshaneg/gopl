# The Go Programming Language Exercises

I'm working through [_The Go Programming Language_](https://www.gopl.io/) by Donovan and Kernighan, and here are the exercises.

I may not work through every exercise, and I've modified the examples based on changes in the language since the book was written, as well as a little personal flair here and there.

I included a [`Makefile`](Makefile) that allows you to execute (or build) each example. The make rules are named with a chapter prefix, then slash, then the name of the program. The exercises that require a build instead of a run have a suffix of `-build` on the make rule.

You can run all the programs from the root of the project with a `make` command, and for those programs that need to be built, the binary is output to the `./bin` directory, and can be executed with `./bin/{binaryname}`

## Attribution

The official repository is [here](https://github.com/adonovan/gopl.io/).

The original code (and therefore this repository) is distributed with the Creative Commons BY-NC-SA 4.0 (Attribution-NonCommercial-ShareAlike) license. The following snippet is included in all their example code and is repeated here.

> Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.  
> License: <https://creativecommons.org/licenses/by-nc-sa/4.0/>
