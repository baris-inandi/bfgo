package boilerplate

const JS_IR_BOILERPLATE string = `(function() {
    let t = new Array(%d).fill(0);
    let p = 0;
    let a = "";
    let v = document.createElement('textarea');
    v.setAttribute('id', 'brainfuck-out');
    v.setAttribute('spellcheck', 'false');
    v.setAttribute('style', 'height:100%%;width:100%%;font-family:monospace;');
    let body = document.getElementsByTagName("body")[0];
    body.replaceChildren(v);
    body.setAttribute('style', 'padding:0;margin:0;height:100vh;width:100vw;');
    let o = document.getElementById("brainfuck-out");%s
})()`
