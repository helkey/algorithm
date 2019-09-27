# 412. Fizz Buzz
# Faster than 85% of Python submissions

from typing import List

class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        def fizzBuzz1(n: int):
            fizz = not(n%3)
            buzz = not(n%5)
            if not fizz and not buzz:
                return str(n)
            fizz = "Fizz" if fizz else ""
            buzz = "Buzz" if buzz else ""
            return fizz + buzz
        return [fizzBuzz1(i + 1) for i in range(n)]

            
s = Solution()
print(s.fizzBuzz(15))
        
