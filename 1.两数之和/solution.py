class Solution:
    def twoSum(self, nums, target):
        kv = {}
        result = []
        for i in range(0, len(nums)):
            sub = target - nums[i]
            if sub in kv.keys():
                result.append(kv[sub])
                result.append(i)
                break
            kv[nums[i]] = i
        return result


if __name__ == "__main__":
    print(Solution().twoSum([2, 7, 11, 15], 9))
