use std::error::Error;
use std::fs::File;
use std::io::{BufReader, BufRead};

fn main() -> Result<(), Box<dyn Error>> {
    let file = File::open("input.txt")?;
    let lines = BufReader::new(file).lines();
    let mut numbers = Vec::new();

    for line in lines {
        if let Ok(line) = line {
            numbers.push(line.parse::<i32>()?);
        }
    }

    for (first_number_index, first_number) in numbers.iter().enumerate() {
        for (second_number_index, second_number) in numbers.iter().enumerate() {
            if first_number_index != second_number_index {
                if first_number + second_number == 2020 {
                    println!("{}", first_number * second_number);
                    return Ok(());
                }
            }
        }
    }

    Ok(())
}