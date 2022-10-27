// Implemente una funci ́on delayInvoc que en cada invocaci ́on incremente la variable total con el
// valor enviado como parametro.

var total = 0;

var delayInvoc = function (a) {
    total += a;
    return function (b) {        
        if(b) return delayInvoc(b);
    };
};


delayInvoc(4)(5);
console.log("4,5: ",total); //9

delayInvoc(4)(5)(8);
console.log("4,5,8: ",total); // 26

delayInvoc(1)(1)(1)(1);
console.log("1,1,1,1: ",total); // 30

