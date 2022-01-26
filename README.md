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
    <th rowspan="2">
      Problem Name
    </th>
    <th rowspan="2">
      Problem Description
    </th>
    <th colspan="3">
      First Solution Benchmarks
    </th>
    <th colspan="3">
      Revision Benchmarks
    </th>
    <th rowspan="2">
      Comments
    </th>
  </tr>
  <tr>
    <th>
      ns/op
    </th>
    <th>
      B/op
    </th>
    <th>
      allocs/op
    </th>
    <th>
      ns/op
    </th>
    <th>
      B/op
    </th>
    <th>
      allocs/op
    </th>
  </tr>
  <tr>
    <td>
      Is Unique
    </td>
    <td>
      Implement an algorithm to determine if a string has all unique characters.
    </td>
    <td>
      397.5
    </td>
    <td>
      96
    </td>
    <td>
      1
    </td>
    <td>
      N/A
    </td>
    <td>
      N/A
    </td>
    <td>
      N/A
    </td>
    <td>
      My solution was so close to the book solution that I didn't bother revising.
    </td>
  </tr>
</table>
