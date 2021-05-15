function generateCombination(list){
    if(!list.length) return [[]]
    const firstElement = list[0]
    rest = list.slice(1)
    const withoutFirstEl = generateCombination(rest)
    const withFirstEl = []
    withoutFirstEl.forEach(el =>{
        withFirstEl.push([...el, firstElement])
    }) 
    return [...withFirstEl, ...withoutFirstEl]
}

console.log(JSON.stringify(generateCombination([1,2,3]), null,2))

// source: https://www.youtube.com/watch?v=NA2Oj9xqaZQ