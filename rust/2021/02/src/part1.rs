use std::fs::File;
use std::io;
use std::io::BufRead;

pub fn part1() {
    let file = io::BufReader::new(
        File::open("02.input").unwrap()
    );
    
    let mut dist: usize = 0;
    let mut depth: usize = 0;

    for line in file.lines() {
        let line = line.unwrap();
        let mut parts = line.split(" ");
        let inst = parts.next().unwrap();
        let num: usize = parts.next().unwrap().parse().unwrap();

        match inst {
            "forward" => dist  += num,
            "down"    => depth += num,
            "up"      => depth -= num,

            _ => panic!("{}", inst),
        }
    }

    println!("{} {} {}", dist, depth, dist*depth)
}
