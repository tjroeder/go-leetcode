# 0/1 Knapsack

# Description
You are given `n` items whose weights and values are known, as well as a knapsack to carry these items. The knapsack cannot carry more than a certain maximum weight, known as its **capacity**.

You need to maximize the total value of the items in your knapsack, while ensuring that the sum of the weights of the selected items does not exceed the capacity of the knapsack.

If there is no combination of weights whose sum is within the capacity constraint, return `0`.

Notes:
```
- An item may not be broken up to fit into the knapsack, i.e., an item either goes into the knapsack in its entirety or not at all.
- We may not add an item more than once to the knapsack.
```

Example 1:
```
capacity: 30
weights: [10, 20, 30]
values: [22, 33, 44]
output: 55 = 22 + 33
```

Example 2:
```
capacity: 5
weights: [1, 2, 3, 4, 5]
values: [10, 5, 4, 8]
output: 15 = 10 + 5
```

Example 3:
```
capacity: 50
weights: [10, 20, 30]
values: [600, 100, 120]
output: 720 = 600 + 120
```

Constraints:
- <code>1 ≤ capacity ≤ 1000</code>
- <code>1 ≤ values.length ≤ 500</code>
- <code>weights.length == values.length</code>
- <code>1 ≤ values[i] ≤ 1000</code>
- <code>1 ≤ weights[i] ≤ capacity</code>
