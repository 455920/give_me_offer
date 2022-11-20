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


# 遍历nums, 将每个num保存在map中 num(key) -> index(value)
# 遍历时, 保存sub = target - 当前num, 判断如果sub是map中的key, 说明之前的数据中有一个值sub与当前num 和为target

if __name__ == "__main__":
    print(Solution().twoSum([2, 7, 11, 15], 9))
