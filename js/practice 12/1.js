//Convierta el array (lineas abajo) de Celsius a Fahrenheit. Utilice map, filter o reduce.

const celsius = [ -15 , -5, 0, 10 , 16 , 20 , 24 , 32];

console.log(celsius.map(x => ( x * 1.8 ) + 32));