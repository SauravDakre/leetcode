function getMax(operations) {
    // Write your code here
    const st = []
    const sk = []
    const res = []
    // maintain 2 stacks
    // st will contain elements in increasing order
    // sk will be normal stack
    for(let i=0;i<operations.length;i++){
        const op = operations[i]
        const ar = op.split(' ')
        if(ar.length > 1){
            if(ar[0] === '1'){
                const t = parseInt(ar[1])
                if(st.length == 0 || (st.length > 0 && st[st.length-1].t < t)){
                    st.push({t,i})
                }
                sk.push({t,i})
                
            }
        }
        // since element can be duplicate, also store index of element
        // before poping from st, check whether index is same or not
        if(ar[0] === '2'){
            if((st[st.length-1].t === sk[sk.length-1].t) && (st[st.length-1].i === sk[sk.length-1].i)){
                st.pop()
            }
            sk.pop()
        }
        // result will contain elements from st stack
        if(ar[0] === '3'){
            res.push(st[st.length-1].t)
        }
        // console.log(op, st, res,sk)
    }
    return res
}