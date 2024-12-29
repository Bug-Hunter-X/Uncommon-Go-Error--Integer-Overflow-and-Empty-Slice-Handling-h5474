# Uncommon Go Error: Integer Overflow and Empty Slice Handling

This repository demonstrates a Go program that calculates the average of a slice of integers. It highlights two potential issues:

1. **Integer Overflow:** If the sum of the integers exceeds the maximum value of an `int`, an overflow occurs, leading to incorrect results.  This is more likely with large datasets.
2. **Empty Slice:** The program doesn't gracefully handle the case where the input slice is empty. This can cause a runtime panic (division by zero).

The `bug.go` file shows the code with the errors. The `bugSolution.go` file presents a corrected version that addresses both issues using appropriate error handling and potentially larger integer types (e.g., `int64`).