// -*- coding:utf-8-unix -*-

use proconio::input;

// ABC086C - Traveling
// https://atcoder.jp/contests/abs/tasks/arc089_a

fn main() {
    input! {
        n: usize,
        mut xs: [i32; n],  // Vec<(i32, i32, i32)>
    }
    let mut sum = 0;
    for x in &xs {
        sum += x;
    }
    let mut min = std::i32::MAX;
    let p = sum / n as i32;
    for i in 0..=1 {
        let cur = score(p + i, &xs);
        if cur < min {
            min = cur
        }
    }
    println!("{}", min);
}
fn score(p: i32, xs: &Vec<i32>) -> i32 {
    let mut total: i32 = 0;
    for x in xs {
        total += (x - p) * (x - p);
    }
    return total;
}
