fn main() {
    let mut v = Vec::new();
    for i in 1..10 {
        v.push(i);
    }
    for i in 1..8 {
        // use std
        println!(" std: v.binary_search(&{}).unwrap(): {}", i, v.binary_search(&i).unwrap());
        let r = binary_search(&v, &i).unwrap();
        println!("orig: binary_search(v, &{}).unwrap(): {}: {}", i, r.0, r.1);
    }
}

fn binary_search (vec: &Vec<i32>, i: &i32) -> Result<(usize, i32), i32> {
    let mut l: i32 = -1;
    let mut r: i32 = vec.len() as i32;
    if vec.len() == 0 {
        return Err(0);
    }
    while (r-l).abs() > 1 {
        let mid = l + ((r-l)>>1);
        if vec[mid as usize] >= *i {
            r = mid
        } else {
            l = mid
        }
    }
    if r == -1 || r == (vec.len() as i32) {
        return Err(r);
    }
    Ok((r as usize, vec[r as usize]))
}
