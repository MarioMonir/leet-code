function lengthOfLastWord(s: string): number {
  let count: number = 0;
  let isFirstWordVisited: boolean = false;

  for (let i = s.length - 1; i >= 0; i--) {
    // We meet the last word and we return the count
    if (s[i] == " " && isFirstWordVisited) {
      return count;
    }

    // We meet the first space and we continue
    if (s[i] == " " && !isFirstWordVisited) {
      continue;
    }

    // we meet the first word and we set the flag to true
    if (!isFirstWordVisited) {
      isFirstWordVisited = true;
    }

    count++;
  }

  return count;
}

console.log(lengthOfLastWord("Hello World"));
console.log(lengthOfLastWord("   fly me   to   the moon  "));
console.log(lengthOfLastWord("luffy is still joyboy"));
