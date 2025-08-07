### Go Basic Data Structures

# Part 1: Palindrome

This section includes three different implementations of a palindrome-checking algorithm. Assosciated packages:

- palindrome: implementaion of three algorithms, tests and benchmarks
- stack: implementation of stack data strucutre used in `CheckWithStack` algorithm

---

## 1. `CheckWithStack`

**Description:**  
This implementation uses a custom stack (based on a slice of runes). It pushes each character of the cleaned string onto the stack, then pops characters while comparing them to the original string to check if it’s a palindrome.

**Time Complexity:**

- **O(n)** — for both pushing and popping characters.

**Memory Allocation:**

- `~n` bytes for the cleaned string
- `~n` bytes for the lowercase conversion
- `~4n` bytes for the stack slice (`[]rune`)
- **Total ≈ 6n bytes**

---

## 2. `CheckWithDoublePointer`

**Description:**  
This approach converts the cleaned string into a rune slice and uses two pointers—one starting at the beginning and one at the end—moving inward and comparing characters.

**Time Complexity:**

- **O(n)** — for iterating through the string.

**Memory Allocation:**

- `~n` bytes for the cleaned string
- `~n` bytes for the lowercase conversion
- `~4n` bytes for the rune slice
- **Total ≈ 6n bytes**

---

## 3. `CheckWithReversedString`

**Description:**  
This method reverses the rune slice in-place and converts it back to a string, then compares it with the cleaned original to determine if it’s a palindrome.

**Time Complexity:**

- **O(n)** — for reversing and comparing.

**Memory Allocation:**

- `~n` bytes for the cleaned string
- `~n` bytes for the lowercase conversion
- `~4n` bytes for the rune slice
- `~n` bytes for the reversed string (`string(runes)`)
- **Total ≈ 7n bytes**

---

## Summary

| Method                    | Time Complexity | Memory Allocation |
| ------------------------- | --------------- | ----------------- |
| `CheckWithStack`          | O(n)            | ~ 6n bytes        |
| `CheckWithDoublePointer`  | O(n)            | ~ 6n bytes        |
| `CheckWithReversedString` | O(n)            | ~ 7n bytes        |

# Part 2: Creating Set

Description:
This function removes duplicates from a slice of strings by using a map[string]bool to track already seen values. It appends each unique string to the result slice.

---

# Part 3: Reversing Characters in Sentence

## 1 `ReverseCharactersOrderRaw`

**Time Complexity**:

- **O(n)** — where n is the number of runes in the input string.

**Memory**:

- n runes for `[]rune(s)`
- n runes for the result slice
- n bytes for the final string (`string(result)`)
- **Total ≈ 2 n runes + n bytes** (≈ 9n bytes)

## 2 `ReverseCharactersOrder`

**Time Complexity**

- **O(n)** - for splitting, reversing, and joining

**Memory**:

- O(w) string-headers for the `[]string` from `Split` (no byte copies)
- n runes (intermediate per-word reversal)
- n bytes for the final `Join` buffer
- **Total ≈ 5 n bytes + O(w)**

---
