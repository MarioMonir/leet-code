function longestCommonPrefix(strs: string[]): string {
  if (strs.length < 2) {
    return "";
  }

  // sort the array by length
  strs.sort((a, b) => a.length - b.length);

  let matchString = "";
  const firstWord = strs[0];

  // iterator for first word
  for (let i = 0; i < firstWord.length; i++) {
    // we will iterate in all words except the first one

    for (let j = 1; j < strs.length; j++) {
      const currentWord = strs[j];

      if (currentWord[i] !== firstWord[i]) return matchString;
    }

    matchString += firstWord[i];
  }

  return matchString;
}

console.log(longestCommonPrefix(["flower", "flow", "flight"])); // "fl"
