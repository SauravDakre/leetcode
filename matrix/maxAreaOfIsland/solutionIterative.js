/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxAreaOfIsland = function (grid) {
    const stack = []
    let max = 0
    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            let count = 0
            if (grid[i][j] == 1) {
                stack.push({ r: i, c: j })
                grid[i][j] = 0
                while (stack.length !== 0) {
                    const pt = stack.pop()
                    count++
                    const pts = [
                        { r: pt.r - 1, c: pt.c },
                        { r: pt.r + 1, c: pt.c },
                        { r: pt.r, c: pt.c - 1 },
                        { r: pt.r, c: pt.c + 1 }
                    ]
                    for (const t of pts) {
                        if (t.r >= 0 && t.r < grid.length && t.c >= 0 && t.c < grid[i].length && grid[t.r][t.c] === 1) {
                            stack.push(t)
                            grid[t.r][t.c] = 0
                        }
                    }
                }
                if (max < count) {
                    max = count
                }
            }
        }
    }
    return max
};