// Cree una funci ́on joinWords que una varios parametros de tipo string.

function joinWords(string1) {
    return (string2) => !string2 ? string1 : joinWords(`${string1} ${string2}`);
}

result = joinWords('Hello')();
console.log(result); // Hello

result = joinWords('There')('is')('no')('spoon.')();
console.log(result); // There is no spoon.