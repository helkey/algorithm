// 242_valid_anagram.scala
// Faster than 78% of Scala submissions

import scala.collection.mutable.HashMap
object Solution {
    def isAnagram(s: String, t: String): Boolean = {
        if (s.length() != t.length()) return false
        var m:HashMap[Char, Int] = HashMap()
        s.foreach{ c=>
            if (m.contains(c)) {
                m(c) += 1
            } else {
                m(c) = 1
            }
        }
        t.foreach{ c=>
            if (!m.contains(c)) return false
            if (m(c)>0) {
                m(c) -= 1
            } else {
                return false
            }
        }
        return true
    }
}
