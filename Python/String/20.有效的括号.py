# Solution 1
# 大神的解法，类似消消乐
class Solution:
    def isValid(self, s: str) -> bool:
        while '{}' in s or '[]' in s or '()' in s:
            s = s.replace('{}', '')
            s = s.replace('[]', '')
            s = s.replace('()', '')
        return s == ''

# Solution 2
# 用栈
class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        left = {'{':'}', '[':']', '(':')'}
        for c in s:
            if c in left:
                stack.append(left[c])
            elif (len(stack) == 0 or c != stack.pop()):
                return False

        return True if len(stack) == 0 else False 
