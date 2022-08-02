# https://leetcode.com/problems/last-stone-weight/

from typing import List
import heapq as hq

def lastStoneWeight(self, stones: List[int]) -> int:
    for x in stones:
        if (len(stones) <= 1):
            return stones

        hq._heapify_max(stones)
        y = hq.heappop((stones)) 
        hq._heapify_max((stones))
        x = hq.heappop((stones)) 

        newelem = y - x
        hq.heappush(stones, newelem)

    if (len(stones) > 1):
        return lastStoneWeight(self = None, stones = stones)
    else:
        return stones



print(lastStoneWeight(self = None, stones = [2,7,4,1,8,1]))
print(lastStoneWeight(self = None, stones = [7,6,7,6,9]))


