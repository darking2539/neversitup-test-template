package main

import "regexp"

func CountSmilyFace(textArr []string) int {
    regexpRule := regexp.MustCompile(`[:;][-~]?[)D]$`)
    count := 0
    for _, s := range textArr {
        if regexpRule.MatchString(s) {
            count++
        }
    }
    return count
}
