// Una manera de calcular el seno y coseno es con la siguiente f ́ormula, la cantidad de operandos
// define la precisi ́on del valor de seno/coseno. Desarrolle un programa que calcule el seno/coseno
// aproximado, utilice memoization y recursividad.

function Factorial(number, memory = []) {
    if (number < 2) return 1;
    if (number in memory) return memory[number];
    return memory[number] = number * Factorial(number-1);
}

var seno = (x , p) => p < 3 ? x : Math.pow(x,p)/Factorial(p) + seno(x,p+2);

console.log(seno(30,5));