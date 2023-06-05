# Least Recently Used (LRU) Cache
Design a data structure that follows the constraints of a [Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU).

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initialize the LRU cache with positive size `capacity`.
- `int get(String key)` Return the value of the key if the key exists, otherwise return -1.
- `void put(String key, int value)` Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.


### Sample Usage
```java
LRUCache lRUCache = new LRUCache(2);
lRUCache.put("a", 1); // cache is {a=1}
lRUCache.put("b", 2); // cache is {a=1, b=2}
lRUCache.get("a");    // return 1
lRUCache.put("c", 3); // LRU key was "b", evicts key "b", cache is {a=1, c=3}
lRUCache.get("b");    // returns -1 (not found)
lRUCache.put("d", 4); // LRU key was "a", evicts key "a", cache is {"d"=4, "c"=3}
lRUCache.get("a");    // return -1 (not found)
lRUCache.get("c");    // return 3
lRUCache.get("d");    // return 4
```

