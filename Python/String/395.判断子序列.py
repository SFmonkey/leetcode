class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) > len(t):
            return False
        if len(s) == 0:
            return True
        s_, t_ = 0, 0
        while t_ <= len(t)-1:
            if s[s_] == t[t_]:
                s_ += 1
                if s_ == len(s):
                    return True
            t_ += 1
        if s_ < len(s):
            return False
        return True
