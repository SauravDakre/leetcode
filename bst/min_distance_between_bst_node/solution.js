/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */

// inorder traversal
var minDiffInBSTInOrder = function(root) {
    let prev = Infinity
    let min = Infinity
    const walk = (root)=>{
        if(root){
            if(root.left) walk(root.left)
            min = Math.min(min, Math.abs(root.val - prev))
            prev = root.val
            if(root.right) walk(root.right)
        }
    }
    walk(root)
    return min
};


// 		b
//	a		c
// a < b < c => min must be one of (a-b) or (b-c)
// i,e one of ( maxLeft - root ) or ( root - minRight)
var minDiffInBST = function(root) {
    let result = Infinity
    
    const walk = (root)=>{
        if(!root){
            return [Infinity, -Infinity]
        }
        [minLeft, maxLeft] = walk(root.left)
        [minRight, maxRight] = walk(root.Right)
        result = Math.min(
            result, 
            Math.abs(maxLeft - root.val), 
            Math.abs(minRight - root.val)
        )
        
        return [
            Math.min(minLeft, root.val), Math.max(maxRight, root.val)
        ]
    }
    return result
};

