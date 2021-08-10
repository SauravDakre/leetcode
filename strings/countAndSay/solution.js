/**
 * @param {number} n
 * @return {string}
 */
var countAndSay = function (n) {
    if (n === 1) {
        return "1"
    }
    if (n === 2) {
        return "11"
    }

    let s = "11"
    for (let i = 3; i <= n; i++) {
        let t = ''
        s += '='
        let c = 1
        for (let j = 0; j < s.length - 1; j++) {
            if (s[j] != s[j + 1]) {
                t += `${c}${s[j]}`
                c = 1
                continue
            }
            c++
        }
        s = t

    }
    return s
};