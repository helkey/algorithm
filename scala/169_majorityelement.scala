// 169. Majority Element
// Faster than 96% of Scala submissions
//   https://leetcode.com/problems/majority-element/

"""
Boyer-Moore Voting Algorithm
Incremented whenever we see an instance of our current candidate for majority element and 
  decremented whenever we see anything else. Whenever count equals 0, we effectively forget about 
  everything in nums up to the current index and consider the current number as the candidate
"""

object Solution {
    def majorityElement(nums: Array[Int]): Int = {
      var count = 0
      var guess = 0
      nums.foreach{ n =>
        if (count == 0) guess = n
        if (n == guess) count += 1
        else count -= 1
      }
      return guess
    }
}

        
