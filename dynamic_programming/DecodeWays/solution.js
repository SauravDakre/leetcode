/**
 * @param {string} s
 * @return {number}
 */
// Dynamic Programming solution
 var numDecodings = function(s) {
    const dp = new Array(s.length+1)
    
    for(let i=2;i<=s.length;i++){
        dp[i] = 0
    }
    
    dp[0] = 1
    dp[1] = ((s.charAt(0)==='0') ? 0 : 1)
    let i=2
    for(;i<=s.length;i++){
        const oneDigit = parseInt(s.substring(i-1,i))
        const twoDigit = parseInt(s.substring(i-2,i))
        // console.log(i,oneDigit, twoDigit)
        if (oneDigit >= 1 ){
            dp[i] += dp[i-1]
        }
        if (twoDigit >= 10 && twoDigit <= 26 ){
            dp[i] += dp[i-2]
        }
    // console.log(dp)
    }
    // console.log(i, dp)
    return dp[s.length]
};