// Con los datos del ejercicio anterior, ahora muestre las tareas con prioridad (priority) 1. Imple-
// mente una versi ́on iterativa y una funcional (puede usar map, filter o reduce).

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

    for(var i = 0; i < data.length; i++) if(data[i].priority == 1) console.log(data[i].name);
}

Iterative(tasks);

function Functional(data) {
    console.log("-- Functional");

    data.filter(x => x.priority == 1).map(x => console.log(x.name));
}

Functional(tasks);