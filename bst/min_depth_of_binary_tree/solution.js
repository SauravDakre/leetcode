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
// Using Depth first search
var minDepthDFS = function(root) {
    if(!root){
        return 0
    }
    if(!root.left && !root.right){
        return 1
    }
    if(!root.left){
        return 1 + minDepth(root.right)
    }
    if(!root.right){
        return 1 + minDepth(root.left)
    }
    return 1 + Math.min(minDepth(root.left), minDepth(root.right))
    
};

// Breadth first search
var minDepthBFS = function(root) {
    if(!root){
        return 0
    }
    const q = []
    q.push(root)
    let minDepth = 1
    while(q.length!=0){
        len = q.length
        while(len-- > 0){
            let t = q.shift()
            if(!t.left && !t.right){
                return minDepth
            }
            if(t.left){
                q.push(t.left)    
            }
            if(t.right){
                q.push(t.right)    
            }    
        }
        minDepth++
    }
};