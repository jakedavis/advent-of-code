// 1_frequency.rs

use std::io;
use std::io::BufReader;
use std::io::File;
use std::io::Lines;

fn main() {
    let f = File::open("../1_input");
    let mut reader = BufReader::new(f);
    let mut buffer = String::new();
    let freq = 0

    for line in reader.lines() {
        freq = freq + line.parse();
    }

    println!("The result is {}", freq);
}
