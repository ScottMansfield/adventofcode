use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;

fn next(buf: &mut [u16; 3], new: u16) -> u16 {
    buf[0] = buf[1];
    buf[1] = buf[2];
    buf[2] = new;

    return buf[0] + buf[1] + buf[2];
}

pub fn part2() {
    let file = File::open("01.input").unwrap();
    let file = BufReader::new(file);

    let mut prev: u16 = 0;
    let mut count: u16 = 0;
    let mut seen: u16 = 0;
    let mut buf: [u16; 3] = [0; 3];

    for line in file.lines() {
        if let Ok(line) = line {
            let temp: u16 = line.parse().unwrap();

            let sum = next(&mut buf, temp);
            // println!("{}", sum);

            seen += 1;
            if seen <= 3 {
                continue;
            }

            if sum > prev {
                count += 1;
            }

            prev = sum;
        }
    }

    println!("{}", count);
}
