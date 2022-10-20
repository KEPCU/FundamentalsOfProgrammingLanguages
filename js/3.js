const vehicles = [
    { make: "Honda", model: 'CR-V', type: 'suv', price: 74045 },
    { make: "Honda", model: 'Accord', type: 'sedan', price: 32455 },
    { make: "Honda", model: 'Accord', type: 'sedan', price: 22455 },
    { make: "Honda", model: 'Accord', type: 'sedan', price: 52455 },
    { make: "Mazda", model: 'Mazda 6', type: 'sedan', price: 24195 },
    { make: "Mazda", model: 'CX-9', type: 'suv', price: 31520 },
    { make: "Toyota", model: '4Runner', type: 'suv', price: 34210 },
    { make: "Toyota", model: 'Sequoia', type: 'suv', price: 45560 },
    { make: "Toyota", model: 'Tacoma', type: 'truck', price: 24320 },
    { make: "Ford", model: 'F-150', type: 'truck', price: 27110 },
    { make: "Ford", model: 'Fusion', type: 'sedan', price: 22120 },
    { make: 'Ford', model: 'Explorer', type: 'suv', price: 31660 }
];

function GetMax(prev , current) {
    return current > prev ? current : prev;
}

console.log("price ", Math.max(vehicles.filter(x => x.make == "Honda").map(x => x.price)));
//document.write("price ", vehicles.filter(x => x.make == "Honda").reduce(( prev , item , i, array ) => prev + item.price , 0));
console.log("price ", vehicles.filter(x => x.make == "Honda").reduce(( prev , item , i, array) => GetMax(prev, item.price) , 0));