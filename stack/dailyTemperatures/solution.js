/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function (t) {
    const stack = []
    const res = new Array(t.length)
    for (let i = 0; i < t.length; i++) {
        // console.log(stack, res)
        if (stack.length == 0) {
            stack.push(i)
            continue
        }

        while (t[i] > t[stack[stack.length - 1]]) {
            const j = stack.pop()
            res[j] = i - j
            continue
        }

        stack.push(i)
    }
    while (stack.length !== 0) {
        res[stack.pop()] = 0
    }
    return res
};