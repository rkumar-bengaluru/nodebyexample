function nsmallest(i) {
    console.log("smallest..." + JSON.stringify(i))
    var largest = Number.MIN_VALUE;
    var sl = largest;
    // 5,4,3,2,1
    for (var k = 0; k < i.length;k++) {
        if (largest < i[k]) {
            sl = largest; 
            largest = i[k]; 
        } else if (sl < i[k]) {
            sl = i[k];
        }
    }
    return sl
}

function nsmallestsorted(i,k) {
    var sorted = i.sort();
    console.log(sorted)
    return sorted[i.length - k - 1]
    
}

function printHeap(input) {
    var n = input.length;
     var rr = ""
    for (var i = 0; i <= n / 2; i++) {
        var r = ''
        if(i + 2*i + 2 <= n) {
            var c1 = 2*i+ 1;
            var c2 = 2*i+ 2;
            
            r = input[i] + `\n` + input[c1] + `   ` + input[c2] + `\n`;
        }
        rr += r;
    }
    console.log(rr)
}

function nsmallestunsorted(i,k) {
    console.log(i);
    printHeap(i);
}

export {nsmallest,nsmallestsorted,nsmallestunsorted}