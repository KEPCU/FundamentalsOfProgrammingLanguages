function Factorial(number, memory = []) {
    if (number < 2) return 1;
    if (number in memory) return memory[number];
    return memory[number] = number * Factorial(number-1);
}

console.log(Factorial(5));

// 2.7182818284590455  2.7182818284590455
var Euler = (number) => number == 0 ? 1 : Euler(number - 1) + 1/Factorial(number);

console.log("Euler(20):", Euler(20));
