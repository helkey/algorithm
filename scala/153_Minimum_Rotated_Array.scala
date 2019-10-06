// 153. Find Minimum in Rotated Sorted Array
//  Find min element of array sorted in ascending order, then rotated at some pivot unknown to you beforehand.

object Solution {
  def findMin(nums: Array[Int]): Int = {
    // Check for special case of no rotation
    var nSt = 0
    var nE = nums.length - 1
    if (nums(nSt) < nums(nE)) return nSt
    while (nE > nSt + 1) {
      var nMid = (nSt + nE)/2
      if (nums(nSt) > nums(nMid)) nE = nMid
      else nSt = nMid + 1
    }
    return nums(nE)
  }
}

var nums = Array(3,4,5,1,2)
println(Solution.findMin(nums))

