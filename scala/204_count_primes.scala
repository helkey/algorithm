//  Number of prime numbers less than a non-negative number, n.

import scala.collection.mutable.ListBuffer
object Solution {
  def countPrimes(n: Int): Int = {
    var primes = new ListBuffer[Integer]()
    def isPrime(n: Int): Boolean = {
      for (p <- primes) {
        if (n % p == 0) return false
        if (p * p > n) return true
      }
      return true
    }
      (2 to n-1).foreach(i =>
        if (isPrime(i)) primes += i // append new prime value
  )
    return primes.length
  }
}
