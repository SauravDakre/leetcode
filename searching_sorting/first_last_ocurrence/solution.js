class Solution {
    find(arr, n, x) {
        //code here
        let f = this.first(arr, x, 0, arr.length - 1)
        let l = this.last(arr, x, 0, arr.length - 1)
        //   console.log('--l', l)
        return [f, l]

    }
    last(arr, x, l, r) {
        while (l <= r) {
            let mid = parseInt((l + r) / 2)
            // console.log(l, mid, r, x)
            if ((arr[mid] == x) && (mid == arr.length || arr[mid] != arr[mid + 1])) {
                return mid
            } else if (arr[mid] > x) {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
        return -1

    }
    first(arr, x, l, r) {
        while (l <= r) {
            let mid = parseInt((l + r) / 2)
            // console.log(l, mid, r, x)
            if ((arr[mid] == x) && (mid == 0 || arr[mid] != arr[mid - 1])) {
                return mid
            } else if (arr[mid] < x) {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        return -1

    }
}