class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        if len(intervals) == 0:
            return 0
        if len(intervals) == 1:
            return 1
        on = [i[0] for i in intervals]
        down = [i[1] for i in intervals]
        on.sort()
        down.sort()
        max_num = 0
        num = 0
        i = 0
        j = 0
        while i < len(on):
            num += 1
            max_num = max(num, max_num)
            if i+1 >= len(on):
                return max_num
            if down[j] <= on[i+1]:
                num -= 1
                j += 1
            i += 1
        return max_num

# 别人的解法，把上下车时间统一排序
class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        events = [(iv[0], 1) for iv in intervals] + [(iv[1], -1) for iv in intervals]
        events.sort()
        ans = cur = 0
        for _, e in events:
            cur += e
            ans = max(ans, cur)
        return ans
