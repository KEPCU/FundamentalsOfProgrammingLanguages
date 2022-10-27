// Implemente una funci ́on curry que tome como argumento cualquier funci ́on f y retorne la versi ́on
// curried de f.

function abc(a, b, c) {
    return a+b+c;
}

function curry(f) {
    return function curry2(...args) {
        if (args.length >= f.length) return f.apply(this, args);
        else return function curry3(...args2) {
            return curry2.apply(this, args.concat(args2));
        }
    };
}
  
var curriedAbc = curry(abc);

console.log("(2)(3)(4): ", curriedAbc(2)(3)(4)); // 9
console.log("(2, 3)(4): ", curriedAbc(2, 3)(4)); // 9
console.log("(2)(3, 4): ", curriedAbc(2)(3, 4)); // 9
console.log("(2, 3, 4): ", curriedAbc(2, 3, 4)) ; // 9
