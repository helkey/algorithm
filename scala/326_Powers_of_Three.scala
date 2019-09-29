object Solution1 {
    def isPowerOfThree(n: Int): Boolean = {
      var i = n
      if (i < 1) return false
      while (i % 3 == 0) i = i/3
      return i == 1
    }
}

object Main {
}

