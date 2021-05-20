func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true
    wm := make(map[string]bool)
    maxlen := 0
    for _,w:=range wordDict{
        wm[w] = true
        if maxlen < len(w) {
            maxlen = len(w)
        }
    }
    for i:=0;i<=len(s);i++{
        for j:=0; j<i;j++{
            
            if(i-j > maxlen || !dp[j]){
                    continue
                }
            if _,ok:=wm[s[j:i]];ok{
                
                dp[i] = true
            }
            // fmt.Println(dp, s[j:i], i, j)
        }
        
    }
    fmt.Println(dp)
    return dp[len(dp)-1]
}