// -*- coding:utf-8-unix -*-

use proconio::input;

// ABC086C - Traveling
// https://atcoder.jp/contests/abs/tasks/arc089_a

fn main() {
    input! {
        a: usize,
        b: usize,
    }
    println!("{}", calc(a, b));
}

fn calc(a: usize, b: usize) -> usize {
    let mut n = 1;
    let mut c = 0;
    while n < b {
        n = n - 1 + a;
        c += 1;
    }
    return c;
}
