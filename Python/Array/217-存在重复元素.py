def containsDuplicate(nums: list[int]) -> bool:
    stack = {}
    for num in nums:
        if num in stack:
            return True
        stack[num] = 1
    return False