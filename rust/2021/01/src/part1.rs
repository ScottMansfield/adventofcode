use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;

pub fn part1() {
    let file = File::open("01.input").unwrap();
    let file = BufReader::new(file);

    let mut prev: u16 = 0;
    let mut count: u16 = 0;

    for line in file.lines() {
        if let Ok(line) = line {
            let temp: u16 = line.parse().unwrap();

            if temp > prev {
                count += 1;
            }

            prev = temp;
        }
    }

    println!("{}", count);
}
