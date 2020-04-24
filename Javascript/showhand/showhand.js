/**
 * Promise 实现
 * @param {*} executor 
 */

function Promise (executor) {
    let _this = this;
        _this.status = 'pending',
        _this.value = undefined,
        _this.reason = undefined,
        _this.onResolvedCallbacks = [],
        _this.onRejectedCallbacks = [];

    function resolve (value) {
        if (_this.status === 'pending') {
            _this.status = 'resolved';
            _this.value = value;

            _this.onResolvedCallbacks.forEach(function(fn) {
                fn(value);
            });
        }
    }

    function reject (reason) {
        if (_this.status === 'pending') {
            _this.status = 'rejected';
            _this.value = reason;

            _this.onRejectedCallbacks.forEach(function(fn) {
                fn(reason);
            });
        }
    }

    try {
        executor(resolve, reject);
    } catch (error) {
        reject(error);
    }
}

Promise.prototype.then = function(onFulfilled, onRejected) {
    let _ = this;
    if (_.status === 'pending') {
        return new Promise(function(resolve, reject) {
            _.onResolvedCallbacks.push(function(value) {
                try {
                    let x = onFulfilled(value);
                    if (x instanceof Promise) {
                        x.then(resolve, reject);
                    } else {
                        resolve(x)
                    }
                } catch (error) {
                    reject(error)
                }
            });
            _.onResolvedCallbacks.push(function(value) {
                try {
                    let x = onRejected(value);
                    if (x instanceof Promise) {
                        x.then(resolve, reject);
                    } else {
                        resolve(x);
                    }
                } catch (error) {
                    reject(error)
                }
            });
        });
    }

    if (_.status === 'resolved') {
        return new Promise(function(resolve, reject) {
            try {
                let x = onFulfilled(_.value);
                if (x instanceof Promise) {
                    x.then(resolve, reject);
                } else {
                    resolve(x);
                }
            } catch (error) {
                reject(error);
            }
        });
    }

    if (_.status === 'rejected') {
        return new Promise(function(resolve, reject) {
            try {
                let x = onRejected(_.reason);
                if (x instanceof Promise) {
                    x.then(resolve, reject);
                } else {
                    resolve(x);
                }
            } catch (error) {
                reject(error);
            }
        });
    }
}

Promise.prototype.resolve = function(value) {
    return new Promise(function(resolve) {
        resolve(value);
    });
}

Promise.prototype.reject = function(value) {
    return new Promise(function(resolve, reject) {
        reject(value);
    });
}

Promise.prototype.race = function(promises) {
    return new Promise(function(resolve, reject) {
        for (let i = 0; i < promises.length; i++) {
            promises[i].then(resolve, reject);
        }
    });
}

Promise.prototype.all = function(promises) {
    return new Promise(function(resolve, reject) {
        let arr = [];

        function processData(index, y) {
            arr[index] = y;
            if (++index === promises.length) {
                resolve(arr);
            }
        }

        for (let i = 0; i < promises.length; i++) {
            promises[i].then(function(y) {
                processData(i, y)
            },reject);
        }
    });
}

/**
 * 防抖
 * @param {*} fn 
 * @param {*} time 
 */
function debounce(fn, time) {
    let timer = null;

    return function() {
        if (timer) {
            clearTimeout(timer);
        }

        setTimeout(() => {
            fn.apply(this, arguments)
        }, time);
    }
}

/**
 * 节流
 * @param {*} fn 
 * @param {*} time 
 */
function throttle(fn, time) {
    let canRun = true;

    return function() {
        if (!canRun) {
            return;
        }

        canRun = false;
        setTimeout(() => {
            fn.apply(this, arguments);
            canRun = true;
        }, time);
    }
}

/**
 * 深拷贝 数组与对象
 * @param {*} obj 
 */
function deepClone(obj) {
    var result = Array.isArray(obj) ? [] : {};
    for(var key in obj) {
        if (obj.hasOwnProperty(key)) {
            if (typeof key === 'object' && obj[key] !== null) {
                result[key] = deepClone(obj[key]);
            } else {
                result[key] = obj[key];
            }
        }
    }

    return result;
}

/**
 * 深拷贝 数组
 * @param {*} arr 
 */
function deepClone(arr) {
    return JSON.parse(JSON.stringify(arr));
}

/**
 * 数组乱序，洗牌算法
 * @param {*} arr 
 */
function shuffle(arr) {
    let len = arr.length;

    while(len--) {
        let idx = parseInt(Math.random * len--);
        [arr[idx], arr[len]] = [arr[len], arr[idx]];
    }

    return arr;
}

/**
 * 数组去重
 * @param {*} arr 
 */
function removeDup(arr) {
    return [...new Set(arr)];
}

/**
 * 数组拍平
 * @param {*} arr 
 */
function flat(arr) {
    let result = [];

    for(let i = 0; i < arr.length; i++) {
        if (Array.isArray(arr[i])) {
            result = result.concat(flat(arr[i]));
        } else {
            result.push(arr[i]);
        }
    }

    return result;
}

function flat(arr, deep) {
    let result = [];

    for(let i = 0; i < arr.length; i++) {
        if (Array.isArray(arr[i]) && deep >= 1) {
            result = result.concat(flat(arr[i], deep - 1));
        } else {
            result.push(arr[i]);
        }
    }

    return result;
}

/**
 * 实现数组 filter
 */
Array.prototype.filter = function(fn, context) {
    if (typeof fn != 'function') {
        throw new TypeError(`${fn} is not a function`);
    }

    let arr = this;
    let result = [];

    for(let i = 0; i < arr.length; i++) {
        let temp = fn.call(context, arr[i], i, arr);

        if (temp) {
            result.push(arr[i]);
        }
    }

    return result;
}

/**
 * call 函数实现
 */
Function.prototype.myCall = function(context) {
    let context = context || window;
    if (typeof this != 'function') {
        throw new TypeError('this is not a function');
    }

    context.fn = this;
    let arr = [];
    for(let i = 1; i < arguments.length; i++) {
        arr.push('argument[' + i + ']');
    }

    let result = eval('context.fn(' + arr + ')');

    delete context.fn;

    return result;
}

/**
 * apply 函数实现
 */
Function.prototype.myApply = function(context, arr) {
    if (typeof this != 'function') {
        throw new TypeError('this is not a function');
    }

    context.fn = this;
    let result = [];
    if (!arr) {
        result = context.fn();
    } else {
        let args = [];
        for(let i = 1; i < arr.length; i++) {
            args.push('arr[' + i + ']');
        }

        result = eval('context.fn(' + args + ')');
    }

    delete context.fn;
    return result;
}

/**
 * bind 函数实现
 * 1. 返回函数，内部用 apply ，改变 this 指向
 * 2. 需要将 bind 函数传的参数与返回函数参数合并
 * 3. 判断当返回函数被作为构造函数使用时，this 指向自身，否则指向 context
 */
Function.prototype.myBind = function(context) {
    if (typeof this != 'function') {
        throw new TypeError('this is not a function');
    }
    let self = this;
    let args = Array.prototype.slice(arguments, 1);

    let F = function() {};

    F.prototype = this.prototype;
    let bound = function() {
        var bindArgs = Array.prototype.slice(arguments)
        return self.apply(this instanceof F ? this : context, args.concat(bindArgs));
    }
    bound.prototype = new F();

    return bound;
}

/**
 * new 模拟实现
 */
function objFactory(fn, ...arg) {
    // let obj = new Object();

    // obj.__proto__ = Constructor.prototype;

    let obj = Object.create(fn.prototype);

    const ret = fn.apply(obj, arg)

    return typeof ret === 'object' ? ret : obj;
}

/**
 * 实现 EventEmitter
 */
class EventEmitter {
    constructor() {
        this.events = {};
    }

    on(name, cb) {
        if (!this.events[name]) {
            this.events[name] = cb;
        } else {
            this.events[name].push(cb);
        }
    }

    emit(name, ...arg) {
        if (this.events[name]) {
            this.events[name].forEach(cb => {
                cb.call(this, ...arg);
            });
        }
    }

    off(name, cb) {
        if(this.events[name]) {
            this.events[name] = this.events[name].filter(fn => {
                return fn != cb
            });
        }
    }

    once(name, cb) {
        let onlyOnce = () => {
            cb.apply(this, arguments);
            this.off(name, onlyOnce);
        }

        this.on(name, onlyOnce);
        return this;
    }
}

/**
 * jsonp 实现
 * @param {*} obj 
 */
function jsonp(obj) {
    const {url, data} = obj;
    if (!url) return;
    return new Promise((resolve, reject) => {
        const cbFn = `jsonp_${Date.now()}`;
        data.callback = cbFn;
        const head = document.querySelector('head');
        const script = document.createElement('script');
        const src = `${url}?${data2Url(data)}`;
        script.src = src;
        head.appendChild(script);

        window[cbFn] = function(res) {
            res ? resolve(res) : reject('error');
            head.removeChild(script);
            window[cbFn] = null;
        }
    });
}

function data2Url(data) {
    return Object.keys(data).reduce((acc, cur) => {
        acc.push(`${cur}=${data[cur]}`);
        return acc;
    }, []).join('&');
}

/**
 * 实现柯里化函数
 * @param {*} fn 
 * @param {*} length 
 */
function currying(fn, length) {
    const len = length || fn.length;

    return function(...args) {
        if (args.length >= len) {
            return fn.apply(this, ...args);
        } else {
            return currying(fn.bind(this, ...args), len - args.length);
        }
    }
}