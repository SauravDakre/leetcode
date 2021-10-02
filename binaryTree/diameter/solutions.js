/*
class Node{
    constructor(data){
        this.data = data;
        this.left = null;
        this.right = null;
    }
}
*/

class Solution {
    // Function to return the diameter of a Binary Tree.
    diameter(root) {
        if(!root){
            return 0
        }
        // your code here
        let l = this.height(root.left)
        let r = this.height(root.right)
        
        const dl = this.diameter(root.left)
        const dr = this.diameter(root.right)
        // console.log(dl,dr, l+r+1)
        return Math.max(dl, dr, (l+r+1))
    }
    
    height(node)
    {
        if(!node){
            return 0
        }
        const l = this.height(node.left)
        const r = this.height(node.right)
        return Math.max(l,r)+1
    }
}