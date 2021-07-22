// main.rs

use std::io::prelude::*;

fn main() -> std::io::Result<()> {
    let f = std::fs::File::open("/Users/jake/dev/github.com/jakedavis/advent-of-code/2018/1_debug")?;
    let reader = std::io::BufReader::new(f);
    let mut seen: Vec<i32> = Vec::new();
    let mut freq: i32 = 0;

    for line in reader.lines() {
        let current: i32 = line
            .expect("Line iterator failed ...")
            .parse()
            .expect("Couldn't parse integer ...");

        freq = freq + current;
        if seen.contains(&freq) {
            println!("The first frequency seen twice is {}", freq);
            break;
        }

        seen.push(freq);
    }

    println!("Seen: {:?}", seen);
    println!("The result is {}", seen.get(seen.len()-1).expect("Wrong index!"));

    Ok(())
}
