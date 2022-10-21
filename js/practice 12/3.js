// Con los datos del ejercicio anterior, revise si los usuarios tienen provilegios de administrador. La
// salida debe mostrar true o false por cada usuario. Utilice map, filter o reduce.

const users = [
    { id: 11 , name : 'Adam', age : 23 , group : 'editor' },
    { id: 47 , name : 'John', age : 28 , group : 'admin' },
    { id: 85 , name : 'William', age : 34 , group : 'editor' },
    { id: 97 , name : 'Oliver', age : 28 , group : 'admin' } ,
    { id: 100 , name : 'Vicente', age : 30 , group : 'admin' }
];

users.map(x => console.log(x.group == 'admin' ? "TRUE" : "FALSE" ,x.id, x.name, x.age, x.group));

console.log("-- Filter 'i'")
users.filter(x => x.name.includes("i")).map(x => console.log(x.group == 'admin' ? "TRUE" : "FALSE" ,x.id, x.name, x.age, x.group));