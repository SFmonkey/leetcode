class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        p = {}
        for i, value in enumerate(nums):
        	diff = target - value
        	if diff in p:
        		return [p[diff], i]
        	p[value] = i
