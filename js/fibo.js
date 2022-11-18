function MemoryFibonacci(number, memory = []) {
    if (number < 2) return number;
    if (number in memory) return memory[number];
    return memory[number] = MemoryFibonacci(number-1) + MemoryFibonacci(number-2);
}


function Euler(number, memory = []) {
    if (number < 1) return 1;
    if (number in memory) return memory[number];
    console.log(1/memory[number] - 1/memory[number-1]);
    return 1/Factorial(number) - 1/Factorial(number-1) < .0001 ? 1/Factorial(number) :  memory[number] = 1/Factorial(number-1) + Euler(number-1);
}

function Factorial(number, memory = []) {
    if (number < 2) return 1;
    if (number in memory) return memory[number];
    return memory[number] = number * Factorial(number-1);
}


function Factorial2(number = 0, memory = [], e = 0) {
    if (number < 2) return 1;
    if (number in memory) return memory[number];
    memory[number] = number * Factorial2(number-1);
    e += memory[number];
    return 1/Factorial(number) - 1/Factorial(number-1) < .0001 ? e : Factorial2(number - 1);
}




let cacheFact = [1];
function factMemoize(key) {
    if (!cacheFact[key]) {
      cacheFact[key] = key * factMemoize(key - 1)
    } else { // just to demo cache:
        //console.log("cache hit:", key)
    }
    return cacheFact[key]
}
/*
console.log(factMemoize(5));

// e = 1/0! + 1/1! + 1/2! ...
// eR(n) = er(n-1) + 1/n!
// eR(0) = 1 
function eR(n){
    console.log(1/factMemoize(n) - 1/factMemoize(n-1));
  if (n == 0)
    return 1;
  else
    return eR(n-1) + 1/factMemoize(n);
}
console.log("eR(20)", eR(20));

//console.log(Euler(4));
console.log(eR(20));
*/
let res = "c"
let empty 
function cafe (ingre) {
    if(ingre = 'le') {
        let res = 'otra'
        return res
    }
}

var tasa = cafe('le')
console.log(tasa)
console.log(empty)