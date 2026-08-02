# `cut` - [coding challenge #4] cut Tool

## 🧐 About Repo

### In this repo, I build my own version of the Unix command line tool: `cut`

Part of this Coding challenge: https://codingchallenges.fyi/challenges/challenge-cut

**Written in:** Go 🔵

## 👨‍🏫 Set-up & Build

1. Clone the repo:

   ```terminal
   $ git clone https://github.com/leekli/coding-challenges
   ```

2. Change directory:

   ```terminal
   $ cd cut
   ```

3. Ensure at least Go **1.23.1** is installed.

4. Build the Go Binary:

   ```terminal
   $ go build cut.go
   ```

5. Two test files are supplied in the `test-data/` folder to use as inputs/examples.

> [!NOTE]
> There is a test suite available and can be ran using `go test -v`

## 💻 Commands & Options available

> [!NOTE]
> All commands below can be used by giving a file name.

> [!NOTE]
> All output can be piped to the Standard Output (stdout).

- **-f[n] [file_name]:** The list specifies fields, separated in the input by the field delimiter character (see the -d option). Output fields are separated by a single occurrence of the field delimiter character.

  - With a file name:

  ```terminal
  $ ./cut -f2 test-data/sample.tsv
  ```

<br></br>

- **-d [file_name]:** Use delim as the field delimiter character instead of the tab character.

  - With a file name:

  ```terminal
  $ ./cut -f2 -d, test-data/fourchords.csv
  ```

  You can pipe the output to stdout as such:

  ```terminal
  $ ./cut -f2 -d, test-data/fourchords.csv | head -n5
  ```

<br></br>
