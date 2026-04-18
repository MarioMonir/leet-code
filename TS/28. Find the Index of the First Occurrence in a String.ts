function strStr(haystack: string, needle: string): number {
  let m = 0;

  for (let i = 0; i < haystack.length; i++) {
    if (haystack[i] !== needle[m]) {
      if (m > 1) i -= m - 1;
      m = 0;
    }

    if (haystack[i] === needle[m]) m++;

    if (m === needle.length) return i - needle.length + 1;
  }

  return -1;
}

console.log(strStr("sadbutsad", "sad"), strStr("sadbutsad", "sad") === 0); // 0
console.log(
  strStr("asabsadbutsad", "sad"),
  strStr("asabsadbutsad", "sad") === 4,
); // 4
console.log(strStr("leetcode", "leeto"), strStr("leetcode", "leeto") === -1); // -1
console.log(strStr("sad", "sad"), strStr("sad", "sad") === 0); // 0
console.log(strStr("a", "a"), strStr("a", "a") === 0); // 0
console.log(
  strStr("mississippi", "issip"),
  strStr("mississippi", "issip") === 4,
); // 4

console.log(strStr("ssaddbutsad", "sad"), strStr("ssaddbutsad", "sad") === 1); // 2
