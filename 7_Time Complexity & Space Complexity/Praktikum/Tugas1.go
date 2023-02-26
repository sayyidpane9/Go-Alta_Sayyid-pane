package main

import "math/rand"

// fungsi yang mengembalikan x^y mod p
func power(x, y, p int) int {
    res := 1
    x = x % p
    for y > 0 {
        if y&1 == 1 {
            res = (res * x) % p
        }
        y = y >> 1
        x = (x * x) % p
    }
    return res
}

// fungsi yang mengembalikan true jika n merupakan bilangan prima,
// false jika n bukan bilangan prima
func primeNumber(n int) bool {
    // kasus khusus untuk bilangan 2 dan 3
    if n == 2 || n == 3 {
        return true
    }
    // kasus khusus untuk bilangan genap
    if n%2 == 0 {
        return false
    }
    // menentukan nilai r dan s sehingga n-1 = 2^r * s
    r, s := 0, n-1
    for s%2 == 0 {
        r++
        s /= 2
    }
    // melakukan pengujian sebanyak 5 kali
    for i := 0; i < 5; i++ {
        // memilih bilangan acak a dalam rentang [2, n-2]
        a := rand.Intn(n-3) + 2
        // menghitung a^s mod n
        x := power(a, s, n)
        // jika hasilnya adalah 1 atau n-1, lanjut ke pengujian berikutnya
        if x == 1 || x == n-1 {
            continue
        }
        // melakukan pengujian sebanyak r-1 kali
        for j := 1; j < r; j++ {
            // menghitung a^(2^j * s) mod n
            x = (x * x) % n
            // jika hasilnya adalah n-1, lanjut ke pengujian berikutnya
            if x == n-1 {
                break
            }
            // jika hasilnya bukan 1 dan bukan n-1, maka n bukan bilangan prima
            if x == 1 {
                return false
            }
        }
        // jika hasil pengujian tidak menunjukkan bahwa n bukan bilangan prima,
        // lanjut ke pengujian berikutnya
        if x != n-1 {
            return false
        }
    }
    // jika semua pengujian berhasil, maka n kemungkinan besar adalah bilangan prima
    return true
}
