function isSubsequence(s: string, t: string): boolean {
  let sp = 0;
  let tp = 0;
  if (s.length === 0) return true;
  while (tp < t.length) {
    if (s[sp] === t[tp]) sp++;
    if (sp === s.length) return true;
    tp++;
  }
  return false;
}

console.log(isSubsequence("abc", "ahbgdc")); // true
console.log(isSubsequence("axc", "ahbgdc")); // false
console.log(isSubsequence("", "")); // true
