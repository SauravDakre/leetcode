/**
 * Given an integer x, return true if x is palindrome integer.
 * An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
    return x === reverseNo(x);
};

function reverseNo(x) {
    let nn = 0;
    while (x > 0) {
        nn = nn * 10 + x % 10;
        x = parseInt(x / 10);
    }
    return nn;
}