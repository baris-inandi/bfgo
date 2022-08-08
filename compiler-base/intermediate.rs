use std::collections::HashMap;

struct Bf {
    tape: HashMap<u16, u8>,
    pointer: u16,
}

#[allow(dead_code)]
impl Bf {
    pub fn new() -> Bf {
        Bf {
            tape: HashMap::new(),
            pointer: 0,
        }
    }
    pub fn c(&self) -> u8 {
        *self.tape.get(&self.pointer).unwrap_or(&0)
    }
    pub fn l(&mut self, n: u16) {
        self.pointer -= n;
    }
    pub fn r(&mut self, n: u16) {
        self.pointer += n;
    }
    pub fn p(&mut self, n: u8) {
        self.tape.insert(
            self.pointer,
            self.tape.get(&self.pointer).unwrap_or(&0).wrapping_add(n),
        );
    }
    pub fn m(&mut self, n: u8) {
        self.tape.insert(
            self.pointer,
            self.tape.get(&self.pointer).unwrap_or(&0).wrapping_sub(n),
        );
    }
    pub fn w(&mut self) -> bool {
        self.c() != 0
    }
    pub fn o(&mut self) {
        print!("{}", self.c() as char);
    }
}

#[allow(unused_variables)]
#[allow(unused_mut)]
fn main() {
    let mut b = Bf::new();
    // brainfuck ir %s
}
