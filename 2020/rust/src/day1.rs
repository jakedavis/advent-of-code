fn parse_input(contents: &String) -> Result<Vec<i32>, std::num::ParseIntError> {
    let entries = contents
        .lines()
        .map(|l| l.parse().unwrap())
        .collect();

    Ok(entries)
}

fn find_candidates(entries: &Vec<i32>, arity: u32, sum: i32) -> Option<i32> {
    for i in entries {
        for j in entries {
            if arity == 2 {
                if i + j == sum {
                    return Some(i * j);
                }
            } else if arity == 3 {
                for k in entries {
                    if i + j + k == sum {
                        return Some(i * j * k);
                    }
                }
            }
        }
    }

    None
}

pub fn part1(input: &String) -> Option<i32> {
    let entries = parse_input(&input).unwrap();
    find_candidates(&entries, 2, 2020)
}

pub fn part2(input: &String) -> Option<i32> {
    let entries = parse_input(&input).unwrap();
    find_candidates(&entries, 3, 2020)
}
