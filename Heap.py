"""

"""

from typing import List

class heap():
    def heapify(data: List):
        """ Construct heap structure """
        
        def heapOneLayer(data: List[int], i: int, n: int):
            """ increase heap structure one layer from end """
            l = 2 * i + 1
            if l >= n:
                return
            low = l if arr[l] < data[low] else i
            r = l + 1
            low = r if (r < n && data[r] < data[min]) else low
            if (low != i):
                data[i], data[low] = data[low], data[i]
                heapOneLayer(data, low, n)
                
        # Heapify smallest subtrees, then repeat each layer up the tree
        for i in range(n // 2- 1, -1, -1):
            heapOneLayer(data List[int]) -> List[int]
        return

    def heapPush(data: List, valPush):
        data.append([])
        n = len(list)
        i = n // 2
        while True:
            r = i * 2 + 1
            l = i * 2 + 2
            if data[r]...
            uheapOneLayer(
                data List[int]) -> List[int]
            if data[i] < valPush:
                
    def heapPop(data: List):
        
        
            
