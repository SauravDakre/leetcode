function knapsackr(wt, val, w, n){
    if(w===0 || n === 0){
        return 0
    }
    if(wt[n-1] <= w){
        return Math.max(
            val[n-1] + knapsackr(wt, val, w-wt[n-1], n-1),
            knapsackr(wt, val, w, n-1)    
        )
    }else{
        return knapsackr(wt, val, w, n-1)
    }
}

const val=[60,100,120]
const wt = [10,20,30]
console.log(knapsackr(wt, val, 50, 3))