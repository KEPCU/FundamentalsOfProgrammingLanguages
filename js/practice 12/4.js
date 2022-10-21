// Con los datos del ejercicio anterior, muestre la frecuencia de usuarios por cada grupo. Por ejemplo
// la salida deber ́ıa ser: {editor: 2, admin: 3}. Utilice map, filter o reduce.

const users = [
    { id: 11 , name : 'Adam', age : 23 , group : 'editor' },
    { id: 47 , name : 'John', age : 28 , group : 'admin' },
    { id: 85 , name : 'William', age : 34 , group : 'editor' },
    { id: 97 , name : 'Oliver', age : 28 , group : 'admin' } ,
    { id: 100 , name : 'Vicente', age : 30 , group : 'admin' }
];

function GetOutput(current, output) {
    current.group == 'editor' ? output.editor++ : output.admin++;
    return output;
}

console.log(users.reduce(( prev , item , i, array) => GetOutput(item, output) , 0, output = { editor: 0, admin: 0 }));