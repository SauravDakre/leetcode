/**
 * @param {number[]} nums
 * @return {number}
 */
 var deleteAndEarn = function(nums) {
    const m = {}
    let max = 0
    for(const n of nums){
        if(m[n]){
            m[n]++
        }else{
            m[n]=1
        }
        
        if(max<n){
            max = n
        }
    }
    // console.log(m)
    const dp = new Array(max+1).fill(0)
    
    for(let i=1;i<=max;i++){
        dp[i] = Math.max(((dp[i-2]||0)+ i*(m[i]||0)), (dp[i-1]||0))
    }
    // console.log(dp)
    return dp[dp.length-1]
};