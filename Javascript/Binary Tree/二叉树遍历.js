/**
 * DFS
 */

// 先序遍历-递归
function preOrder(node) {
    if (node) {
        console.log(node.val);
        preOrder(node.left);
        preOrder(node.right);
    }
}

// 先序遍历-非递归
function preOrder(node) {
    let stack = [node];

    while(stack.length) {
        let node = stack.pop();

        console.log(node.val);
        node.right && stack.push(node.right);
        node.left && stack.push(node.left);
    }
}



// 中序遍历-递归
function middleOrder(node) {
    if (node) {
        middleOrder(node.left);
        console.log(node.val);
        middleOrder(node.right);
    }
}

// 中序遍历-非递归
function middleOrder(node) {
    let stack = [];

    while(stack.length || node) {
        if (node) {
            stack.push(node);
            node = node.left;
        } else {
            node = stack.pop();
            console.log(node.val);
            node = node.right;
        }
    }
}

// 后序遍历-递归
function postOrder(node) {
    if (node) {
        postOrder(node.left);
        postOrder(node.right);
        console.log(node.val);
    }
}

// 后序遍历-非递归
function postOrder(node) {
    let stack = [node];
    let tmp = null;

    while(stack.length) {
        tmp = stack[stack.length - 1];
        if (tmp.left && node !== tmp.left && node !== tmp.right) {
            stack.push(tmp.left);
        } else if (tmp.right && node !== tmp.right) {
            stack.push(tmp.right);
        } else {
            console.log(stack.pop().val);
            node = tmp;
        }
    }
}

/**
 * BFS
 */

function levelOrder(node) {
    let queue = [node];

    while(queue.length) {
        const node = queue.shift();

        console.log(node.val);

        node.left && node.push(node.left);
        node.right && node.push(node.right);
    }
}