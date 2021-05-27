/**
 * @param {string} text1
 * @param {string} text2
 * @return {number}
 */
var longestCommonSubsequence = function(text1, text2) {
    const ar = []
    for (let i = 0;i<=text1.length;i++){
        const c = []
        for(let j = 0; j<=text2.length;j++){
            c.push(0)
        }
        ar.push(c)
    }
    // console.log(ar)
    let i,j
    for (i = 1;i<ar.length;i++){
        for(j = 1; j<ar[i].length;j++){
            if(text1[i-1] === text2[j-1]){
                ar[i][j] = ar[i-1][j-1] + 1
            }else{
                ar[i][j] = Math.max(ar[i-1][j], ar[i][j-1])
            }
        }
    }
    // console.log(ar)
    return ar[ar.length-1][ar[0].length-1]
};