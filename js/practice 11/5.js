// Muestre el score total (pilotingScore + shootingScore) de los usuarios Force (isForceUser: true).
// Implemente una versi ́on iterativa y una funcional (puede usar map, filter o reduce).

var personnel = [
    {
        id: 5 ,
        name : "Luke Skywalker",
        pilotingScore : 98 ,
        shootingScore : 56 ,
        isForceUser : true ,
    },
    {
        id: 82 ,
        name : "Sabine Wren",
        pilotingScore : 73 ,
        shootingScore : 99 ,
        isForceUser : false ,
    },
    {
        id: 22 ,
        name : "Zeb Orellios",
        pilotingScore : 20 ,
        shootingScore : 59 ,
        isForceUser : false ,
    },
    {
        id: 15 ,
        name : "Ezra Bridger",
        pilotingScore : 43 ,
        shootingScore : 67 ,
        isForceUser : true ,
    },
    {
        id: 11 ,
        name : "Caleb Dume",
        pilotingScore : 71 ,
        shootingScore : 85 ,
        isForceUser : true ,
    }
];

function Iterative(data) {
    console.log("-- Iterative");
    for(var i = 0; i < data.length; i++) if(data[i].isForceUser) console.log(data[i].pilotingScore + data[i].shootingScore);

    let totalPrice = 0;
    for(var i = 0; i < data.length; i++) if(data[i].isForceUser) totalPrice += data[i].pilotingScore + data[i].shootingScore;
    console.log("total: ", totalPrice);
}

Iterative(personnel);

function Functional(data) {
    console.log("-- Functional");
    data.filter(x => x.isForceUser).map(x => console.log(x.pilotingScore + x.shootingScore));

    console.log("total: ", data.filter(x => x.isForceUser).reduce(( sum , item , i, array ) => sum + item.pilotingScore + item.shootingScore , 0));
}

Functional(personnel);