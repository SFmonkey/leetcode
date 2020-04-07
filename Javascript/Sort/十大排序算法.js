var arr=[3,44,38,5,47,15,36,26,27,2,46,4,19,50,48];

// 1.冒泡排序
function bubbleSort(arr) {
    var len = arr.length;

    for(var i = 0; i < len; i++) {
        for(var j = 0; j < len - 1 - i; j++) {
            if (arr[j] > arr[j + 1]) {
                var temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }

    return arr;
}

// 2.选择排序
function selectionSort(arr) {
    var len = arr.length;
    var minIndex;
    for (let i = 0; i < len; i ++) {
        minIndex = i;
        for (let j = i + 1; j < len; j ++) {
            if (arr[minIndex] > arr[j]) {
                minIndex = j;
            }
        }

        if (minIndex !== i) {
            var temp = arr[i];
            arr[i] = arr[minIndex];
            arr[minIndex] = temp;
        }
    }

    return arr;
}

// 3.插入排序
function insertionSort(arr) {
    for (let i = 1; i < arr.length; i ++) {
        let key = arr[i];
        let j =  i - 1;

        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j --;
        }

        arr[j + 1] = key;
    }

    return arr;
}

// 4.希尔排序
function shellSort(arr) {
    var len = arr.length,
        temp,
        gap = 1;

    while(gap < len / 5) {
        gap = gap * 5 + 1;
    }

    for(gap; gap > 0; gap = Math.floor(gap / 5)) {
        for (var i = gap; i < len; i++) {
            temp = arr[i];
            for (var j = i - gap; j >=0 && arr[j] > temp; j-=gap) {
                arr[j + gap] = arr[j];
            }
            arr[j + gap] = temp;
        }
    }

    return arr;
}

// 5.归并排序
function mergeSort(arr) {
    var len = arr.length;
    if (len < 2) {
        return arr;
    }

    var middle = Math.floor(len / 2),
        left = arr.slice(0, middle),
        right = arr.slice(middle);

    return merge(mergeSort(left), mergeSort(right));
}

function merge(left, right) {
    var result = [];
    while (left.length && right.length) {
        if (left[0] <= right[0]) {
            result.push(left.shift());
        } else {
            result.push(right.shift());
        }
    }

    while (left.length) {
        result.push(left.shift());
    }

    while(right.length) {
        result.push(right.shift());
    }

    return result;
}

// 6.快速排序
function quickSort(arr) {
    if (arr.length <= 1) return arr;
    let pivotIndex = Math.floor(arr.length / 2);
    let pivot = arr.splice(pivotIndex, 1)[0];
    let left = [];
    let right = [];

    for(var i = 0; i < arr.length; i++) {
        if (arr[i] < pivot) {
            left.push(arr[i]);
        } else {
            right.push(arr[i]);
        }
    }

    return [...quickSort(left), pivot, ...quickSort(right)];
}

// 7.堆排序
function heapSort(arr) {
    let heapSize = arr.length, temp;

    for(let i = Math.floor(heapSize / 2) - 1; i >=0; i--) {
        heapify(arr, i, heapSize);
    }

    for (var j = heapSize - 1; j >= 1; j--) {
        temp = arr[0];
        arr[0] = arr[j];
        arr[j] = temp;
        heapify(arr, 0, --heapSize );
    }

    return arr;
}

function heapify(arr, x, len) {
    let l = 2 * x + 1, r = 2 * x + 2, largest = x, temp;

    if (l < len && arr[l] > arr[largest]) {
        largest = l;
    }

    if (r < len && arr[r] > arr[largest]) {
        largest = r;
    }

    if (largest != x) {
        temp = arr[x];
        arr[x] = arr[largest];
        arr[largest] = temp;
        heapify(arr, largest, len);
    }
}

// 8.计数排序
function countingSort(arr, max) {
    let len = arr.length;
    let res = [];

    let temp = [];
    for (let i = 0; i <= max; i ++) {
        temp[i] = 0;
    }

    for(let j = 0; j < len; j++) {
        temp[arr[j]] ++;
    }

    for(let k = 0; k <=max; k++) {
        while(temp[k]--) {
            res.push(k);
        }
    }

    return res;
}

// 9.桶排序
function bucketSort(arr, num) {
    if (arr.length <= 1) return arr;

    let len = arr.length,
        buckets = [],
        result = [],
        min = max = arr[0],
        regex = '/^[1-9]+[0-9]$/',
        space,
        n = 0;

    num = num || ((num > 1 && regex.test(num)) ? num : 10);

    for(let i = 1; i < len; i++) {
        min = min <= arr[i] ? min : arr[i];
        max = max >= arr[i] ? max : arr[i];
    }

    space = (max - min + 1) / num;
    for(let i = 0; i < len; i++) {
        let idx = Math.floor((arr[i] - min) / space);

        if (buckets[idx]) {
            let k = buckets[idx].length - 1;
            while(k >=0 && buckets[idx][k] > arr[i]) {
                buckets[idx][k + 1] = buckets[idx][k];
                k --;
            }
            buckets[idx][k + 1] = arr[i];
        } else {
            buckets[idx] = [];
            buckets[idx].push(arr[i])
        }
    }

    while(n < num) {
        result = result.concat(buckets[n]);
        n ++;
    }

    return result;
}

// 10.基数排序
function radixSort(arr, maxDigit) {
    let mod = 10;
    let dev = 1;
    let counter = [];

    for(let i = 0; i < maxDigit; i ++, dev *= 10, mod *= 10) {
        for(let j = 0; j < arr.length; j++) {
            let bucket = parseInt((arr[j] % mod) / dev);
            if (counter[bucket] == null) {
                counter[bucket] = [];
            }

            counter[bucket].push(arr[j]);
        }

        let pos = 0;
        for(let j = 0; j < counter.length; j++) {
            let val = null;
            if (counter[j] != null) {
                while ((val = counter[j].shift()) != null) {
                    arr[pos++] = val;
                }
            }
        }
    }

    return arr;
}

console.log(radixSort(arr, 2));