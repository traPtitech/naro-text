#[derive(Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

fn distance(a: Point, b: Point) -> f64 {
    let dx = (a.x - b.x) as f64;
    let dy = (a.y - b.y) as f64;
    (dx * dx + dy * dy).sqrt()
}

fn main() {
    let points = [
        Point { x: 4, y: 5 },
        Point { x: 7, y: 1 },
        Point { x: 2, y: 9 },
    ];
    let mut min_distance = f64::INFINITY;
    let origin = Point { x: 0, y: 0 };
    // 入力した点のうち最も原点に近い点を探し、その距離を求める
    for &p in &points {
        if min_distance > distance(origin, p) {
            min_distance = distance(origin, p);
        }
    }
    println!("{}", min_distance);
}
