
pub fn gcd(x: i32, y: i32) -> i32 {
    let (mut a, mut b) = (x, y);
    let mut t: i32;
    while b > 0 {
        t = a/b;
        let (na, nb) = (b, a-t*b);
        a = na;
        b = nb;
    }
    return a
}

pub fn lcm(x: i32, y: i32) -> i32 {
    return x*y/gcd(x, y)
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn gcd_works() {
        assert_eq!(gcd(2, 3), 1);
        assert_eq!(gcd(6, 4), 2);
        assert_eq!(gcd(6, 8), 2);
        assert_eq!(gcd(3, 6), 3);
        assert_eq!(gcd(3, 2), 1);
        assert_eq!(gcd(4, 6), 2);
        assert_eq!(gcd(8, 6), 2);
        assert_eq!(gcd(6, 3), 3);
    }
    #[test]
    fn lcm_works() {
        assert_eq!(lcm(2, 3), 6);
        assert_eq!(lcm(6, 4), 12);
        assert_eq!(lcm(6, 8), 24);
        assert_eq!(lcm(3, 6), 6);
        assert_eq!(lcm(3, 2), 6);
        assert_eq!(lcm(4, 6), 12);
        assert_eq!(lcm(8, 6), 24);
        assert_eq!(lcm(6, 3), 6);
    }
}
