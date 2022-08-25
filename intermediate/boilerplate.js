if (document.readyState !== 'loading') {
    runBrainfuck();
} else {
    document.addEventListener('DOMContentLoaded',function(){runBrainfuck()});
}
function runBrainfuck() {
    let t = new Array(30000).fill(0);
    let p = 0;
    let a = "";
    let v = document.createElement('textarea');
    v.setAttribute('id', 'brainfuck-out');
    v.setAttribute('style', 'height:100%%;width:100%%;');
    let body = document.getElementsByTagName("body")[0];
    body.replaceChildren(v);
    body.setAttribute('style', 'padding:0;margin:0;height:100vh;width:100vw;');
    let o = document.getElementById("brainfuck-out");
}