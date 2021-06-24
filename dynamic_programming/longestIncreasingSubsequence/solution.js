/**
 * @param {number[]} nums
 * @return {number}
 */
 var lengthOfLIS = function(nums) {
    let dp = new Array(nums.length)
    for(let j=0;j<nums.length;j++){
        dp[j] = 1
    }
    // console.log(dp)
    let max = -Infinity
    for(let i=1;i<nums.length;i++){
        for(let j=0;j<i;j++){
            if(nums[i] > nums[j]){
                dp[i] = Math.max(dp[i], dp[j] + 1)
                if(max < dp[i]){
                    max = dp[i]
                }
            }
        }
    }
    return max
};