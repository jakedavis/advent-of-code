mod day1;
fn get_input(path: &str) -> String {
    let f = std::fs::read(path).expect("Invalid path");
    String::from_utf8(f).expect("Invalid UTF-8")
}

fn main() {
    match day1::run() {
        Ok(result) => println!("{:?}", result),
        Err(error) => panic!("{}", error),
    }
}
