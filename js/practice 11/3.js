// Con los datos del ejercicio anterior, muestre la cantidad total de tiempo que tomar ́an todas las
// tareas. Implemente una versi ́on iterativa y una funcional (puede usar map, filter o reduce).

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
    let totalTime = 0;
    for(var i = 0; i < data.length; i++) if(data[i].priority == 1) totalTime += data[i].duration; 
    console.log("-- Iterative");
    console.log("total time: ", totalTime);
}

Iterative(tasks);

function Functional(data) {
    console.log("-- Functional");
    console.log("total time ", data.filter(x => x.priority == 1).reduce(( sum , item , i, array ) => sum + item.duration , 0));
}

Functional(tasks);