const Arithmetics = {
	add: (a, b) => {
		return `${a} + ${b} = ${a + b}`;
	},
	subtract: (a, b) => {
		return `${a} - ${b} = ${a - b}`
	},
	multiply: (a, b) => {
		return `${a} * ${b} = ${a * b}`
	},
	division: (a, b) => {
		if (b != 0) return `${a} / ${b} = ${a / b}`;
		return `Cannot Divide by Zero!!!`;
	}
}

function sayHello() {
	return "Hello, ";
}
function greeting(helloMessage, name) {
	console.log(helloMessage() + name);
}
// Pass `sayHello` as an argument to `greeting` function
greeting(sayHello, "JavaScript!");
// Hello, JavaScript!