/**
 * @param {string} text1
 * @param {string} text2
 * @return {number}
 */
 var longestCommonSubstring = function(text1, text2) {
    const ar = []
    for (let i = 0;i<=text1.length;i++){
        const c = []
        for(let j = 0; j<=text2.length;j++){
            c.push(0)
        }
        ar.push(c)
    }
    // console.log(ar)
    let i,j,max=0
    for (i = 1;i<ar.length;i++){
        for(j = 1; j<ar[i].length;j++){
            if(text1[i-1] === text2[j-1]){
                ar[i][j] = ar[i-1][j-1] + 1
            }else{
                ar[i][j] = 0
            }
            if(max < ar[i][j]){
                max = ar[i][j]
            }
        }
    }
    // console.log(ar)
    return max
};

console.log(longestCommonSubstring("GeeksforGeeks", 'GeeksQuiz'))
console.log(longestCommonSubstring("zxabcdezy", 'yzabcdezx'))
console.log(longestCommonSubstring("abcdxyz", 'xyzabcd'))