fn parse_input(file: &str) -> Result<Vec<String>, Box<dyn std::error::Error>> {
    let raw_input = std::fs::read_to_string(file)?;
    let mut report: Vec<String> = vec![];
    for line in raw_input.lines() {
        report.push(line.to_owned())
    }

    Ok(report)
}

pub fn solve(file: &str) {
    let input = parse_input(file).unwrap();

    let input_len = input.len();
    let unit_size = input[0].len();
    let mut count = vec![0; unit_size];
    for num in input {
        for (idx, n) in num.chars().enumerate() {
            match n.to_string().as_str() {
                "0" => (),
                "1" => count[idx] += 1,
                _   => panic!("WTF"),
            }
        }
    };

    let mut b_gamma: Vec<String> = vec![String::from(""); unit_size];
    let mut b_epsilon: Vec<String> = vec![String::from(""); unit_size];

    for (idx, c) in count.iter().enumerate() {
        b_gamma[idx] = format!("{}", (*c as f32 / input_len as f32).round() as u32);
        b_epsilon[idx] = format!("{}", ((input_len - *c) as f32 / input_len as f32).round() as u32);
    }

    // Convert to decimal
    let gamma = usize::from_str_radix(b_gamma.join("").as_str(), 2).unwrap() as u32;
    let epsilon = usize::from_str_radix(b_epsilon.join("").as_str(), 2).unwrap() as u32;

    println!("[3.1] gamma * epsilon = {}", gamma * epsilon);
    //println!("[3.2] gamma * epsilon = {}", gamma * epsilon);
}
