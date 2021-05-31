const val=[60,100,120]
const wt = [10,20,30]
const w = 50
const n = 3
const t = []
// we will create matrix of length w+1, n+1
for(let i=0;i<w+1;i++){
    const q = []
    for(let j=0;j<n+1;j++){
        q.push(-1)
    }
    t.push(q)
}
function knapsackr(wt, val, w, n){
    if(w===0 || n === 0){
        return 0
    }
    if(t[w][n] !== -1){
        return t[w][n]
    }
    if(wt[n-1] <= w){
        t[w][n] = Math.max(
            val[n-1] + knapsackr(wt, val, w-wt[n-1], n-1),
            knapsackr(wt, val, w, n-1)    
        )
        return t[w][n]
    }else{
        t[w][n] = knapsackr(wt, val, w, n-1)
        return t[w][n]
    }
}

console.log(knapsackr(wt, val, w, n))