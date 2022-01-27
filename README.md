# Cracking the Coding Interview, 6th Ed. 

In order to stay sharp, I try to solve a few of these every week. The repository is organized by chapter. Each chapter contains a number of problems. There is a subdirectory for each problem. Generally speaking, I approach a problem by solving it all by myself, then I consult the solution or hints to see ways I could improve upon it. Rather than overwriting my current work when checking the solution, I will create a new package and adjust from there. I will then benchmark and compare my original version to the version I created after viewing the solution or hints. You can view aggregated benchmarks in the table below, but keep in mind that they're only relatively valid -- one solution may be faster than another, but that could change on a more cpu- or memory-constrained machine.

Additionally, in all cases, I recommend actually running the benchmarks for the package in question. I only include the benchmarks for the "base case" in the table, but oftentimes (as mentioned in my comments), there are other benchmarks that test inputs of increasing length or complexity to see how the algorithm scales.

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
    <th>ns/op</th>
    <th>B/op</th>
    <th>allocs/op</th>
    <th>ns/op</th>
    <th>B/op</th>
    <th>allocs/op</th>
  </tr>
  <tr>
    <td>Is Unique</td>
    <td>Implement an algorithm to determine if a string has all unique characters.</td>
    <td>397.5</td>
    <td>96</td>
    <td>1</td>
    <td>N/A</td>
    <td>N/A</td>
    <td>N/A</td>
    <td>My solution was so close to the book solution that I didn't bother revising.</td>
  </tr>
  <tr>
    <td>Check Permutation</td>
    <td>Given two strings, write a method to decide if one is a permutation of the other.</td>
    <td>1984</td>
    <td>496</td>
    <td>6</td>
    <td>3844</td>
    <td>1731</td>
    <td>6</td>
    <td><p>The benchmarks on this one are a bit misleading -- run the full package benchmarks to see the full story. My original solution was to sort both strings and compare them, exiting early if a rune comparison failed. In the best case (strings of unequal length), this runs in O(1) time. In the normal case, it bottlenecks at the sort step with O(log(N)) time, since both sorts have to run before starting comparison.</p>
    <p>The revised solution uses a rune-int map to count the character occurrences of the first string and compare them to the character occurrences of the second string, exiting early if there are any differences. With small strings, this performs worse. However, with larger strings, it begins to shine: As with the first solution, with strings of unequal length, it runs in O(1). With equal-length strings, it runs in O(N) time, where N is the length of both strings.</p></td>
  </tr>
  <tr>
    <td>Urlify</td>
    <td>
      <p>Write a method to replace all spaces in a string with '%20'. You may assume that the given string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.</p>
      <pre>
        EXAMPLE<nbsp>
        Input:   "Mr John Smith    "
        Output:  "Mr%20John%20Smith"
      </pre>
    </td>
    <td>116.6</td>
    <td>48</td>
    <td>2</td>
    <td>105.5</td>
    <td>0</td>
    <td>0</td>
    <td>
      <p>This is one of those cases where I would really advocate for the less-efficient but much more readable algorithm. Using <code>strings.Replace</code> and <code>strings.Trim</code> on the input is two lines and is really fast. That's how it's done in the "clean" version. The "clever" version, using no libraries, avoids allocations by using a rune array and looking backwards through the string. It's so eerily close to the solution the book proposes, I'm not writing a "revised" version.</p>
    </td>
  </tr>
</table>
