# 76. Minimum Window Substring

https://leetcode.com/problems/minimum-window-substring/

# Difficulty:

Hard

# Description

Given two strings `s` and `t` of lengths `m` and `n` respectively, return the **minimum window substring** of `s` such that every character in `t` **(including duplicates)** is included in the window. If there is no such substring, return the empty string `""`.

The testcases will be generated such that the answer is `unique`.

Note: Leetcode and Educative differ for this question, Leetcode the minimum substring can be rearranged to match the substring t, while for Educative they do not allow for rearranging. I'll try both. This description and examples are from Leetcode.

Example 1:
```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
```

Example 2:
```
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
```

Example 3:
```
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
```


Constraints:

- `m == s.length`
- `n == t.length`
- <code>1 <= m, n <= 10<sup>5</sup></code>
- `s` and `t` consist of uppercase and lowercase English letters.

 

Follow up: Could you find an algorithm that runs in `O(m + n)` time?