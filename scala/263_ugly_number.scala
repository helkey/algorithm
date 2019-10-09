// 263. Ugly Number
// Faster than 77% of Scala solutions

// Enforce tail-call recursion
import scala.annotation.tailrec

@tailrec
object Solution {
    def isUgly(num: Int): Boolean = {
        if (num == 1) return true
        if (num <= 0) return false
        var mod = 0
        if (num % 2 == 0) mod = 2
        else if (num % 3 == 0) mod = 3
        else if (num % 5 == 0) mod = 5
        else return false
        return isUgly(num / mod)
    }
}
      
								        }
									}