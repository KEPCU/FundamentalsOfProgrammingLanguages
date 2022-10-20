// Muestre solo el nombre (name) del conjunto de datos mostrado linea abajo. Implemente una
// version iterativa y una funcional (puede usar map, filter o reduce).
let tasks = [
    {
        name : "Buy milk from the shop",
        duration : 20 ,
        priority  : 1
    },
    {
        name  : "Clean the house",
        duration : 120 ,
        priority : 3
    },
    {
        name : "Study JS functions",
        duration  : 180 ,
        priority  : 1
    }
];

function Iterative(data) {
    console.log("-- Iterative");
    for(var i = 0; i < data.length; i++) console.log(data[i].name);
}

Iterative(tasks);

function Funtional(data) {
    console.log("-- Funcional");
    data.map(x => console.log(x.name));
}

Funtional(tasks);