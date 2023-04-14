func isPalindrome(x int) bool {
    if (x < 0 || (x % 10 == 0 && x != 0)) {
        return false
    } 
    var revertedNumber int
    for x > revertedNumber {
        revertedNumber = revertedNumber * 10 + x % 10
        x /= 10
    }
    return x == revertedNumber || x == revertedNumber/10
}

