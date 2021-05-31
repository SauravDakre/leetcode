/**
 * @param {number[]} nums
 * @return {number[][]}
 */
 var permute = function(nums, set=[], answers=[]) {
     console.log(nums, set, answers)
    if(!nums.length) answers.push([...set])
    for(let i=0;i<nums.length;i++){
        set.push(nums[i])
        const newN = nums.filter(n=>n!==nums[i])
        permute(newN,set, answers)
        set.pop()
    }
    return answers
};

console.log(permute([1,2,3]))