# algo

## max_subarray_sum:
<details> 
    <summary>Maximum subarray.</summary>

    Given an integer array nums, find the subarray with the largest sum, and return its sum.

    ## Example 1:
    Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6
    Explanation: The subarray [4,-1,2,1] has the largest sum 6.
    
    ## Example 2:
    Input: nums = [1]
    Output: 1
    Explanation: The subarray [1] has the largest sum 1.

    ## Example 3:
    Input: nums = [5,4,-1,7,8]
    Output: 23
    Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
</details>

## valid_parentheses:
<details> 
    <summary>Valid parentheses.</summary>

    Given a string s containing just the characters '(', ')', '{', '}', '[' and '] determine if the input string is valid.

    An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

    ## Example 1:
    Input: s = "()"
    Output: true

    ## Example 2:
    Input: s = "()[]{}"
    Output: true

    ## Example 3:
    Input: s = "(]"
    Output: false
</details>

## longest_palindromic_substring:
<details> 
    <summary>Longest Palindromic Substring.</summary>

    Given a string s, return the longest palindromic substring in s.

    ## Example 1:
    Input: s = "babad"
    Output: "bab"
    Explanation: "aba" is also a valid answer.

    ## Example 2:
    Input: s = "cbbd"
    Output: "bb"
</details>

## regular_expression_matching:
<details> 
    <summary>Regular Expression Matching.</summary>

    Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

    '.' Matches any single character.​​​​
    '*' Matches zero or more of the preceding element.
    The matching should cover the entire input string (not partial).

    ## Example 1:
    Input: s = "aa", p = "a*"
    Output: true
    Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

    ## Example 2:
    Input: s = "ab", p = ".*"
    Output: true
    Explanation: ".*" means "zero or more (*) of any character (.)".
</details>
