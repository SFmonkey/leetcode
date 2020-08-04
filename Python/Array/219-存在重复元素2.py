def containsNearbyDuplicate(nums: list[int], k: int) -> bool:
    arr = {}
    for i in range(len(nums)):
        if nums[i] in arr and i - arr[nums[i]] <= k:
            return True
        arr[nums[i]] = i
    return False

