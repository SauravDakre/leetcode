/**
 * @param {number[]} prices
 * @return {number}
 */

 var maxProfit = function(prices) {
    let mp = 0
    let min = Infinity
    for(const x of prices){
        if(min > x){
            min = x
        }
        
        if(mp < (x - min)){
            mp = x - min
        }
        
    }
    return mp
}

var maxProfitUsingStack = function(prices) {
    const ar = []
    let max = 0
    for(const x of prices){
        if(ar.length === 0){
            ar.push(x)
            continue
        }

        while(ar[ar.length-1] > x){
            ar.pop()
        }
        ar.push(x)

        if(ar.length >= 2){
            if(max < (ar[ar.length-1] - ar[0])){
                max = (ar[ar.length-1] - ar[0])
            }
        }
    }

    return max
};