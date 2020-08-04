# 时间复杂度O(n), 空间复杂度O(n)
def findDuplicate(nums: list[int]) -> list:
    hm = {}
    for n in nums:
        if n not in hm:
            hm[n] = 0
        hm[n] += 1
        if hm[n] > 1:
            res.append(n)
    return res

# 时间复杂度O(n), 空间复杂度O(1)
# 环形链表快慢指针找环入口
def findDuplicate2(nums: list[int]) -> list:
    slow = nums[0]
    fast = nums[nums[0]]
    while slow != fast:
        slow = nums[slow]
        fast = nums[nums[fast]]
    root = 0
    while root != slow:
        root = nums[root]
        slow = nums[slow]
    return root