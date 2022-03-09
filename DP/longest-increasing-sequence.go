package DP

/*
Longest Increasing Subsequence is another example of a classic dynamic programming problem.
In this problem, we need to determine the length of the longest (first characteristic) subsequence that is strictly increasing.
For example, if we had the input nums = [1, 2, 6, 3, 5] \text{nums = [1, 2, 6, 3, 5]} nums = [1, 2, 6, 3, 5],
the answer would be 4, from the subsequence [1, 2, 3, 5] \text{[1, 2, 3, 5]} [1, 2, 3, 5].
Again, the important decision comes when we arrive at the 6 - do we take it or not take it? If we decide to take it,
then we get to increase our current length by 1, but it affects the future - we can no longer take the 3 or 5.
Of course, with such a small example,
it's easy to see why we shouldn't take it - but how are we supposed to design
an algorithm that can always make the correct decision with huge inputs?
Imagine if nums contained 10,00010,00010,000 numbers instead.
*/
