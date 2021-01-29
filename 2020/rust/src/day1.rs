fn open(path: &str) -> Result<String, std::io::Error> {
    let f = std::fs::read(path)?;
    let s = String::from_utf8(f).expect("Invalid UTF-8");

    Ok(s)
}

fn parse(contents: &String) -> Result<Vec<i32>, std::num::ParseIntError> {
    let entries = contents
        .lines()
        .map(|l| l.parse().unwrap())
        .collect();

    Ok(entries)
}

pub fn run() -> Result<Vec<i32>, Box<dyn std::error::Error>> {
    let handle  = open("../1_input")?;
    let entries = parse(&handle)?;

    let mut result: Vec<i32> = vec![];
    for i in &entries {
        for j in &entries {
            for k in &entries {
                if i + j + k == 2020 {
                    if !result.contains(&(i * j * k)) {
                        result.push(i * j * k);
                    }
                }
            }
        }
    }

    Ok(result)
}
