function isPalindrome(s: string): boolean {
  let formattedString = s.replace(/[^a-z0-9]/gi, "").toLowerCase();
  let start = 0;
  let end = formattedString.length - 1;

  while (start < end) {
    if (formattedString[start] !== formattedString[end]) return false;
    start++;
    end--;
  }

  return true;
}

console.log(isPalindrome("A man, a plan, a canal: Panama")); // true
console.log(isPalindrome("race a car")); // false
console.log(isPalindrome(" ")); // true
