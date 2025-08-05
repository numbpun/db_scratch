## Type of queries

Data structures dictates what kids of queries are supported.

Only small subset queries can benefit from data structures.

Real-world OLTP queries fall into 1 on 3:

1. Scan whole Data set without an index (only if table is small)

2. Point Query: Query index by a single key

3. Range Query: Query a range of keys in sort order

Non Indexed Queries are concern of OLAP
- Range Queries require DS to keep keys in sort order. DS support 2 operations:

1. Seek: Find starting key of rnage
2. Iterate: Visit previous/next key in sort order

Point Queries just "seek" without "iterate". So, a sorting data structure

HashTables only work for point query (Get, set, delete) since you do not need sorted keys.

But using Hash (even in-memory one) has valuable challenges:

1. How to efficiently grow a hastable? Keys must be moved to a larger hashrable when load facotr is too high. Moving Everything at once is O(N). Rehashing must he done progressively even for in-memory apps like Redis
2. In-place updates, space resuse etc.

# 2.3 Sorted Arrays

So, for sorted, lets use simplest sorting DS: Sorted Array.

Binary Search in `O(log N)`. For variable-length data such as strings (KV), use an array of pointers (offsets) to do binary searches.

Updating a sorted array is `O(n)` (either in place or not). Therefore not practical, but can be extended to other updatable data structures.

To reduce update costs:

1. Divide array into several non-overlapping arrays(nested sorted arrays) -> B+tree (multi-level n-ary tree) with extra issue of maintaining these small arrays (tree nodes)

2. Another "Updatable Array" is Log-Structured Merge Tree (LSM)
Updates first buffered in a smaller array (or other sorting data structures), then merged into main array when it becomes too large. The update cost is amortized by propagating smaller arrays into larger arrays.

# 2.4 B tree

Balanced n-ary tree comparable to balanced binary trees

Each node can store variable #keys (and branches) upto n and n>2

Reduces IOs per second (Limiting factor for tree lookups). Each level of a tree is a disk read in lookup

Large n is advantageous but tradeoffs:
- Large trees slower to update
- Read Latency increses for larger nodes.

Node size are chosen as only a few OS pages
