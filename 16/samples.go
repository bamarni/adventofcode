package main

const rawSamples = `Before: [0, 1, 2, 1]
14 1 3 3
After:  [0, 1, 2, 1]

Before: [3, 2, 2, 3]
13 2 1 3
After:  [3, 2, 2, 1]

Before: [2, 2, 2, 2]
13 2 1 0
After:  [1, 2, 2, 2]

Before: [0, 1, 0, 3]
12 1 0 2
After:  [0, 1, 1, 3]

Before: [0, 1, 0, 3]
11 3 1 3
After:  [0, 1, 0, 0]

Before: [0, 1, 2, 1]
15 1 2 2
After:  [0, 1, 0, 1]

Before: [2, 1, 2, 0]
15 1 2 0
After:  [0, 1, 2, 0]

Before: [3, 1, 2, 0]
11 2 2 1
After:  [3, 1, 2, 0]

Before: [3, 2, 2, 1]
3 3 2 3
After:  [3, 2, 2, 1]

Before: [1, 0, 3, 3]
11 3 3 2
After:  [1, 0, 1, 3]

Before: [1, 0, 2, 1]
5 0 2 0
After:  [0, 0, 2, 1]

Before: [1, 1, 3, 0]
9 1 2 1
After:  [1, 0, 3, 0]

Before: [0, 1, 1, 1]
12 1 0 1
After:  [0, 1, 1, 1]

Before: [1, 3, 2, 2]
5 0 2 1
After:  [1, 0, 2, 2]

Before: [0, 1, 0, 3]
12 1 0 0
After:  [1, 1, 0, 3]

Before: [2, 3, 2, 2]
13 2 0 2
After:  [2, 3, 1, 2]

Before: [1, 1, 3, 1]
1 1 0 0
After:  [1, 1, 3, 1]

Before: [3, 2, 3, 3]
11 3 3 0
After:  [1, 2, 3, 3]

Before: [0, 0, 2, 1]
6 0 0 3
After:  [0, 0, 2, 0]

Before: [1, 3, 0, 3]
4 0 2 3
After:  [1, 3, 0, 0]

Before: [1, 1, 2, 0]
5 0 2 3
After:  [1, 1, 2, 0]

Before: [1, 1, 1, 2]
1 1 0 1
After:  [1, 1, 1, 2]

Before: [3, 0, 1, 1]
10 3 3 3
After:  [3, 0, 1, 0]

Before: [3, 0, 2, 3]
11 3 3 3
After:  [3, 0, 2, 1]

Before: [1, 1, 0, 3]
2 1 3 1
After:  [1, 0, 0, 3]

Before: [1, 1, 2, 0]
1 1 0 2
After:  [1, 1, 1, 0]

Before: [1, 1, 2, 2]
5 0 2 2
After:  [1, 1, 0, 2]

Before: [0, 1, 2, 0]
12 1 0 1
After:  [0, 1, 2, 0]

Before: [0, 1, 2, 0]
12 1 0 2
After:  [0, 1, 1, 0]

Before: [0, 1, 0, 0]
12 1 0 2
After:  [0, 1, 1, 0]

Before: [2, 1, 2, 0]
15 1 2 2
After:  [2, 1, 0, 0]

Before: [1, 1, 2, 1]
15 1 2 2
After:  [1, 1, 0, 1]

Before: [1, 1, 2, 0]
1 1 0 0
After:  [1, 1, 2, 0]

Before: [2, 1, 0, 0]
7 2 0 2
After:  [2, 1, 1, 0]

Before: [0, 0, 3, 0]
0 3 2 0
After:  [1, 0, 3, 0]

Before: [1, 3, 0, 2]
4 0 2 1
After:  [1, 0, 0, 2]

Before: [1, 0, 2, 2]
5 0 2 2
After:  [1, 0, 0, 2]

Before: [1, 0, 2, 3]
5 0 2 0
After:  [0, 0, 2, 3]

Before: [0, 0, 3, 3]
13 3 2 1
After:  [0, 1, 3, 3]

Before: [0, 3, 1, 1]
6 0 0 2
After:  [0, 3, 0, 1]

Before: [3, 1, 2, 0]
15 1 2 2
After:  [3, 1, 0, 0]

Before: [1, 3, 3, 3]
13 3 2 0
After:  [1, 3, 3, 3]

Before: [1, 0, 2, 1]
5 0 2 3
After:  [1, 0, 2, 0]

Before: [3, 1, 3, 2]
9 1 2 0
After:  [0, 1, 3, 2]

Before: [0, 1, 3, 3]
12 1 0 1
After:  [0, 1, 3, 3]

Before: [3, 3, 0, 2]
0 2 3 2
After:  [3, 3, 1, 2]

Before: [0, 1, 3, 0]
12 1 0 2
After:  [0, 1, 1, 0]

Before: [3, 2, 2, 2]
11 2 2 3
After:  [3, 2, 2, 1]

Before: [3, 1, 3, 0]
9 1 2 3
After:  [3, 1, 3, 0]

Before: [0, 3, 2, 3]
8 2 2 2
After:  [0, 3, 2, 3]

Before: [2, 2, 1, 3]
2 2 3 0
After:  [0, 2, 1, 3]

Before: [2, 2, 2, 2]
8 2 2 0
After:  [2, 2, 2, 2]

Before: [0, 1, 3, 3]
2 1 3 0
After:  [0, 1, 3, 3]

Before: [1, 1, 3, 2]
1 1 0 2
After:  [1, 1, 1, 2]

Before: [0, 1, 2, 3]
15 1 2 1
After:  [0, 0, 2, 3]

Before: [1, 1, 2, 1]
5 0 2 2
After:  [1, 1, 0, 1]

Before: [3, 0, 3, 3]
13 3 2 2
After:  [3, 0, 1, 3]

Before: [1, 1, 2, 1]
14 1 3 0
After:  [1, 1, 2, 1]

Before: [0, 1, 1, 3]
12 1 0 3
After:  [0, 1, 1, 1]

Before: [1, 3, 2, 1]
8 2 2 3
After:  [1, 3, 2, 2]

Before: [1, 3, 2, 2]
5 0 2 0
After:  [0, 3, 2, 2]

Before: [0, 1, 1, 3]
12 1 0 2
After:  [0, 1, 1, 3]

Before: [1, 1, 2, 3]
15 1 2 1
After:  [1, 0, 2, 3]

Before: [1, 2, 2, 3]
5 0 2 2
After:  [1, 2, 0, 3]

Before: [3, 1, 0, 2]
10 3 3 2
After:  [3, 1, 0, 2]

Before: [2, 1, 0, 1]
14 1 3 3
After:  [2, 1, 0, 1]

Before: [2, 0, 1, 3]
2 2 3 0
After:  [0, 0, 1, 3]

Before: [1, 3, 3, 2]
10 3 3 2
After:  [1, 3, 0, 2]

Before: [0, 1, 2, 3]
15 1 2 2
After:  [0, 1, 0, 3]

Before: [3, 3, 1, 3]
2 2 3 2
After:  [3, 3, 0, 3]

Before: [1, 1, 1, 0]
1 1 0 1
After:  [1, 1, 1, 0]

Before: [2, 1, 3, 2]
7 3 2 2
After:  [2, 1, 1, 2]

Before: [1, 3, 1, 3]
2 2 3 2
After:  [1, 3, 0, 3]

Before: [2, 1, 3, 1]
9 1 2 3
After:  [2, 1, 3, 0]

Before: [1, 1, 0, 2]
0 2 3 0
After:  [1, 1, 0, 2]

Before: [3, 3, 1, 3]
13 3 0 1
After:  [3, 1, 1, 3]

Before: [1, 0, 0, 1]
4 0 2 3
After:  [1, 0, 0, 0]

Before: [2, 2, 2, 2]
13 2 0 1
After:  [2, 1, 2, 2]

Before: [0, 0, 3, 2]
6 0 0 0
After:  [0, 0, 3, 2]

Before: [1, 2, 0, 0]
4 0 2 3
After:  [1, 2, 0, 0]

Before: [1, 1, 0, 0]
1 1 0 2
After:  [1, 1, 1, 0]

Before: [3, 1, 2, 1]
15 1 2 1
After:  [3, 0, 2, 1]

Before: [0, 0, 3, 0]
0 3 2 3
After:  [0, 0, 3, 1]

Before: [2, 1, 3, 1]
14 1 3 1
After:  [2, 1, 3, 1]

Before: [0, 3, 3, 2]
6 0 0 3
After:  [0, 3, 3, 0]

Before: [1, 1, 1, 0]
1 1 0 0
After:  [1, 1, 1, 0]

Before: [0, 0, 0, 2]
11 0 0 3
After:  [0, 0, 0, 1]

Before: [3, 1, 0, 3]
2 1 3 0
After:  [0, 1, 0, 3]

Before: [3, 3, 3, 2]
7 3 2 2
After:  [3, 3, 1, 2]

Before: [1, 1, 3, 1]
1 1 0 3
After:  [1, 1, 3, 1]

Before: [2, 1, 3, 0]
9 1 2 1
After:  [2, 0, 3, 0]

Before: [2, 3, 0, 1]
7 2 0 0
After:  [1, 3, 0, 1]

Before: [0, 2, 3, 3]
13 3 2 2
After:  [0, 2, 1, 3]

Before: [0, 3, 1, 3]
2 2 3 2
After:  [0, 3, 0, 3]

Before: [2, 3, 2, 3]
11 3 2 0
After:  [0, 3, 2, 3]

Before: [1, 0, 1, 3]
2 2 3 2
After:  [1, 0, 0, 3]

Before: [3, 1, 1, 1]
14 1 3 0
After:  [1, 1, 1, 1]

Before: [1, 1, 2, 1]
1 1 0 3
After:  [1, 1, 2, 1]

Before: [1, 2, 2, 1]
3 3 2 0
After:  [1, 2, 2, 1]

Before: [0, 0, 3, 2]
7 3 2 0
After:  [1, 0, 3, 2]

Before: [1, 1, 2, 3]
5 0 2 2
After:  [1, 1, 0, 3]

Before: [3, 1, 2, 1]
3 3 2 1
After:  [3, 1, 2, 1]

Before: [1, 1, 1, 1]
1 1 0 1
After:  [1, 1, 1, 1]

Before: [3, 2, 2, 3]
8 2 2 1
After:  [3, 2, 2, 3]

Before: [3, 1, 1, 1]
14 1 3 3
After:  [3, 1, 1, 1]

Before: [2, 1, 3, 2]
10 3 3 2
After:  [2, 1, 0, 2]

Before: [0, 1, 0, 3]
6 0 0 1
After:  [0, 0, 0, 3]

Before: [0, 1, 2, 2]
12 1 0 2
After:  [0, 1, 1, 2]

Before: [0, 1, 2, 1]
15 1 2 1
After:  [0, 0, 2, 1]

Before: [2, 0, 2, 1]
3 3 2 1
After:  [2, 1, 2, 1]

Before: [0, 1, 2, 1]
12 1 0 3
After:  [0, 1, 2, 1]

Before: [1, 1, 2, 2]
10 3 3 2
After:  [1, 1, 0, 2]

Before: [2, 0, 3, 3]
8 3 3 2
After:  [2, 0, 3, 3]

Before: [1, 2, 0, 3]
4 0 2 1
After:  [1, 0, 0, 3]

Before: [1, 2, 0, 1]
4 0 2 2
After:  [1, 2, 0, 1]

Before: [0, 0, 3, 0]
0 3 2 2
After:  [0, 0, 1, 0]

Before: [1, 3, 3, 0]
0 3 2 0
After:  [1, 3, 3, 0]

Before: [0, 1, 3, 3]
2 1 3 2
After:  [0, 1, 0, 3]

Before: [0, 1, 3, 2]
6 0 0 3
After:  [0, 1, 3, 0]

Before: [1, 1, 1, 3]
1 1 0 2
After:  [1, 1, 1, 3]

Before: [1, 0, 0, 3]
4 0 2 0
After:  [0, 0, 0, 3]

Before: [3, 1, 3, 0]
9 1 2 2
After:  [3, 1, 0, 0]

Before: [1, 1, 1, 2]
1 1 0 0
After:  [1, 1, 1, 2]

Before: [1, 1, 2, 0]
5 0 2 0
After:  [0, 1, 2, 0]

Before: [0, 0, 0, 1]
6 0 0 3
After:  [0, 0, 0, 0]

Before: [0, 1, 2, 3]
11 3 2 1
After:  [0, 0, 2, 3]

Before: [1, 2, 2, 3]
5 0 2 3
After:  [1, 2, 2, 0]

Before: [0, 3, 2, 1]
6 0 0 3
After:  [0, 3, 2, 0]

Before: [0, 1, 3, 0]
9 1 2 2
After:  [0, 1, 0, 0]

Before: [1, 1, 1, 1]
10 2 3 1
After:  [1, 0, 1, 1]

Before: [1, 0, 2, 3]
8 3 3 3
After:  [1, 0, 2, 3]

Before: [2, 1, 2, 1]
8 2 2 0
After:  [2, 1, 2, 1]

Before: [2, 0, 0, 1]
7 2 0 2
After:  [2, 0, 1, 1]

Before: [1, 0, 2, 3]
5 0 2 2
After:  [1, 0, 0, 3]

Before: [1, 1, 3, 3]
13 3 2 3
After:  [1, 1, 3, 1]

Before: [1, 1, 0, 3]
1 1 0 0
After:  [1, 1, 0, 3]

Before: [2, 2, 2, 3]
13 2 1 3
After:  [2, 2, 2, 1]

Before: [1, 3, 2, 3]
5 0 2 1
After:  [1, 0, 2, 3]

Before: [1, 2, 2, 3]
11 2 2 0
After:  [1, 2, 2, 3]

Before: [3, 3, 2, 3]
8 3 3 1
After:  [3, 3, 2, 3]

Before: [2, 1, 2, 1]
14 1 3 1
After:  [2, 1, 2, 1]

Before: [1, 1, 2, 1]
3 3 2 1
After:  [1, 1, 2, 1]

Before: [2, 1, 2, 2]
8 2 2 0
After:  [2, 1, 2, 2]

Before: [3, 0, 3, 0]
0 3 2 2
After:  [3, 0, 1, 0]

Before: [1, 1, 3, 3]
9 1 2 1
After:  [1, 0, 3, 3]

Before: [1, 1, 0, 1]
1 1 0 2
After:  [1, 1, 1, 1]

Before: [3, 0, 0, 2]
0 2 3 2
After:  [3, 0, 1, 2]

Before: [1, 0, 0, 3]
4 0 2 1
After:  [1, 0, 0, 3]

Before: [2, 0, 3, 1]
10 3 3 3
After:  [2, 0, 3, 0]

Before: [0, 0, 0, 2]
0 2 3 1
After:  [0, 1, 0, 2]

Before: [0, 2, 2, 3]
2 1 3 2
After:  [0, 2, 0, 3]

Before: [2, 1, 3, 3]
9 1 2 0
After:  [0, 1, 3, 3]

Before: [1, 1, 2, 0]
5 0 2 2
After:  [1, 1, 0, 0]

Before: [0, 1, 1, 1]
12 1 0 0
After:  [1, 1, 1, 1]

Before: [2, 3, 3, 0]
0 3 2 2
After:  [2, 3, 1, 0]

Before: [2, 2, 0, 3]
7 2 0 1
After:  [2, 1, 0, 3]

Before: [2, 1, 0, 0]
7 2 0 0
After:  [1, 1, 0, 0]

Before: [0, 1, 1, 3]
11 3 2 2
After:  [0, 1, 0, 3]

Before: [3, 2, 1, 0]
7 3 0 3
After:  [3, 2, 1, 1]

Before: [1, 1, 2, 3]
15 1 2 2
After:  [1, 1, 0, 3]

Before: [1, 1, 1, 2]
1 1 0 3
After:  [1, 1, 1, 1]

Before: [3, 1, 2, 0]
15 1 2 0
After:  [0, 1, 2, 0]

Before: [3, 1, 2, 1]
14 1 3 1
After:  [3, 1, 2, 1]

Before: [2, 1, 3, 3]
9 1 2 2
After:  [2, 1, 0, 3]

Before: [2, 3, 2, 1]
3 3 2 3
After:  [2, 3, 2, 1]

Before: [1, 2, 3, 3]
13 3 2 3
After:  [1, 2, 3, 1]

Before: [3, 0, 2, 1]
10 3 3 1
After:  [3, 0, 2, 1]

Before: [1, 1, 2, 2]
11 2 2 2
After:  [1, 1, 1, 2]

Before: [0, 2, 2, 2]
11 2 2 3
After:  [0, 2, 2, 1]

Before: [3, 1, 1, 0]
7 3 0 1
After:  [3, 1, 1, 0]

Before: [1, 1, 0, 2]
4 0 2 3
After:  [1, 1, 0, 0]

Before: [2, 1, 0, 1]
14 1 3 0
After:  [1, 1, 0, 1]

Before: [2, 0, 2, 2]
10 3 3 1
After:  [2, 0, 2, 2]

Before: [0, 3, 0, 2]
6 0 0 2
After:  [0, 3, 0, 2]

Before: [0, 0, 2, 1]
10 3 3 0
After:  [0, 0, 2, 1]

Before: [2, 1, 1, 3]
2 1 3 1
After:  [2, 0, 1, 3]

Before: [0, 1, 0, 2]
12 1 0 1
After:  [0, 1, 0, 2]

Before: [0, 2, 1, 0]
11 0 0 2
After:  [0, 2, 1, 0]

Before: [2, 2, 2, 3]
2 2 3 1
After:  [2, 0, 2, 3]

Before: [2, 3, 0, 2]
10 3 3 0
After:  [0, 3, 0, 2]

Before: [2, 0, 2, 1]
3 3 2 3
After:  [2, 0, 2, 1]

Before: [3, 1, 3, 0]
0 3 2 2
After:  [3, 1, 1, 0]

Before: [3, 1, 2, 1]
3 3 2 0
After:  [1, 1, 2, 1]

Before: [0, 1, 2, 3]
15 1 2 0
After:  [0, 1, 2, 3]

Before: [0, 1, 2, 2]
12 1 0 1
After:  [0, 1, 2, 2]

Before: [3, 0, 3, 3]
13 3 2 3
After:  [3, 0, 3, 1]

Before: [1, 1, 3, 0]
1 1 0 0
After:  [1, 1, 3, 0]

Before: [3, 2, 2, 3]
2 1 3 1
After:  [3, 0, 2, 3]

Before: [3, 1, 3, 3]
9 1 2 2
After:  [3, 1, 0, 3]

Before: [2, 2, 1, 3]
8 3 3 2
After:  [2, 2, 3, 3]

Before: [2, 1, 3, 2]
11 2 1 0
After:  [0, 1, 3, 2]

Before: [3, 0, 3, 0]
0 3 2 1
After:  [3, 1, 3, 0]

Before: [2, 1, 2, 3]
11 3 2 2
After:  [2, 1, 0, 3]

Before: [1, 0, 3, 0]
0 3 2 2
After:  [1, 0, 1, 0]

Before: [1, 2, 2, 2]
5 0 2 0
After:  [0, 2, 2, 2]

Before: [0, 1, 1, 3]
11 3 1 3
After:  [0, 1, 1, 0]

Before: [1, 1, 3, 3]
8 3 3 1
After:  [1, 3, 3, 3]

Before: [3, 1, 0, 0]
7 3 0 1
After:  [3, 1, 0, 0]

Before: [1, 1, 2, 0]
1 1 0 1
After:  [1, 1, 2, 0]

Before: [3, 1, 3, 3]
9 1 2 1
After:  [3, 0, 3, 3]

Before: [1, 2, 2, 0]
5 0 2 1
After:  [1, 0, 2, 0]

Before: [2, 1, 1, 1]
14 1 3 2
After:  [2, 1, 1, 1]

Before: [0, 1, 1, 3]
6 0 0 1
After:  [0, 0, 1, 3]

Before: [3, 1, 3, 1]
13 2 3 0
After:  [0, 1, 3, 1]

Before: [1, 0, 2, 1]
3 3 2 2
After:  [1, 0, 1, 1]

Before: [2, 2, 3, 0]
0 3 2 3
After:  [2, 2, 3, 1]

Before: [2, 2, 0, 2]
7 2 0 0
After:  [1, 2, 0, 2]

Before: [3, 3, 3, 0]
0 3 2 1
After:  [3, 1, 3, 0]

Before: [0, 1, 3, 0]
0 3 2 1
After:  [0, 1, 3, 0]

Before: [0, 2, 2, 2]
10 3 3 1
After:  [0, 0, 2, 2]

Before: [1, 2, 3, 1]
13 2 3 3
After:  [1, 2, 3, 0]

Before: [1, 1, 2, 3]
11 2 2 2
After:  [1, 1, 1, 3]

Before: [3, 0, 2, 0]
11 2 2 3
After:  [3, 0, 2, 1]

Before: [3, 2, 3, 0]
0 3 2 3
After:  [3, 2, 3, 1]

Before: [1, 1, 3, 1]
1 1 0 1
After:  [1, 1, 3, 1]

Before: [1, 3, 0, 1]
4 0 2 2
After:  [1, 3, 0, 1]

Before: [0, 3, 1, 1]
10 3 3 3
After:  [0, 3, 1, 0]

Before: [0, 2, 2, 0]
8 2 2 1
After:  [0, 2, 2, 0]

Before: [0, 1, 3, 1]
12 1 0 0
After:  [1, 1, 3, 1]

Before: [1, 1, 3, 0]
9 1 2 0
After:  [0, 1, 3, 0]

Before: [0, 1, 2, 2]
6 0 0 3
After:  [0, 1, 2, 0]

Before: [2, 2, 2, 3]
11 3 1 0
After:  [0, 2, 2, 3]

Before: [1, 1, 3, 0]
1 1 0 1
After:  [1, 1, 3, 0]

Before: [2, 3, 2, 1]
3 3 2 0
After:  [1, 3, 2, 1]

Before: [3, 0, 3, 3]
13 3 2 1
After:  [3, 1, 3, 3]

Before: [1, 0, 2, 0]
11 2 2 2
After:  [1, 0, 1, 0]

Before: [2, 2, 3, 0]
0 3 2 2
After:  [2, 2, 1, 0]

Before: [0, 0, 2, 3]
6 0 0 2
After:  [0, 0, 0, 3]

Before: [3, 1, 3, 1]
9 1 2 0
After:  [0, 1, 3, 1]

Before: [1, 2, 1, 1]
10 2 3 0
After:  [0, 2, 1, 1]

Before: [0, 0, 1, 0]
6 0 0 3
After:  [0, 0, 1, 0]

Before: [0, 1, 3, 2]
9 1 2 3
After:  [0, 1, 3, 0]

Before: [3, 2, 2, 3]
2 1 3 2
After:  [3, 2, 0, 3]

Before: [0, 3, 2, 1]
3 3 2 0
After:  [1, 3, 2, 1]

Before: [1, 0, 3, 0]
0 3 2 3
After:  [1, 0, 3, 1]

Before: [1, 1, 1, 3]
1 1 0 1
After:  [1, 1, 1, 3]

Before: [1, 0, 2, 0]
5 0 2 1
After:  [1, 0, 2, 0]

Before: [0, 0, 0, 2]
10 3 3 3
After:  [0, 0, 0, 0]

Before: [0, 1, 2, 2]
15 1 2 0
After:  [0, 1, 2, 2]

Before: [0, 0, 1, 3]
8 3 3 2
After:  [0, 0, 3, 3]

Before: [1, 1, 2, 1]
3 3 2 2
After:  [1, 1, 1, 1]

Before: [0, 1, 1, 1]
6 0 0 3
After:  [0, 1, 1, 0]

Before: [0, 2, 1, 1]
6 0 0 2
After:  [0, 2, 0, 1]

Before: [3, 2, 2, 1]
13 2 1 1
After:  [3, 1, 2, 1]

Before: [2, 2, 2, 1]
3 3 2 0
After:  [1, 2, 2, 1]

Before: [0, 1, 3, 1]
9 1 2 0
After:  [0, 1, 3, 1]

Before: [0, 1, 3, 3]
12 1 0 0
After:  [1, 1, 3, 3]

Before: [0, 1, 2, 3]
11 2 2 1
After:  [0, 1, 2, 3]

Before: [2, 1, 2, 2]
15 1 2 2
After:  [2, 1, 0, 2]

Before: [2, 1, 3, 1]
9 1 2 0
After:  [0, 1, 3, 1]

Before: [3, 1, 2, 3]
2 1 3 1
After:  [3, 0, 2, 3]

Before: [2, 0, 3, 3]
13 3 2 2
After:  [2, 0, 1, 3]

Before: [2, 0, 2, 3]
8 2 2 3
After:  [2, 0, 2, 2]

Before: [0, 1, 2, 1]
3 3 2 3
After:  [0, 1, 2, 1]

Before: [3, 1, 3, 0]
0 3 2 0
After:  [1, 1, 3, 0]

Before: [1, 1, 0, 0]
1 1 0 0
After:  [1, 1, 0, 0]

Before: [1, 1, 3, 2]
9 1 2 3
After:  [1, 1, 3, 0]

Before: [2, 0, 2, 2]
13 2 0 1
After:  [2, 1, 2, 2]

Before: [0, 1, 2, 1]
3 3 2 2
After:  [0, 1, 1, 1]

Before: [2, 1, 2, 1]
15 1 2 0
After:  [0, 1, 2, 1]

Before: [2, 1, 1, 3]
2 2 3 2
After:  [2, 1, 0, 3]

Before: [1, 1, 3, 2]
9 1 2 1
After:  [1, 0, 3, 2]

Before: [2, 1, 0, 3]
7 2 0 1
After:  [2, 1, 0, 3]

Before: [0, 1, 3, 2]
12 1 0 0
After:  [1, 1, 3, 2]

Before: [1, 0, 0, 1]
4 0 2 2
After:  [1, 0, 0, 1]

Before: [1, 0, 0, 0]
4 0 2 3
After:  [1, 0, 0, 0]

Before: [3, 1, 3, 2]
9 1 2 1
After:  [3, 0, 3, 2]

Before: [2, 2, 0, 0]
7 2 0 1
After:  [2, 1, 0, 0]

Before: [1, 1, 0, 1]
1 1 0 0
After:  [1, 1, 0, 1]

Before: [3, 1, 0, 1]
10 3 3 2
After:  [3, 1, 0, 1]

Before: [0, 2, 2, 1]
3 3 2 2
After:  [0, 2, 1, 1]

Before: [0, 3, 2, 0]
6 0 0 3
After:  [0, 3, 2, 0]

Before: [0, 3, 0, 2]
0 2 3 1
After:  [0, 1, 0, 2]

Before: [0, 2, 3, 3]
8 3 3 2
After:  [0, 2, 3, 3]

Before: [0, 1, 3, 0]
12 1 0 1
After:  [0, 1, 3, 0]

Before: [0, 1, 1, 1]
14 1 3 2
After:  [0, 1, 1, 1]

Before: [1, 3, 0, 3]
4 0 2 0
After:  [0, 3, 0, 3]

Before: [3, 1, 3, 3]
13 3 0 3
After:  [3, 1, 3, 1]

Before: [0, 2, 3, 0]
6 0 0 1
After:  [0, 0, 3, 0]

Before: [1, 1, 2, 3]
15 1 2 3
After:  [1, 1, 2, 0]

Before: [0, 2, 0, 1]
10 3 3 1
After:  [0, 0, 0, 1]

Before: [1, 2, 0, 2]
4 0 2 1
After:  [1, 0, 0, 2]

Before: [0, 1, 1, 2]
10 3 3 2
After:  [0, 1, 0, 2]

Before: [3, 1, 1, 3]
2 2 3 2
After:  [3, 1, 0, 3]

Before: [1, 1, 1, 3]
8 3 3 2
After:  [1, 1, 3, 3]

Before: [3, 2, 2, 1]
3 3 2 1
After:  [3, 1, 2, 1]

Before: [1, 1, 3, 0]
9 1 2 3
After:  [1, 1, 3, 0]

Before: [3, 0, 2, 1]
3 3 2 0
After:  [1, 0, 2, 1]

Before: [0, 1, 0, 2]
12 1 0 2
After:  [0, 1, 1, 2]

Before: [1, 1, 1, 1]
14 1 3 1
After:  [1, 1, 1, 1]

Before: [0, 1, 0, 1]
14 1 3 2
After:  [0, 1, 1, 1]

Before: [2, 2, 2, 3]
8 2 2 3
After:  [2, 2, 2, 2]

Before: [0, 1, 3, 0]
0 3 2 2
After:  [0, 1, 1, 0]

Before: [3, 3, 1, 3]
2 2 3 3
After:  [3, 3, 1, 0]

Before: [0, 1, 2, 3]
2 1 3 0
After:  [0, 1, 2, 3]

Before: [0, 1, 3, 1]
13 2 3 2
After:  [0, 1, 0, 1]

Before: [3, 2, 3, 1]
10 3 3 1
After:  [3, 0, 3, 1]

Before: [0, 1, 2, 3]
12 1 0 2
After:  [0, 1, 1, 3]

Before: [1, 1, 3, 1]
14 1 3 0
After:  [1, 1, 3, 1]

Before: [1, 1, 0, 0]
4 0 2 0
After:  [0, 1, 0, 0]

Before: [1, 2, 3, 0]
0 3 2 3
After:  [1, 2, 3, 1]

Before: [0, 1, 0, 1]
12 1 0 3
After:  [0, 1, 0, 1]

Before: [0, 1, 1, 2]
12 1 0 0
After:  [1, 1, 1, 2]

Before: [0, 1, 3, 0]
0 3 2 3
After:  [0, 1, 3, 1]

Before: [3, 2, 3, 2]
7 3 2 2
After:  [3, 2, 1, 2]

Before: [3, 0, 3, 0]
0 3 2 0
After:  [1, 0, 3, 0]

Before: [1, 0, 2, 2]
5 0 2 3
After:  [1, 0, 2, 0]

Before: [0, 1, 0, 1]
14 1 3 1
After:  [0, 1, 0, 1]

Before: [1, 2, 2, 1]
3 3 2 3
After:  [1, 2, 2, 1]

Before: [1, 1, 2, 2]
1 1 0 1
After:  [1, 1, 2, 2]

Before: [0, 1, 0, 3]
8 3 3 3
After:  [0, 1, 0, 3]

Before: [0, 3, 2, 0]
8 2 2 1
After:  [0, 2, 2, 0]

Before: [1, 1, 0, 0]
1 1 0 1
After:  [1, 1, 0, 0]

Before: [3, 1, 1, 0]
7 3 0 3
After:  [3, 1, 1, 1]

Before: [1, 0, 2, 3]
5 0 2 1
After:  [1, 0, 2, 3]

Before: [1, 3, 2, 1]
5 0 2 1
After:  [1, 0, 2, 1]

Before: [2, 3, 3, 3]
13 3 2 1
After:  [2, 1, 3, 3]

Before: [2, 3, 2, 1]
8 2 2 2
After:  [2, 3, 2, 1]

Before: [1, 1, 2, 1]
10 3 3 3
After:  [1, 1, 2, 0]

Before: [1, 1, 1, 3]
1 1 0 0
After:  [1, 1, 1, 3]

Before: [0, 1, 3, 1]
9 1 2 2
After:  [0, 1, 0, 1]

Before: [0, 3, 0, 3]
6 0 0 1
After:  [0, 0, 0, 3]

Before: [1, 0, 1, 2]
10 3 3 1
After:  [1, 0, 1, 2]

Before: [2, 1, 3, 0]
0 3 2 1
After:  [2, 1, 3, 0]

Before: [1, 0, 3, 0]
0 3 2 0
After:  [1, 0, 3, 0]

Before: [3, 1, 2, 0]
15 1 2 1
After:  [3, 0, 2, 0]

Before: [2, 3, 3, 0]
0 3 2 3
After:  [2, 3, 3, 1]

Before: [2, 2, 2, 0]
8 2 2 3
After:  [2, 2, 2, 2]

Before: [3, 2, 3, 2]
7 3 2 1
After:  [3, 1, 3, 2]

Before: [2, 1, 3, 2]
9 1 2 0
After:  [0, 1, 3, 2]

Before: [1, 1, 3, 1]
9 1 2 0
After:  [0, 1, 3, 1]

Before: [0, 1, 1, 3]
11 0 0 3
After:  [0, 1, 1, 1]

Before: [3, 2, 0, 3]
11 3 3 0
After:  [1, 2, 0, 3]

Before: [3, 2, 3, 1]
13 2 3 2
After:  [3, 2, 0, 1]

Before: [2, 0, 0, 1]
7 2 0 1
After:  [2, 1, 0, 1]

Before: [1, 3, 0, 2]
0 2 3 2
After:  [1, 3, 1, 2]

Before: [0, 2, 2, 2]
6 0 0 3
After:  [0, 2, 2, 0]

Before: [1, 0, 0, 2]
4 0 2 3
After:  [1, 0, 0, 0]

Before: [2, 1, 2, 1]
3 3 2 3
After:  [2, 1, 2, 1]

Before: [1, 1, 2, 0]
1 1 0 3
After:  [1, 1, 2, 1]

Before: [3, 3, 3, 3]
13 3 2 3
After:  [3, 3, 3, 1]

Before: [0, 3, 1, 1]
6 0 0 0
After:  [0, 3, 1, 1]

Before: [0, 2, 1, 2]
11 0 0 1
After:  [0, 1, 1, 2]

Before: [1, 1, 0, 3]
1 1 0 2
After:  [1, 1, 1, 3]

Before: [0, 2, 2, 3]
6 0 0 0
After:  [0, 2, 2, 3]

Before: [1, 1, 3, 0]
1 1 0 2
After:  [1, 1, 1, 0]

Before: [1, 1, 2, 3]
1 1 0 2
After:  [1, 1, 1, 3]

Before: [1, 1, 3, 3]
13 3 2 0
After:  [1, 1, 3, 3]

Before: [1, 1, 3, 3]
1 1 0 1
After:  [1, 1, 3, 3]

Before: [1, 1, 3, 3]
1 1 0 2
After:  [1, 1, 1, 3]

Before: [1, 2, 0, 1]
4 0 2 3
After:  [1, 2, 0, 0]

Before: [0, 1, 2, 3]
12 1 0 0
After:  [1, 1, 2, 3]

Before: [3, 3, 0, 2]
0 2 3 0
After:  [1, 3, 0, 2]

Before: [2, 3, 0, 1]
7 2 0 3
After:  [2, 3, 0, 1]

Before: [0, 3, 3, 1]
13 2 3 2
After:  [0, 3, 0, 1]

Before: [3, 1, 2, 1]
3 3 2 2
After:  [3, 1, 1, 1]

Before: [0, 1, 2, 3]
2 1 3 3
After:  [0, 1, 2, 0]

Before: [1, 1, 1, 1]
14 1 3 2
After:  [1, 1, 1, 1]

Before: [1, 1, 0, 1]
14 1 3 0
After:  [1, 1, 0, 1]

Before: [1, 1, 0, 3]
11 3 1 2
After:  [1, 1, 0, 3]

Before: [2, 1, 2, 3]
2 2 3 3
After:  [2, 1, 2, 0]

Before: [0, 3, 2, 0]
8 2 2 3
After:  [0, 3, 2, 2]

Before: [0, 3, 0, 3]
6 0 0 3
After:  [0, 3, 0, 0]

Before: [1, 0, 2, 1]
5 0 2 2
After:  [1, 0, 0, 1]

Before: [3, 0, 3, 2]
10 3 3 2
After:  [3, 0, 0, 2]

Before: [1, 2, 3, 1]
13 2 3 0
After:  [0, 2, 3, 1]

Before: [3, 2, 1, 1]
10 2 3 2
After:  [3, 2, 0, 1]

Before: [3, 3, 2, 1]
3 3 2 2
After:  [3, 3, 1, 1]

Before: [2, 0, 1, 3]
2 2 3 1
After:  [2, 0, 1, 3]

Before: [2, 0, 2, 0]
11 2 2 0
After:  [1, 0, 2, 0]

Before: [1, 0, 0, 2]
4 0 2 0
After:  [0, 0, 0, 2]

Before: [1, 0, 0, 0]
4 0 2 1
After:  [1, 0, 0, 0]

Before: [1, 3, 2, 1]
5 0 2 3
After:  [1, 3, 2, 0]

Before: [0, 3, 1, 2]
11 0 0 1
After:  [0, 1, 1, 2]

Before: [1, 3, 2, 1]
3 3 2 3
After:  [1, 3, 2, 1]

Before: [3, 2, 3, 1]
13 2 3 0
After:  [0, 2, 3, 1]

Before: [3, 2, 1, 3]
2 2 3 2
After:  [3, 2, 0, 3]

Before: [0, 1, 0, 1]
14 1 3 0
After:  [1, 1, 0, 1]

Before: [3, 3, 2, 3]
8 3 3 0
After:  [3, 3, 2, 3]

Before: [1, 1, 2, 1]
5 0 2 1
After:  [1, 0, 2, 1]

Before: [0, 0, 2, 1]
3 3 2 0
After:  [1, 0, 2, 1]

Before: [2, 2, 0, 1]
7 2 0 3
After:  [2, 2, 0, 1]

Before: [1, 1, 0, 2]
1 1 0 3
After:  [1, 1, 0, 1]

Before: [1, 2, 0, 3]
4 0 2 2
After:  [1, 2, 0, 3]

Before: [0, 0, 2, 3]
8 2 2 1
After:  [0, 2, 2, 3]

Before: [1, 1, 0, 0]
4 0 2 1
After:  [1, 0, 0, 0]

Before: [2, 1, 2, 3]
15 1 2 1
After:  [2, 0, 2, 3]

Before: [2, 2, 1, 2]
10 3 3 2
After:  [2, 2, 0, 2]

Before: [0, 1, 2, 1]
15 1 2 3
After:  [0, 1, 2, 0]

Before: [2, 1, 2, 2]
10 3 3 1
After:  [2, 0, 2, 2]

Before: [3, 1, 0, 2]
10 3 3 0
After:  [0, 1, 0, 2]

Before: [1, 1, 3, 1]
1 1 0 2
After:  [1, 1, 1, 1]

Before: [2, 3, 2, 1]
3 3 2 2
After:  [2, 3, 1, 1]

Before: [1, 2, 2, 2]
5 0 2 1
After:  [1, 0, 2, 2]

Before: [1, 3, 2, 1]
3 3 2 0
After:  [1, 3, 2, 1]

Before: [0, 2, 2, 1]
6 0 0 1
After:  [0, 0, 2, 1]

Before: [1, 1, 2, 3]
1 1 0 3
After:  [1, 1, 2, 1]

Before: [0, 3, 2, 1]
3 3 2 3
After:  [0, 3, 2, 1]

Before: [2, 2, 2, 1]
13 2 1 1
After:  [2, 1, 2, 1]

Before: [0, 1, 3, 1]
12 1 0 3
After:  [0, 1, 3, 1]

Before: [0, 0, 3, 2]
6 0 0 1
After:  [0, 0, 3, 2]

Before: [1, 2, 2, 1]
3 3 2 2
After:  [1, 2, 1, 1]

Before: [0, 1, 2, 3]
6 0 0 1
After:  [0, 0, 2, 3]

Before: [3, 2, 2, 3]
8 2 2 3
After:  [3, 2, 2, 2]

Before: [2, 1, 1, 3]
8 3 3 3
After:  [2, 1, 1, 3]

Before: [1, 0, 1, 1]
10 2 3 0
After:  [0, 0, 1, 1]

Before: [0, 0, 3, 3]
8 3 3 3
After:  [0, 0, 3, 3]

Before: [0, 1, 0, 0]
12 1 0 0
After:  [1, 1, 0, 0]

Before: [2, 2, 1, 3]
11 3 2 1
After:  [2, 0, 1, 3]

Before: [1, 0, 3, 2]
7 3 2 1
After:  [1, 1, 3, 2]

Before: [2, 1, 2, 1]
15 1 2 3
After:  [2, 1, 2, 0]

Before: [3, 2, 0, 0]
7 3 0 2
After:  [3, 2, 1, 0]

Before: [1, 1, 2, 1]
3 3 2 0
After:  [1, 1, 2, 1]

Before: [2, 1, 3, 3]
8 3 3 2
After:  [2, 1, 3, 3]

Before: [1, 0, 2, 1]
3 3 2 1
After:  [1, 1, 2, 1]

Before: [2, 1, 1, 3]
2 1 3 2
After:  [2, 1, 0, 3]

Before: [0, 2, 2, 1]
3 3 2 3
After:  [0, 2, 2, 1]

Before: [2, 1, 0, 2]
7 2 0 0
After:  [1, 1, 0, 2]

Before: [3, 0, 2, 1]
3 3 2 2
After:  [3, 0, 1, 1]

Before: [3, 1, 2, 3]
15 1 2 2
After:  [3, 1, 0, 3]

Before: [1, 1, 2, 1]
14 1 3 2
After:  [1, 1, 1, 1]

Before: [0, 0, 3, 3]
8 3 3 0
After:  [3, 0, 3, 3]

Before: [2, 3, 0, 1]
10 3 3 1
After:  [2, 0, 0, 1]

Before: [1, 1, 0, 1]
14 1 3 3
After:  [1, 1, 0, 1]

Before: [1, 2, 2, 2]
13 2 1 1
After:  [1, 1, 2, 2]

Before: [1, 2, 0, 0]
4 0 2 0
After:  [0, 2, 0, 0]

Before: [3, 1, 0, 1]
10 3 3 3
After:  [3, 1, 0, 0]

Before: [1, 1, 3, 0]
0 3 2 0
After:  [1, 1, 3, 0]

Before: [3, 1, 3, 2]
9 1 2 3
After:  [3, 1, 3, 0]

Before: [1, 3, 3, 0]
0 3 2 3
After:  [1, 3, 3, 1]

Before: [1, 2, 0, 1]
4 0 2 0
After:  [0, 2, 0, 1]

Before: [0, 1, 3, 2]
7 3 2 0
After:  [1, 1, 3, 2]

Before: [0, 1, 2, 1]
10 3 3 1
After:  [0, 0, 2, 1]

Before: [1, 1, 3, 3]
1 1 0 3
After:  [1, 1, 3, 1]

Before: [2, 1, 3, 0]
9 1 2 2
After:  [2, 1, 0, 0]

Before: [0, 1, 3, 0]
11 0 0 2
After:  [0, 1, 1, 0]

Before: [1, 2, 2, 1]
5 0 2 3
After:  [1, 2, 2, 0]

Before: [0, 2, 2, 2]
6 0 0 2
After:  [0, 2, 0, 2]

Before: [0, 1, 2, 2]
12 1 0 3
After:  [0, 1, 2, 1]

Before: [1, 1, 0, 3]
4 0 2 3
After:  [1, 1, 0, 0]

Before: [1, 2, 0, 3]
2 1 3 3
After:  [1, 2, 0, 0]

Before: [2, 1, 2, 3]
15 1 2 0
After:  [0, 1, 2, 3]

Before: [3, 1, 3, 1]
14 1 3 2
After:  [3, 1, 1, 1]

Before: [0, 1, 3, 1]
6 0 0 3
After:  [0, 1, 3, 0]

Before: [1, 1, 0, 2]
1 1 0 2
After:  [1, 1, 1, 2]

Before: [0, 0, 0, 3]
8 3 3 1
After:  [0, 3, 0, 3]

Before: [2, 2, 3, 3]
2 1 3 3
After:  [2, 2, 3, 0]

Before: [0, 2, 1, 0]
6 0 0 2
After:  [0, 2, 0, 0]

Before: [3, 3, 2, 1]
3 3 2 1
After:  [3, 1, 2, 1]

Before: [0, 1, 2, 1]
3 3 2 0
After:  [1, 1, 2, 1]

Before: [3, 2, 2, 3]
13 2 1 0
After:  [1, 2, 2, 3]

Before: [1, 0, 2, 1]
3 3 2 3
After:  [1, 0, 2, 1]

Before: [1, 1, 2, 1]
1 1 0 1
After:  [1, 1, 2, 1]

Before: [1, 0, 3, 0]
0 3 2 1
After:  [1, 1, 3, 0]

Before: [2, 1, 2, 1]
3 3 2 1
After:  [2, 1, 2, 1]

Before: [2, 3, 2, 1]
8 2 2 0
After:  [2, 3, 2, 1]

Before: [3, 0, 3, 1]
13 2 3 1
After:  [3, 0, 3, 1]

Before: [0, 2, 2, 3]
11 2 2 2
After:  [0, 2, 1, 3]

Before: [0, 3, 1, 1]
6 0 0 1
After:  [0, 0, 1, 1]

Before: [0, 2, 3, 2]
6 0 0 3
After:  [0, 2, 3, 0]

Before: [1, 1, 2, 2]
1 1 0 2
After:  [1, 1, 1, 2]

Before: [2, 1, 3, 0]
9 1 2 0
After:  [0, 1, 3, 0]

Before: [0, 1, 2, 3]
8 3 3 3
After:  [0, 1, 2, 3]

Before: [1, 1, 3, 1]
9 1 2 3
After:  [1, 1, 3, 0]

Before: [3, 3, 1, 0]
7 3 0 3
After:  [3, 3, 1, 1]

Before: [0, 1, 0, 0]
12 1 0 1
After:  [0, 1, 0, 0]

Before: [0, 0, 1, 1]
6 0 0 3
After:  [0, 0, 1, 0]

Before: [1, 3, 2, 1]
3 3 2 1
After:  [1, 1, 2, 1]

Before: [1, 1, 3, 3]
11 3 1 0
After:  [0, 1, 3, 3]

Before: [0, 0, 0, 1]
6 0 0 1
After:  [0, 0, 0, 1]

Before: [2, 3, 0, 3]
7 2 0 2
After:  [2, 3, 1, 3]

Before: [1, 3, 1, 1]
10 2 3 0
After:  [0, 3, 1, 1]

Before: [1, 1, 0, 1]
1 1 0 1
After:  [1, 1, 0, 1]

Before: [1, 1, 2, 2]
15 1 2 2
After:  [1, 1, 0, 2]

Before: [1, 1, 0, 2]
0 2 3 1
After:  [1, 1, 0, 2]

Before: [3, 1, 2, 1]
14 1 3 3
After:  [3, 1, 2, 1]

Before: [3, 0, 3, 2]
7 3 2 3
After:  [3, 0, 3, 1]

Before: [0, 3, 2, 3]
2 2 3 3
After:  [0, 3, 2, 0]

Before: [1, 0, 0, 2]
0 2 3 2
After:  [1, 0, 1, 2]

Before: [2, 0, 1, 3]
2 2 3 2
After:  [2, 0, 0, 3]

Before: [1, 1, 1, 1]
14 1 3 0
After:  [1, 1, 1, 1]

Before: [1, 2, 2, 3]
2 1 3 0
After:  [0, 2, 2, 3]

Before: [0, 1, 0, 3]
12 1 0 1
After:  [0, 1, 0, 3]

Before: [3, 2, 3, 0]
7 3 0 0
After:  [1, 2, 3, 0]

Before: [1, 0, 2, 3]
8 2 2 1
After:  [1, 2, 2, 3]

Before: [0, 1, 1, 1]
12 1 0 2
After:  [0, 1, 1, 1]

Before: [1, 3, 2, 2]
5 0 2 2
After:  [1, 3, 0, 2]

Before: [0, 1, 3, 3]
9 1 2 3
After:  [0, 1, 3, 0]

Before: [2, 2, 0, 1]
7 2 0 1
After:  [2, 1, 0, 1]

Before: [2, 1, 3, 0]
9 1 2 3
After:  [2, 1, 3, 0]

Before: [1, 1, 0, 3]
4 0 2 0
After:  [0, 1, 0, 3]

Before: [1, 2, 3, 2]
10 3 3 0
After:  [0, 2, 3, 2]

Before: [1, 1, 0, 1]
1 1 0 3
After:  [1, 1, 0, 1]

Before: [0, 2, 2, 1]
10 3 3 0
After:  [0, 2, 2, 1]

Before: [0, 2, 1, 3]
2 1 3 0
After:  [0, 2, 1, 3]

Before: [1, 1, 0, 1]
4 0 2 1
After:  [1, 0, 0, 1]

Before: [2, 2, 3, 3]
2 1 3 1
After:  [2, 0, 3, 3]

Before: [1, 0, 2, 0]
5 0 2 3
After:  [1, 0, 2, 0]

Before: [3, 1, 1, 3]
2 1 3 1
After:  [3, 0, 1, 3]

Before: [2, 3, 0, 2]
0 2 3 2
After:  [2, 3, 1, 2]

Before: [0, 1, 1, 1]
14 1 3 0
After:  [1, 1, 1, 1]

Before: [0, 2, 3, 3]
6 0 0 3
After:  [0, 2, 3, 0]

Before: [0, 1, 3, 1]
14 1 3 3
After:  [0, 1, 3, 1]

Before: [0, 1, 2, 0]
15 1 2 2
After:  [0, 1, 0, 0]

Before: [0, 1, 1, 0]
12 1 0 2
After:  [0, 1, 1, 0]

Before: [3, 2, 2, 1]
3 3 2 2
After:  [3, 2, 1, 1]

Before: [1, 3, 0, 3]
4 0 2 1
After:  [1, 0, 0, 3]

Before: [0, 1, 3, 1]
14 1 3 1
After:  [0, 1, 3, 1]

Before: [1, 2, 0, 3]
4 0 2 0
After:  [0, 2, 0, 3]

Before: [1, 0, 3, 3]
8 3 3 3
After:  [1, 0, 3, 3]

Before: [1, 0, 0, 3]
4 0 2 2
After:  [1, 0, 0, 3]

Before: [1, 2, 2, 3]
11 3 1 3
After:  [1, 2, 2, 0]

Before: [1, 2, 2, 1]
5 0 2 1
After:  [1, 0, 2, 1]

Before: [0, 1, 2, 0]
15 1 2 0
After:  [0, 1, 2, 0]

Before: [0, 1, 2, 3]
15 1 2 3
After:  [0, 1, 2, 0]

Before: [1, 2, 3, 0]
0 3 2 1
After:  [1, 1, 3, 0]

Before: [1, 3, 0, 2]
4 0 2 2
After:  [1, 3, 0, 2]

Before: [1, 1, 2, 0]
15 1 2 0
After:  [0, 1, 2, 0]

Before: [3, 0, 2, 1]
3 3 2 1
After:  [3, 1, 2, 1]

Before: [2, 1, 1, 1]
14 1 3 3
After:  [2, 1, 1, 1]

Before: [0, 2, 3, 0]
0 3 2 3
After:  [0, 2, 3, 1]

Before: [0, 1, 0, 1]
14 1 3 3
After:  [0, 1, 0, 1]

Before: [2, 1, 2, 1]
14 1 3 0
After:  [1, 1, 2, 1]

Before: [1, 1, 1, 0]
1 1 0 3
After:  [1, 1, 1, 1]

Before: [0, 2, 2, 3]
8 3 3 2
After:  [0, 2, 3, 3]

Before: [1, 3, 0, 1]
4 0 2 1
After:  [1, 0, 0, 1]

Before: [2, 1, 3, 1]
11 2 1 0
After:  [0, 1, 3, 1]

Before: [1, 3, 0, 1]
4 0 2 0
After:  [0, 3, 0, 1]

Before: [0, 1, 1, 1]
14 1 3 1
After:  [0, 1, 1, 1]

Before: [1, 2, 0, 3]
8 3 3 1
After:  [1, 3, 0, 3]

Before: [1, 1, 2, 3]
1 1 0 1
After:  [1, 1, 2, 3]

Before: [1, 1, 1, 1]
14 1 3 3
After:  [1, 1, 1, 1]

Before: [0, 1, 0, 3]
6 0 0 0
After:  [0, 1, 0, 3]

Before: [1, 1, 2, 0]
15 1 2 3
After:  [1, 1, 2, 0]

Before: [0, 1, 1, 0]
12 1 0 3
After:  [0, 1, 1, 1]

Before: [2, 1, 2, 2]
15 1 2 3
After:  [2, 1, 2, 0]

Before: [1, 1, 3, 1]
9 1 2 2
After:  [1, 1, 0, 1]

Before: [1, 1, 0, 1]
4 0 2 0
After:  [0, 1, 0, 1]

Before: [0, 3, 1, 1]
10 2 3 0
After:  [0, 3, 1, 1]

Before: [1, 3, 2, 0]
5 0 2 2
After:  [1, 3, 0, 0]

Before: [1, 1, 3, 3]
9 1 2 0
After:  [0, 1, 3, 3]

Before: [1, 1, 2, 2]
5 0 2 1
After:  [1, 0, 2, 2]

Before: [0, 0, 3, 2]
7 3 2 1
After:  [0, 1, 3, 2]

Before: [0, 1, 3, 1]
9 1 2 1
After:  [0, 0, 3, 1]

Before: [3, 1, 0, 2]
0 2 3 3
After:  [3, 1, 0, 1]

Before: [0, 1, 0, 2]
12 1 0 0
After:  [1, 1, 0, 2]

Before: [0, 3, 2, 3]
8 3 3 1
After:  [0, 3, 2, 3]

Before: [2, 1, 2, 0]
13 2 0 0
After:  [1, 1, 2, 0]

Before: [3, 1, 1, 3]
13 3 0 1
After:  [3, 1, 1, 3]

Before: [1, 0, 2, 0]
11 2 2 0
After:  [1, 0, 2, 0]

Before: [1, 0, 1, 3]
2 2 3 0
After:  [0, 0, 1, 3]

Before: [1, 1, 2, 3]
2 1 3 0
After:  [0, 1, 2, 3]

Before: [0, 1, 3, 2]
12 1 0 1
After:  [0, 1, 3, 2]

Before: [0, 1, 3, 3]
9 1 2 2
After:  [0, 1, 0, 3]

Before: [3, 1, 3, 0]
0 3 2 1
After:  [3, 1, 3, 0]

Before: [3, 0, 2, 0]
7 3 0 0
After:  [1, 0, 2, 0]

Before: [0, 1, 1, 3]
2 2 3 0
After:  [0, 1, 1, 3]

Before: [1, 1, 3, 1]
13 2 3 2
After:  [1, 1, 0, 1]

Before: [1, 1, 2, 3]
5 0 2 0
After:  [0, 1, 2, 3]

Before: [0, 3, 2, 3]
11 0 0 3
After:  [0, 3, 2, 1]

Before: [1, 0, 2, 3]
2 2 3 2
After:  [1, 0, 0, 3]

Before: [1, 1, 0, 0]
1 1 0 3
After:  [1, 1, 0, 1]

Before: [1, 0, 0, 0]
4 0 2 2
After:  [1, 0, 0, 0]

Before: [0, 0, 1, 3]
11 3 2 1
After:  [0, 0, 1, 3]

Before: [0, 3, 2, 1]
3 3 2 1
After:  [0, 1, 2, 1]

Before: [2, 0, 2, 1]
13 2 0 2
After:  [2, 0, 1, 1]

Before: [1, 1, 0, 1]
14 1 3 2
After:  [1, 1, 1, 1]

Before: [0, 1, 2, 1]
10 3 3 3
After:  [0, 1, 2, 0]

Before: [0, 3, 2, 2]
10 3 3 3
After:  [0, 3, 2, 0]

Before: [3, 1, 3, 1]
9 1 2 1
After:  [3, 0, 3, 1]

Before: [1, 3, 0, 2]
4 0 2 3
After:  [1, 3, 0, 0]

Before: [0, 2, 2, 1]
6 0 0 0
After:  [0, 2, 2, 1]

Before: [1, 2, 2, 1]
5 0 2 2
After:  [1, 2, 0, 1]

Before: [0, 0, 2, 1]
3 3 2 2
After:  [0, 0, 1, 1]

Before: [2, 1, 3, 2]
9 1 2 3
After:  [2, 1, 3, 0]

Before: [1, 2, 2, 1]
8 2 2 0
After:  [2, 2, 2, 1]

Before: [3, 3, 2, 3]
11 2 2 2
After:  [3, 3, 1, 3]

Before: [1, 1, 0, 2]
4 0 2 2
After:  [1, 1, 0, 2]

Before: [0, 1, 1, 2]
6 0 0 2
After:  [0, 1, 0, 2]

Before: [1, 1, 2, 2]
15 1 2 3
After:  [1, 1, 2, 0]

Before: [2, 2, 1, 1]
10 3 3 0
After:  [0, 2, 1, 1]

Before: [2, 1, 1, 1]
14 1 3 1
After:  [2, 1, 1, 1]

Before: [3, 3, 0, 2]
0 2 3 3
After:  [3, 3, 0, 1]

Before: [1, 0, 3, 2]
7 3 2 3
After:  [1, 0, 3, 1]

Before: [2, 0, 3, 0]
0 3 2 3
After:  [2, 0, 3, 1]

Before: [0, 0, 1, 0]
6 0 0 1
After:  [0, 0, 1, 0]

Before: [1, 1, 3, 0]
9 1 2 2
After:  [1, 1, 0, 0]

Before: [0, 1, 2, 3]
11 0 0 2
After:  [0, 1, 1, 3]

Before: [1, 2, 1, 3]
2 1 3 1
After:  [1, 0, 1, 3]

Before: [2, 1, 2, 3]
15 1 2 3
After:  [2, 1, 2, 0]

Before: [0, 2, 2, 3]
8 2 2 0
After:  [2, 2, 2, 3]

Before: [1, 1, 2, 1]
3 3 2 3
After:  [1, 1, 2, 1]

Before: [1, 1, 2, 2]
15 1 2 1
After:  [1, 0, 2, 2]

Before: [1, 3, 0, 0]
4 0 2 2
After:  [1, 3, 0, 0]

Before: [1, 2, 2, 3]
2 1 3 1
After:  [1, 0, 2, 3]

Before: [0, 1, 0, 3]
12 1 0 3
After:  [0, 1, 0, 1]

Before: [1, 3, 2, 0]
5 0 2 1
After:  [1, 0, 2, 0]

Before: [1, 0, 2, 1]
5 0 2 1
After:  [1, 0, 2, 1]

Before: [1, 0, 2, 0]
5 0 2 2
After:  [1, 0, 0, 0]

Before: [2, 1, 2, 1]
14 1 3 3
After:  [2, 1, 2, 1]

Before: [1, 1, 3, 2]
1 1 0 0
After:  [1, 1, 3, 2]

Before: [0, 2, 2, 2]
6 0 0 0
After:  [0, 2, 2, 2]

Before: [0, 2, 0, 0]
6 0 0 0
After:  [0, 2, 0, 0]

Before: [1, 1, 3, 3]
1 1 0 0
After:  [1, 1, 3, 3]

Before: [1, 1, 3, 1]
14 1 3 1
After:  [1, 1, 3, 1]

Before: [3, 2, 2, 1]
3 3 2 0
After:  [1, 2, 2, 1]

Before: [0, 1, 2, 0]
12 1 0 0
After:  [1, 1, 2, 0]

Before: [1, 3, 2, 3]
5 0 2 0
After:  [0, 3, 2, 3]

Before: [0, 1, 3, 0]
9 1 2 1
After:  [0, 0, 3, 0]

Before: [1, 1, 2, 0]
5 0 2 1
After:  [1, 0, 2, 0]

Before: [3, 1, 2, 0]
7 3 0 1
After:  [3, 1, 2, 0]

Before: [1, 1, 3, 3]
2 1 3 2
After:  [1, 1, 0, 3]

Before: [2, 1, 3, 3]
9 1 2 1
After:  [2, 0, 3, 3]

Before: [1, 3, 2, 3]
5 0 2 3
After:  [1, 3, 2, 0]

Before: [1, 0, 2, 1]
3 3 2 0
After:  [1, 0, 2, 1]

Before: [1, 0, 0, 3]
4 0 2 3
After:  [1, 0, 0, 0]

Before: [0, 1, 2, 1]
11 0 0 0
After:  [1, 1, 2, 1]

Before: [1, 1, 2, 3]
1 1 0 0
After:  [1, 1, 2, 3]

Before: [2, 2, 2, 1]
3 3 2 1
After:  [2, 1, 2, 1]

Before: [1, 2, 3, 1]
13 2 3 2
After:  [1, 2, 0, 1]

Before: [2, 1, 2, 1]
3 3 2 0
After:  [1, 1, 2, 1]

Before: [1, 1, 2, 3]
2 1 3 1
After:  [1, 0, 2, 3]

Before: [0, 2, 2, 1]
3 3 2 0
After:  [1, 2, 2, 1]

Before: [3, 3, 2, 0]
7 3 0 0
After:  [1, 3, 2, 0]

Before: [1, 1, 1, 1]
1 1 0 0
After:  [1, 1, 1, 1]

Before: [1, 2, 0, 1]
4 0 2 1
After:  [1, 0, 0, 1]

Before: [1, 0, 2, 2]
5 0 2 1
After:  [1, 0, 2, 2]

Before: [1, 1, 0, 1]
4 0 2 3
After:  [1, 1, 0, 0]

Before: [2, 2, 0, 3]
8 3 3 3
After:  [2, 2, 0, 3]

Before: [1, 2, 0, 2]
4 0 2 2
After:  [1, 2, 0, 2]

Before: [0, 3, 3, 3]
13 3 2 3
After:  [0, 3, 3, 1]

Before: [2, 1, 2, 3]
13 2 0 3
After:  [2, 1, 2, 1]

Before: [2, 3, 0, 3]
8 3 3 0
After:  [3, 3, 0, 3]

Before: [1, 3, 2, 1]
5 0 2 0
After:  [0, 3, 2, 1]

Before: [2, 0, 2, 2]
8 2 2 3
After:  [2, 0, 2, 2]

Before: [1, 1, 0, 2]
4 0 2 0
After:  [0, 1, 0, 2]

Before: [3, 0, 1, 1]
10 2 3 3
After:  [3, 0, 1, 0]

Before: [1, 1, 3, 2]
9 1 2 0
After:  [0, 1, 3, 2]

Before: [1, 3, 0, 0]
4 0 2 0
After:  [0, 3, 0, 0]

Before: [3, 1, 1, 1]
14 1 3 2
After:  [3, 1, 1, 1]

Before: [3, 2, 3, 2]
10 3 3 2
After:  [3, 2, 0, 2]

Before: [2, 3, 1, 1]
10 2 3 3
After:  [2, 3, 1, 0]

Before: [1, 2, 2, 1]
5 0 2 0
After:  [0, 2, 2, 1]

Before: [1, 3, 0, 1]
4 0 2 3
After:  [1, 3, 0, 0]

Before: [0, 1, 3, 3]
12 1 0 3
After:  [0, 1, 3, 1]

Before: [3, 1, 3, 1]
14 1 3 3
After:  [3, 1, 3, 1]

Before: [2, 1, 2, 1]
3 3 2 2
After:  [2, 1, 1, 1]

Before: [2, 0, 3, 2]
7 3 2 2
After:  [2, 0, 1, 2]

Before: [0, 3, 2, 1]
3 3 2 2
After:  [0, 3, 1, 1]

Before: [1, 1, 0, 3]
1 1 0 3
After:  [1, 1, 0, 1]

Before: [0, 1, 1, 3]
12 1 0 1
After:  [0, 1, 1, 3]

Before: [0, 1, 3, 1]
12 1 0 2
After:  [0, 1, 1, 1]

Before: [2, 1, 3, 1]
9 1 2 1
After:  [2, 0, 3, 1]

Before: [1, 2, 2, 1]
3 3 2 1
After:  [1, 1, 2, 1]

Before: [2, 2, 0, 3]
2 1 3 1
After:  [2, 0, 0, 3]

Before: [1, 0, 0, 1]
4 0 2 0
After:  [0, 0, 0, 1]

Before: [3, 0, 2, 3]
8 3 3 3
After:  [3, 0, 2, 3]

Before: [1, 3, 2, 3]
5 0 2 2
After:  [1, 3, 0, 3]

Before: [3, 1, 3, 2]
9 1 2 2
After:  [3, 1, 0, 2]

Before: [1, 1, 0, 3]
1 1 0 1
After:  [1, 1, 0, 3]

Before: [1, 3, 0, 2]
4 0 2 0
After:  [0, 3, 0, 2]

Before: [0, 1, 2, 1]
12 1 0 1
After:  [0, 1, 2, 1]

Before: [1, 0, 0, 3]
8 3 3 0
After:  [3, 0, 0, 3]

Before: [3, 3, 0, 0]
7 3 0 1
After:  [3, 1, 0, 0]

Before: [1, 1, 3, 2]
7 3 2 2
After:  [1, 1, 1, 2]

Before: [2, 0, 1, 2]
10 3 3 1
After:  [2, 0, 1, 2]

Before: [1, 1, 0, 2]
1 1 0 1
After:  [1, 1, 0, 2]

Before: [3, 1, 0, 0]
7 3 0 2
After:  [3, 1, 1, 0]

Before: [2, 1, 0, 2]
0 2 3 3
After:  [2, 1, 0, 1]

Before: [1, 1, 1, 3]
1 1 0 3
After:  [1, 1, 1, 1]

Before: [1, 0, 2, 3]
5 0 2 3
After:  [1, 0, 2, 0]

Before: [2, 1, 3, 1]
14 1 3 2
After:  [2, 1, 1, 1]

Before: [3, 2, 0, 3]
11 3 1 0
After:  [0, 2, 0, 3]

Before: [2, 1, 2, 3]
8 2 2 0
After:  [2, 1, 2, 3]

Before: [0, 2, 0, 2]
11 0 0 0
After:  [1, 2, 0, 2]

Before: [2, 2, 0, 2]
10 3 3 1
After:  [2, 0, 0, 2]

Before: [1, 1, 3, 2]
1 1 0 1
After:  [1, 1, 3, 2]

Before: [2, 1, 2, 2]
15 1 2 0
After:  [0, 1, 2, 2]

Before: [2, 1, 2, 2]
10 3 3 2
After:  [2, 1, 0, 2]

Before: [0, 1, 3, 0]
12 1 0 0
After:  [1, 1, 3, 0]

Before: [1, 0, 3, 1]
10 3 3 0
After:  [0, 0, 3, 1]

Before: [0, 1, 2, 1]
3 3 2 1
After:  [0, 1, 2, 1]

Before: [2, 1, 1, 2]
10 3 3 3
After:  [2, 1, 1, 0]

Before: [0, 1, 3, 2]
7 3 2 2
After:  [0, 1, 1, 2]

Before: [1, 1, 1, 1]
1 1 0 2
After:  [1, 1, 1, 1]

Before: [2, 1, 2, 0]
15 1 2 3
After:  [2, 1, 2, 0]

Before: [1, 2, 0, 2]
4 0 2 3
After:  [1, 2, 0, 0]

Before: [1, 1, 3, 3]
9 1 2 2
After:  [1, 1, 0, 3]

Before: [3, 1, 1, 3]
13 3 0 3
After:  [3, 1, 1, 1]

Before: [0, 0, 1, 3]
2 2 3 1
After:  [0, 0, 1, 3]

Before: [3, 2, 0, 2]
0 2 3 3
After:  [3, 2, 0, 1]

Before: [1, 1, 1, 3]
2 1 3 1
After:  [1, 0, 1, 3]

Before: [1, 1, 3, 1]
9 1 2 1
After:  [1, 0, 3, 1]

Before: [1, 1, 3, 2]
1 1 0 3
After:  [1, 1, 3, 1]

Before: [1, 3, 2, 2]
10 3 3 1
After:  [1, 0, 2, 2]

Before: [0, 1, 3, 0]
9 1 2 0
After:  [0, 1, 3, 0]

Before: [3, 0, 2, 1]
3 3 2 3
After:  [3, 0, 2, 1]

Before: [3, 3, 2, 1]
3 3 2 0
After:  [1, 3, 2, 1]

Before: [2, 2, 2, 0]
13 2 1 1
After:  [2, 1, 2, 0]

Before: [2, 0, 3, 2]
7 3 2 3
After:  [2, 0, 3, 1]

Before: [1, 1, 1, 2]
1 1 0 2
After:  [1, 1, 1, 2]

Before: [1, 3, 3, 1]
10 3 3 0
After:  [0, 3, 3, 1]

Before: [2, 0, 2, 1]
8 2 2 1
After:  [2, 2, 2, 1]

Before: [1, 3, 2, 0]
5 0 2 0
After:  [0, 3, 2, 0]

Before: [1, 0, 2, 0]
5 0 2 0
After:  [0, 0, 2, 0]

Before: [1, 1, 2, 3]
5 0 2 1
After:  [1, 0, 2, 3]

Before: [0, 0, 3, 0]
0 3 2 1
After:  [0, 1, 3, 0]

Before: [1, 3, 0, 0]
4 0 2 1
After:  [1, 0, 0, 0]

Before: [0, 1, 3, 0]
12 1 0 3
After:  [0, 1, 3, 1]

Before: [3, 1, 1, 1]
14 1 3 1
After:  [3, 1, 1, 1]

Before: [0, 1, 2, 0]
15 1 2 3
After:  [0, 1, 2, 0]

Before: [3, 2, 1, 3]
13 3 0 0
After:  [1, 2, 1, 3]

Before: [3, 1, 2, 1]
15 1 2 0
After:  [0, 1, 2, 1]

Before: [1, 1, 2, 1]
1 1 0 2
After:  [1, 1, 1, 1]

Before: [0, 1, 1, 2]
11 0 0 1
After:  [0, 1, 1, 2]

Before: [3, 2, 0, 3]
8 3 3 0
After:  [3, 2, 0, 3]

Before: [2, 3, 2, 3]
11 2 2 2
After:  [2, 3, 1, 3]

Before: [0, 2, 3, 2]
7 3 2 1
After:  [0, 1, 3, 2]

Before: [0, 1, 2, 1]
8 2 2 1
After:  [0, 2, 2, 1]

Before: [0, 3, 1, 1]
6 0 0 3
After:  [0, 3, 1, 0]

Before: [0, 2, 3, 2]
7 3 2 3
After:  [0, 2, 3, 1]

Before: [0, 1, 2, 0]
6 0 0 3
After:  [0, 1, 2, 0]

Before: [0, 1, 0, 3]
6 0 0 2
After:  [0, 1, 0, 3]

Before: [1, 1, 3, 2]
9 1 2 2
After:  [1, 1, 0, 2]

Before: [0, 0, 2, 0]
11 2 2 1
After:  [0, 1, 2, 0]

Before: [3, 1, 3, 3]
9 1 2 3
After:  [3, 1, 3, 0]

Before: [0, 2, 3, 2]
6 0 0 2
After:  [0, 2, 0, 2]

Before: [3, 2, 2, 3]
13 3 0 0
After:  [1, 2, 2, 3]

Before: [0, 3, 3, 0]
6 0 0 3
After:  [0, 3, 3, 0]

Before: [0, 1, 1, 0]
12 1 0 0
After:  [1, 1, 1, 0]

Before: [1, 1, 3, 1]
14 1 3 2
After:  [1, 1, 1, 1]

Before: [3, 1, 2, 3]
15 1 2 0
After:  [0, 1, 2, 3]

Before: [1, 2, 2, 2]
5 0 2 2
After:  [1, 2, 0, 2]

Before: [3, 1, 1, 1]
10 2 3 1
After:  [3, 0, 1, 1]

Before: [2, 0, 2, 3]
8 3 3 3
After:  [2, 0, 2, 3]

Before: [2, 2, 2, 1]
3 3 2 3
After:  [2, 2, 2, 1]

Before: [1, 0, 0, 3]
8 3 3 1
After:  [1, 3, 0, 3]

Before: [3, 0, 2, 2]
10 3 3 2
After:  [3, 0, 0, 2]

Before: [3, 1, 2, 2]
15 1 2 0
After:  [0, 1, 2, 2]

Before: [1, 2, 2, 0]
5 0 2 3
After:  [1, 2, 2, 0]

Before: [3, 1, 0, 3]
13 3 0 1
After:  [3, 1, 0, 3]

Before: [0, 1, 0, 0]
12 1 0 3
After:  [0, 1, 0, 1]

Before: [1, 3, 3, 3]
13 3 2 3
After:  [1, 3, 3, 1]

Before: [0, 1, 2, 1]
12 1 0 2
After:  [0, 1, 1, 1]

Before: [3, 0, 2, 0]
7 3 0 2
After:  [3, 0, 1, 0]

Before: [1, 3, 2, 2]
5 0 2 3
After:  [1, 3, 2, 0]

Before: [1, 1, 2, 1]
1 1 0 0
After:  [1, 1, 2, 1]

Before: [3, 1, 3, 0]
9 1 2 1
After:  [3, 0, 3, 0]

Before: [3, 3, 0, 3]
13 3 0 1
After:  [3, 1, 0, 3]

Before: [2, 1, 2, 1]
15 1 2 1
After:  [2, 0, 2, 1]

Before: [3, 1, 3, 0]
9 1 2 0
After:  [0, 1, 3, 0]

Before: [1, 1, 2, 3]
5 0 2 3
After:  [1, 1, 2, 0]

Before: [0, 1, 0, 1]
12 1 0 1
After:  [0, 1, 0, 1]

Before: [2, 2, 2, 2]
13 2 1 1
After:  [2, 1, 2, 2]

Before: [1, 3, 0, 2]
0 2 3 1
After:  [1, 1, 0, 2]

Before: [2, 2, 0, 3]
8 3 3 1
After:  [2, 3, 0, 3]

Before: [2, 0, 0, 2]
0 2 3 2
After:  [2, 0, 1, 2]

Before: [0, 1, 3, 0]
6 0 0 0
After:  [0, 1, 3, 0]

Before: [0, 1, 0, 2]
0 2 3 1
After:  [0, 1, 0, 2]

Before: [0, 1, 0, 3]
11 3 3 3
After:  [0, 1, 0, 1]

Before: [1, 1, 3, 0]
0 3 2 1
After:  [1, 1, 3, 0]

Before: [1, 0, 0, 2]
4 0 2 2
After:  [1, 0, 0, 2]

Before: [2, 2, 2, 3]
13 2 0 0
After:  [1, 2, 2, 3]

Before: [0, 0, 0, 2]
6 0 0 1
After:  [0, 0, 0, 2]

Before: [0, 1, 3, 1]
9 1 2 3
After:  [0, 1, 3, 0]

Before: [1, 2, 2, 2]
5 0 2 3
After:  [1, 2, 2, 0]

Before: [1, 2, 2, 3]
13 2 1 1
After:  [1, 1, 2, 3]

Before: [1, 2, 1, 3]
2 2 3 3
After:  [1, 2, 1, 0]

Before: [0, 1, 2, 1]
15 1 2 0
After:  [0, 1, 2, 1]

Before: [1, 1, 2, 1]
15 1 2 0
After:  [0, 1, 2, 1]

Before: [3, 3, 2, 1]
8 2 2 0
After:  [2, 3, 2, 1]

Before: [1, 1, 0, 1]
4 0 2 2
After:  [1, 1, 0, 1]

Before: [2, 1, 2, 3]
2 1 3 1
After:  [2, 0, 2, 3]

Before: [0, 1, 3, 3]
11 0 0 2
After:  [0, 1, 1, 3]

Before: [0, 0, 2, 2]
11 2 2 3
After:  [0, 0, 2, 1]

Before: [1, 3, 2, 1]
5 0 2 2
After:  [1, 3, 0, 1]

Before: [0, 0, 0, 3]
11 0 0 2
After:  [0, 0, 1, 3]

Before: [3, 3, 2, 3]
13 3 0 0
After:  [1, 3, 2, 3]

Before: [2, 1, 3, 1]
14 1 3 3
After:  [2, 1, 3, 1]

Before: [3, 0, 2, 0]
11 2 2 0
After:  [1, 0, 2, 0]

Before: [0, 2, 0, 2]
6 0 0 1
After:  [0, 0, 0, 2]

Before: [3, 1, 2, 1]
14 1 3 0
After:  [1, 1, 2, 1]

Before: [1, 1, 1, 0]
1 1 0 2
After:  [1, 1, 1, 0]

Before: [0, 2, 2, 1]
3 3 2 1
After:  [0, 1, 2, 1]

Before: [0, 0, 2, 3]
11 2 2 2
After:  [0, 0, 1, 3]

Before: [3, 1, 1, 1]
10 3 3 2
After:  [3, 1, 0, 1]

Before: [2, 1, 0, 1]
14 1 3 1
After:  [2, 1, 0, 1]

Before: [1, 1, 2, 2]
15 1 2 0
After:  [0, 1, 2, 2]

Before: [1, 2, 1, 1]
10 3 3 1
After:  [1, 0, 1, 1]

Before: [1, 2, 0, 2]
0 2 3 0
After:  [1, 2, 0, 2]`
