// Muestre la edad promedio de las especies tipo “dog”. Utilice map, filter o reduce.

arr = [
    { name :'luna', sex :'f', age :7 , species :'dog'},
    { name :'jimmy', sex :'m', age :122 , species :'human'},
    { name :'snoop', sex :'m', age :60 , species :'human'},
    { name :'jennifer', sex :'f', age :250 , species :'human'},
    { name :'yeller', sex:'20', age :20 , species :'dog'},
];

console.log("average: ", arr.filter(x => x.species == 'dog').reduce((prev,curr,i,array) => prev + (curr.age)/array.length,0));