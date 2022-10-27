// Implemente una funci ́on que calcule el volumen de un cilindro. Incluya la versi ́on normal y una
// aplicando Currying.

var NormalCylinderVolume = (r, h) => Math.PI * Math.pow(r,2) * h;

console.log("Normal: ",NormalCylinderVolume(2, 3));

var CurryingCylinderVolume =(r) => {
    return (h) => Math.PI * Math.pow(r,2) * h;
}
console.log("Currying: ",CurryingCylinderVolume(2)(3));