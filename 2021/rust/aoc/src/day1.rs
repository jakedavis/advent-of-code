use std::cmp::Ordering;

/// Takes a path to a file (the input) and parses it according to this day's problem. Since every
/// problem, and therefore its data, is meaningfully different, we implement this each day.
fn parse_input(file: &str) -> Result<Vec<u32>, Box<dyn std::error::Error>> {
    let mut report: Vec<u32> = vec![];

    let input = std::fs::read_to_string(file)?;
    for line in input.lines() {
        let parsed = line.parse::<u32>()?;
        report.push(parsed);
    };

    Ok(report)
}

/// Finds the number of elevation increases for the given report, broken into a Windows instance.
/// The elevation is said to be increasing if the first element of the window is less than the
/// second.
fn elevation_increases(input: &Vec<u32>) -> u32 {
    input.windows(2).map(|it|
        match it[0].cmp(&it[1]) {
            Ordering::Less => 1,
            _ => 0,
        }
    ).sum()
}

pub fn solve(file: &str) {
    let report = parse_input(file).unwrap();
    let sum_2 = elevation_increases(&report);

    let report_3 = report.windows(3).map(|it| it.iter().sum()).collect();
    let sum_3 = elevation_increases(&report_3);

    println!("[1.1] Elevation increases (batches of 1): {sum_2}");
    println!("[1.2] Elevation increases (batches of 3): {sum_3}");
}
