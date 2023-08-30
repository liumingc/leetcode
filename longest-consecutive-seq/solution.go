package main

func longestConsecutive(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    if n < 1 {
        return 0
    }
    maxN := 1
    seqN := 1
    for i:=1; i<n; i++ {
        if nums[i] == nums[i-1] {
            continue
        }
        if nums[i] - nums[i-1] == 1 {
            // is consecutive
            seqN++
            if seqN > maxN {
                maxN = seqN
            }
            continue
        }
        seqN = 1 // reset
    }
    return maxN
}
