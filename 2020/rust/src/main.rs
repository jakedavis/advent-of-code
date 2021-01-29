mod day1;

fn main() {
    match day1::run() {
        Ok(result) => println!("{:?}", result),
        Err(error) => panic!("{}", error),
    }
}
