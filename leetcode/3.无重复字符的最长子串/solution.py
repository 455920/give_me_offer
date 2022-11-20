class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        win_kv = set()
        max_len = 0
        win_left = 0
        for i in range(0, len(s)):
            if s[i] not in win_kv:
                win_kv.add(s[i])
            else:
                while True:
                    if s[win_left] == s[i]:
                        win_left += 1
                        break
                    else:
                        win_kv.remove(s[win_left])
                        win_left += 1
            max_len = max(max_len, i - win_left + 1)
        return max_len


if __name__ == "__main__":
    print(Solution().lengthOfLongestSubstring("123"))
    print(Solution().lengthOfLongestSubstring("121"))
    print(Solution().lengthOfLongestSubstring("1231"))
    print(Solution().lengthOfLongestSubstring("1231231"))

