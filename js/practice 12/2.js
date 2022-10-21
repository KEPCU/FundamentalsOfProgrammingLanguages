// Muestre los usuarios que contienen in “i” en su nombre (name). Utilice map, filter o reduce.

const users = [
    { id: 11 , name : 'Adam', age : 23 , group : 'editor' },
    { id: 47 , name : 'John', age : 28 , group : 'admin' },
    { id: 85 , name : 'William', age : 34 , group : 'editor' },
    { id: 97 , name : 'Oliver', age : 28 , group : 'admin' } ,
    { id: 100 , name : 'Vicente', age : 30 , group : 'admin' }
];

users.filter(x => x.name.includes("i")).map(x => console.log(x.id, x.name, x.age, x.group));