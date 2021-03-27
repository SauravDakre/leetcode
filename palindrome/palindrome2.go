package palindrome
/*
680. Valid Palindrome II
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.
*/

func validPalindrome(s string) bool {
    for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{ 	// check if string is valid palindrome as per logic in isPalindrome function
        if s[i] != s[j]{	// if the char do not match then check by omitting first char and last char
            return isSubstrPalindrome(s,i,j-1) || isSubstrPalindrome(s,i+1,j)
        }
    }
    return true
}

func isSubstrPalindrome(s string, i,j int) bool {
    for ;i<j;i,j=i+1,j-1{
        if s[i] != s[j]{
            return false
        }
    }
    return true
}
