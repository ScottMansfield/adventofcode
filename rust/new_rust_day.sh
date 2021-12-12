mkdir -p $1/$2
cd $1/$2

# Touch example and input files
touch $2.example
touch $2.input

# Generate src dir and cargo.toml
cargo init --bin --name day$2

# Ignore output dir
echo "/target" > .gitignore

# Create day structure
cat > src/main.rs<< EOF
mod part1;
mod part2;

fn main() {
    part1::part1();
    part2::part2();
}
EOF

cat > src/part1.rs<< EOF
pub fn part1() {
    println!("Hello, Part1!");
}
EOF

cat > src/part2.rs<< EOF
pub fn part2() {
    println!("Hello, Part2!");
}
EOF

cargo run
