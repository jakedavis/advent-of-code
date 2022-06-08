/// Represents the possible directions in which the submarine can go. Each variant also takes a u32
/// that represents the amount to move for the given direction.
enum Direction {
    Forward(u32),
    Down(u32),
    Up(u32),
}

/// Methods required to interpret a report which contains directions with magnitudes.
trait Positionable {
    fn forward(&self, amt: i32) -> Self;
    fn up(&self, amt: i32) -> Self;
    fn down(&self, amt: i32) -> Self;
    fn resolve_position(&self) -> i32;
}

/// Submarine with different semantics for interpreting directions: in particular, an aim field
/// dictates the degree to which the depth is adjusted.
struct AimedSubmarine {
    horizontal: i32,
    depth: i32,
    aim: i32,
}

impl AimedSubmarine {
    fn new() -> Self {
        AimedSubmarine {
            horizontal: 0,
            depth: 0,
            aim: 0,
        }
    }
}

impl Positionable for AimedSubmarine {
    /// Add the amount to the horizontal position and increase the depth by the aim multiplied by
    /// the amount.
    fn forward(&self, amt: i32) -> AimedSubmarine {
        AimedSubmarine {
            horizontal: self.horizontal + amt,
            depth: self.depth + self.aim * amt,
            aim: self.aim,
        }
    }

    /// Decrease the aim of the submarine
    fn up(&self, amt: i32) -> AimedSubmarine {
        AimedSubmarine {
            horizontal: self.horizontal,
            depth: self.depth,
            aim: self.aim - amt,
        }
    }

    /// Increase the aim of the submarine
    fn down(&self, amt: i32) -> AimedSubmarine {
        AimedSubmarine {
            horizontal: self.horizontal,
            depth: self.depth,
            aim: self.aim + amt,
        }
    }

    fn resolve_position(&self) -> i32 {
        self.horizontal * self.depth
    }
}

/// Representation of where the submarine is. Has a horizontal quantity, which moves permanently in
/// the positive direction, and a depth. Note going up decreases the depth and vice versa for down.
struct Submarine {
    horizontal: i32,
    depth: i32,
}

impl Submarine {
    /// Creates a new Position with initial values 0, 0
    fn new() -> Self {
        Submarine {
            horizontal: 0,
            depth: 0,
        }
    }
}

impl Positionable for Submarine {
    /// Modify the horizontal component of the position
    fn forward(&self, amt: i32) -> Submarine {
        Submarine {
            horizontal: self.horizontal + amt,
            depth: self.depth,
        }
    }

    /// Decrease the depth of the submarine
    fn up(&self, amt: i32) -> Submarine {
        Submarine {
            horizontal: self.horizontal,
            depth: self.depth - amt,
        }
    }

    /// Increase the depth of the submarine
    fn down(&self, amt: i32) -> Submarine {
        Submarine {
            horizontal: self.horizontal,
            depth: self.depth + amt,
        }
    }

    /// We'll be asked to multiply the depth by the horizontal value as part of the exercise, which
    /// we lovingly call the "weird calculation".
    fn resolve_position(&self) -> i32 {
        self.horizontal * self.depth
    }
}

/// Takes an file path as input and returns a Result with either a vector of Direction enums or
/// some variety of error.
fn parse_input(file: &str) -> Result<Vec<Direction>, Box<dyn std::error::Error>> {
    let raw_input = std::fs::read_to_string(file)?;
    let mut directions: Vec<Direction> = vec![];

    for line in raw_input.lines() {
        let mut splits = line.split(" ");
        let dir = splits.next().unwrap();
        let amt = splits.next().unwrap().parse::<u32>()?;

        let d = match dir {
            "forward" => Direction::Forward(amt),
            "down"    => Direction::Down(amt),
            "up"      => Direction::Up(amt),
            _         => panic!("IDK"),
        };

        directions.push(d)
    }

    Ok(directions)
}

pub fn solve(file: &str) {
    let directions = parse_input(file).unwrap();
    let mut p1 = Submarine::new();
    let mut p2 = AimedSubmarine::new();

    for direction in directions {
        p1 = match direction {
            Direction::Forward(n) => p1.forward(n as i32),
            Direction::Down(n)    => p1.down(n as i32),
            Direction::Up(n)      => p1.up(n as i32)
        };

        p2 = match direction {
            Direction::Forward(n) => p2.forward(n as i32),
            Direction::Down(n)    => p2.down(n as i32),
            Direction::Up(n)      => p2.up(n as i32)
        };
    };

    println!("[2.1] Standard submarine final position: {}", p1.resolve_position());
    println!("[2.2] Aimed submarine final position: {}", p2.resolve_position());
}
