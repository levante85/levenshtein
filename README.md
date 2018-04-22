# levenshtein

The levenshtein package implements the efficient version of the levenshtein distance algorithm which calculates how many edits it takes to change one string into another one ( it takes into account, insertion, deletion and substitution ).

For example consider the following two strings "GUMBO" and "GAMBOL":
* GUMBO -> GAMBO  ( 1 substitution )
* GAMBO -> GAMBOL ( 1 insertion )

It's runtime is O(n*m) where n is the length of the first string and m is the
length of the second string, space requirements are equal O(n*m). The recursive
inefficient but more straight forward version is 15-20x slower on small strings,
I have included a benchmark to show the difference ( the recursive version is a
golang port of the wikipedia article ).
A more detailed explanation of the algorithm with implementations in java, c++ and vb, can be seen at the following [link](https://people.cs.pitt.edu/~kirk/cs1501/Pruhs/Spring2006/assignments/editdistance/Levenshtein%20Distance.htm).