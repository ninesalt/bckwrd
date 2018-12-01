# bckwrd 

Bckwrd is a proof-of-concept implementation of an inference engine that does backwards chaining to evaluate queries. The idea here is to use the A* search algorithm instead of DFS (which Prolog uses). Bckwrd is written in Go to make use of it's efficiency (mainly pointers) which should speed up evaluation.

