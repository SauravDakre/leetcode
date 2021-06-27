/**
 * @param {character[][]} board
 * @param {string} word
 * @return {boolean}
 */
 var exist = function(board, word) {
    for (let i=0;i<(board.length);i++){
        for (let j=0;j<(board[i].length);j++){
            if(board[i][j] === word[0] && search(board, word, 0, i, j)){
                return true
            }
        }
    }
    return false
};

function search(board, words, i,r,c ) {
    // console.log(board, words[i],i, r, c)
    if(i === words.length){
        return true
    }
    if(r < 0 || r >= board.length || c < 0 || c >= board[0].length || board[r][c] !== words[i]){
        return false
    }
    // console.log('rc',r,c, board.length, board[0].length)
    const tmp = board[r][c]
    board[r][c] = "-1"
    const found = (search(board, words, i+1, r+1, c) || 
        search(board, words, i+1, r-1, c) ||
        search(board, words, i+1, r, c+1) ||
        search(board, words, i+1, r, c-1))
    board[r][c] = tmp
    return found
}

