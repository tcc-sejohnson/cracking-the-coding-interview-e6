# Cracking the Coding Interview, 6th Ed. 

In order to stay sharp, I try to solve a few of these every week. The repository is organized by chapter. Each chapter contains a number of problems. There is a subdirectory for each problem. Generally speaking, I approach a problem by solving it all by myself, then I consult the solution or hints to see ways I could improve upon it. Rather than overwriting my current work when checking the solution, I will create a new package and adjust from there. I will then benchmark and compare my original version to the version I created after viewing the solution or hints. You can view aggregated benchmarks in the table below, but keep in mind that they're only relatively valid -- one solution may be faster than another, but that could change on a more cpu- or memory-constrained machine.

## Benchmarking
```bash
# Assumes you're in the probject root

# Run all tests, including benchmarks
go test -bench . ./...

# Run tests for a specific subdirectory
DIR='ch1/1.1'
go test -v ./$DIR/...
```

## Results

### Chapter 1 - Arrays and Strings

<table>
  <tr>
    <th>
      Problem Name
    </th>
    <th>
      Problem Description
    </th>
    <th>
      First Solution Benchmarks
    </th>
    <th>
      Revision Benchmarks
    </th>
    <th>
      Comments
    </th>
  </tr>
</table>
