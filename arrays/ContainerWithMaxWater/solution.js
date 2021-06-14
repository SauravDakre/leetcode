/**
 * @param {number[]} height
 * @return {number}
 */
 var maxArea = function(height) {
    let i=0
    let j = height.length -1 
    let sum = -Infinity
    while(i<j){
        let min = Math.min(height[i], height[j])
        if(sum < (min * (j-i))){
            sum = min * (j-i)
        }
        if(height[i] < height[j]){
            i++
        }else{
            j--
        }
    }
    return sum
};