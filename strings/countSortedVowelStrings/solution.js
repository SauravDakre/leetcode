/**
 * @param {number} n
 * @return {number}
 */
 var countVowelStrings = function(n) {
    const ar = [1,1,1,1,1]
    for(let i=1;i<n;i++){
        for(let j=ar.length-2;j>=0;j--){
            ar[j] = ar[j]+ar[j+1]
        }
    }
    return ar.reduce((c, a)=>c+a, 0)
};

