```go
func f_rec(n int) uint64 {
    if n == 1 || n == 2 {
        return 1
    }
    return f_rec(n-1) + f_rec(n-2)
}
```

1. `f_rec(7) = 13`
2. I assume that the return type is declared as `uint64` because the output of the function can be very large and is always positive.
3. The function returns the nth number of the Fibonacci sequence.

```go
func f_iter1(n int) uint64 {
    var f uint64
    var f1 , f2 uint64 = 1, 1

    if n == 2 || n == 1 {
        return 1
    }

    for n >= 3 {
        n--
        f = f1 + f2
        f1 = f2
        f2 = f
    }
    return f
}
```

```go
func f_iter2(n int) uint64 {
    var f uint64
    var f1 , f2 uint64 = 1, 1

    if n == 2 || n == 1 {
        return 1
    }

    for i := 2; i <= n; i ++ {
        f = f1 + f2
        f1 = f2
        f2 = f
    }
    return f
}
```

4. Are the two functions equivalent?

   - `f_iter1(0) = f_iter2(0) = 0`
   - `f_iter1(1) = f_iter2(1) = 1`
   - `f_iter1(2) = f_iter2(2) = 1`
   - `f_iter1(3) = 2` and `f_iter2(3) = 3`
   - The two functions are not equivalent. `f_iter2` gives the (n+1)th element of the Fibonacci sequence

5. `f_iter1` is equivalent to `f_rec` for n>0.
6. To make `f_iter2` equivalent to `f_rec` (for n>0) we have to change the condition of the for loop to `i < n`.
7. Operations:

   - `f_rec`: time complexity `O(2n-2)=O(n)`, space complexity: `O(2n)=O(n)`;
   - `f_irer1`: time complexity `O(n)`, space complexity: `O(1)`;
   - `f_irer2`: time complexity `O(n)`, space complexity: `O(1)`;

```go
func f_riter(a, b uint64, n int) uint64 {
    if n == 2 {
        return a
    }

    if n == 1 {
        return b
    }

    return f_riter(a+b, a, n-1)
}
```

8. `f_riter(1, 1, n)` is equivalent to `f_rec(n)` for n>0.
9. Execution of `f_rec(7)` and `f_riter(1, 1, 7)`:

   ```go
   // f_rec(7)
   // f_rec(6) + f_rec(5)
   // f_rec(5) + f_rec(4) + f_rec(4) + f_rec(3)
   // f_rec(4) + f_rec(3) + f_rec(3) + f_rec(2) + f_rec(3) + f_rec(2) + (f_rec(2) + f_rec(1)
   // f_rec(3) + f_rec(2) + f_rec(2) + f_rec(1) + (f_rec(2) + f_rec(1)) + 1 + f_rec(2) + f_rec(1) + 1 + 1 + 1
   // f_rec(2) + f_rec(1) + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1
   // 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1
   // 13

   // f_riter(1, 1, 7)
   // f_riter(2, 1, 6)
   // f_riter(3, 2, 5)
   // f_riter(5, 3, 4)
   // f_riter(8, 5, 3)
   // f_riter(13, 8, 2)
   // 13
   ```

10. The space complexity of `f_riter` is `O(n)`.
