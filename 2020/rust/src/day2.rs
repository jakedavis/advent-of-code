struct PasswordPolicy {
    occurrences: (i32, i32),
    letter: char,
    password: String,
}

impl PasswordPolicy {
    fn valid_numerically(&self) -> bool {
        let occurs = self.password
            .chars()
            .filter(|c| self.letter == *c)
            .count() as i32;

        let (low, high) = self.occurrences;
        low <= occurs && occurs <= high
    }

    fn valid_positionally(&self) -> bool {
        let chars: Vec<char> = self.password.chars().collect();
        let (low, high) = self.occurrences;

        (self.letter == chars[(low as usize)-1]) ^ (self.letter == chars[(high as usize)-1])
    }
}

fn process_line(line: &str) -> PasswordPolicy {
    let splits: Vec<&str> = line.split(" ").collect();
    let range:  Vec<i32>  = splits[0]
        .split("-")
        .map(|r| r.parse::<i32>().unwrap())
        .collect();

    PasswordPolicy {
        occurrences: (range[0], range[1]),
        letter: splits[1].chars().next().unwrap(),
        password: splits[2].to_string(),
    }
}

fn parse(input: &String) -> Vec<PasswordPolicy> {
    input
        .lines()
        .map(|l| process_line(&l))
        .collect()
}

pub fn part1(input: &String) -> Option<i32> {
    Some(parse(input).iter().filter(|p| p.valid_numerically()).count() as i32)
}

pub fn part2(input: &String) -> Option<i32> {
    Some(parse(input).iter().filter(|p| p.valid_positionally()).count() as i32)
}
