""" 146. LRU Cache
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
  get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
  put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, 
it should invalidate the least recently used item before inserting a new item.
The cache is initialized with a positive capacity. https://leetcode.com/problems/lru-cache/

Doubly Linked List to maintain the key and value pairs. 
  The one in the front is least recently used.
  The one in the rear is the most frequently used.
  Dict will have reference of each node for constant access. 

Faster than 87% of Python submissions
"""


class DListNode:
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.next = None
        self.prev = None
        
class LRUCache_separate_HashMap:
    def __init__(self, capacity: int):
        self.capacity = capacity # max number of cache items to keep
        self.nItems = 0 # number of cache items
        self.head = None # newest item (head of list)
        self.tail = None # oldest item (tail of list)
        self.dict = {} # dict of pointers to double-linked list
        
    def get(self, key: int) -> int:
        if not key in self.dict:
            return -1
        else:
            node = self.dict[key]
            print("GET", self.nItems, node.key, node.val, node.next, node.prev)
            self.move_to_end(node) # now most recently accessed node
            return node.val

    def put(self, key: int, value: int) -> None:
        node = None
        if self.nItems < self.capacity:
            # Allocate new node (at beginning)
            node = DListNode(key, value)
            node.next = self.head
            self.head = node
            if node.next != None:
                (node.next).prev = node
            if self.tail == None:
                self.tail = node
            self.nItems += 1
        else:
            # Cache eviction overwrite oldest node
            node = self.tail
            self.move_to_end(node)
            del self.dict[node.key] # delete old dict entry
        node.val = value
        self.dict[key] = node
        node.key = key # save key to enable delete from dict
        self.print_cache()
        
    def move_to_end(self, node):
        """ Move selected node to front of list (most recent) """
        if self.head == node: # node already at front
            return
        # a = node.prev
        # print("A", a)
        if node.next != None:
            (node.next).prev = node.prev
        if self.tail == node:
            self.tail = node.prev
        (node.prev).next = node.next
        node.next = self.head
        self.nead = node
        node.prev = None

    def print_cache(self):
        node = self.head
        while node != None:
            print ("NODE", node.key, node.val, node.prev, node.next)
            node = node.next
        print("TAIL", (self.tail).key)
        print()


""" TRIVIAL using collections.OrderedDict structure
  OrderedDict acts as Doubly linked list + Dict with the entries on the left being least recent to the entries on the right.
  When we call the get() function we move the entry if present to the right (make it more recent).
  Time complexity of move_to_end is O(1) since its implemented as a linked list.
"""
from collections import OrderedDict
class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity # max number of cache items to keep
        self.nItems = 0 # number of cache items
        self.dict = OrderedDict() # dict that keeps tract of order

    def get(self, key: int) -> int:
        if key in self.dict:
            val = self.dict[key]
            self.dict.move_to_end(key)
            return val
        else:
            return -1

    def put(self, key: int, value: int) -> None:
        if key in self.dict:
            self.dict[key] = value
            self.dict.move_to_end(key)
            return
        if len(self.dict) >= self.capacity:
            # Evict least-frequently used item
            self.dict.popitem(last=False)
        self.dict[key] = value
        
cache = LRUCache(2) # 2 = capacity
if True:
    cache.put(1, 1.1)
    cache.put(2, 2.2)
    print(cache.get(1)) # returns 1.1
    cache.put(3, 3.3)   # evicts key 2
    print(cache.get(2)) # assert -1 (not found)
    cache.put(4, 4.4)   # evicts key 1
    print(cache.get(1)) # assert -1 (not found)
    print(cache.get(3)) # assert 3.3
    print(cache.get(4)) # assert 4.4

if False:
    cache.put(2, 1)
    cache.put(2, 2)
    print(cache.get(2)) # returns 2
    cache.put(1, 1)   # evicts key 2
    cache.put(4, 1)   # evicts key 1
    print(cache.get(2)) # assert -1 (not found)


