const fizzBuzz = (num: number): "FizzBuzz" | "Fizz" | "Buzz" | number => {
  if (num % 15 === 0) return "FizzBuzz";

  if (num % 5 === 0) return "Fizz";

  if (num % 3 === 0) return "Buzz";

  return num;
};

const main = () => {
  for (let i: number = 1; i < 100; i++) {
    console.log(fizzBuzz(i));
  }
};

main();
