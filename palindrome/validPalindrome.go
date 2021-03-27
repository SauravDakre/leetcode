/*
Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

ex:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
*/

func isPalindrome(s string) bool {
    ns := ""
    for _,t:=range s{
        if t>='A' && t <= 'Z' || t >='a' && t <= 'z' || t >= '0' && t <= '9'{
            ns += string(t)
        }
    }

	ns = strings.ToLower(ns)

	i,j:=0,len(ns)-1
    for ; i<j; i,j=i+1,j-1{
        if ns[i] != ns[j] {
            return false
        }
    }
    return true
}
