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
 * @return {number[]}
 */
var rightSideView = function (root) {
    if (!root) {
        return []
    }
    return iterate(root)
};

function iterate(root) {
    const q = []
    const res = []
    q.push(root)
    while (q.length != 0) {
        // console.log('q--',q.length)
        // console.log('r--',res)
        let c = q.length
        let tmp
        while (c--) {
            tmp = q.shift()
            // console.log(tmp, q.length)
            if (tmp.left) {
                q.push(tmp.left)
            }
            if (tmp.right) {
                q.push(tmp.right)
            }
        }
        res.push(tmp.val)
    }
    return res
}