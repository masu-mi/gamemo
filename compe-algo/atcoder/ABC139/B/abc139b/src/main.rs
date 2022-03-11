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
    return (a+b-3)/(a-1)
}
