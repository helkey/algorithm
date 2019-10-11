# 190. Reverse Bits
# https://leetcode.com/problems/reverse-bits/

# Faster than 28% of solutions
# Bit-shift mask solution
class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        r = 0
        m1 = 1
        m2 = 1 << 32
        for _ in range(32):
            b = n & m1
            if (b != 0):
                r = r | m2
            m1 << 1
            m2 >> 1
        return r
    
# test case:
  n = 4294967293
  r = 322122541
  print(reverseBits(n))          


"""
# Faster than 28% of solutions
# Bit-shifting solution
class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        r = 0
        for _ in range(32):
            r = (r << 1) | (n & 1)
            n = n >> 1
            # print(r, n)
        return r """
