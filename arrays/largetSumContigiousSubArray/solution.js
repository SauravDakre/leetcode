// Kadane's Algorithm

function maxSubarraySum(arr, N){
    // code here
    let maxSoFar = -Infinity
    let maxEndingHere = 0
    for(let i=0;i<N;i++){
        maxEndingHere += arr[i]
        if(maxSoFar < maxEndingHere){
            maxSoFar = maxEndingHere
        }
        if(maxEndingHere < 0){
            maxEndingHere = 0
        }
    }
    return maxSoFar
} 

console.log(maxSubarraySum([-1,-2,-3,-4], 4))