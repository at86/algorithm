function fabnacci(place) {
    place = parseInt(place);
    if (place <= 0) {
        throw 'must be a int bigger than 0';
    }
    else if (place <= 2) {
        return 1;
    } else {
        let pre = 1, ppre = 1;
        let sum = 0;
        for (var i = 2; i < place; i++) {
            sum = ppre + pre;
            ppre = pre;
            pre = sum;
        }
        return sum;
    }
}

// fabnacci(0);

// let arr = [];
// for (let i = 1; i < 20; i++) {
//     let r = fabnacci(i);
//     arr.push(r);
// }
// console.log(arr.toString())

console.time('not recurse');
console.log(fabnacci(30));
console.timeEnd('not recurse');


function fibonacci(n) {
    if (n < 2) {
        return n;
    }
    return fibonacci(n - 2) + fibonacci(n - 1);
}
console.time('recurse');
console.log(fibonacci(30));
console.timeEnd('recurse');
