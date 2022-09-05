# `stdiotest`

stdiotest is a testing utility which tests the stdout output of a given program with specified stdin. It also supports running the tests in parallel to improve efficiency.

## How to use

1. Install the cli tool: `go install github.com/GalvinGao/stdiotest@latest`
2. Create a `stdiotest.yaml` file in the root of your project. The file should look like this:

```yaml
test_cases:
  - cmd: "./main"
    exit_code: 0
    stdin: "Galvin"
    stdout: |
      Hello, Galvin!

  - cmd: "./main"
    exit_code: 0
    stdin: "Alice"
    stdout: |
      Hello, Alice!

  - cmd: "./main"
    exit_code: 1
    stdin: "Bob"
    stdout: |
      Bob is not allowed to use this program.-

```

3. Run `stdiotest run` in the root of your project.

See `stdiotest help` for more information.
