package main

import "sort"

func Permutations(text string) []string {

	var result []string

	//if array equals 0 return empty array
	if len(text) == 0 {
		return result
	}

	//change string to byte for select easy to swap position
	bytes := []byte(text)

	//sorting for left handside is lowest byte (dcba -> abcd)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	length := len(bytes)

	for {
		result = append(result, string(bytes))

		i := length - 2 // select ab[c]d
		j := length - 1 // select abc[d]

		for i >= 0 && bytes[i] >= bytes[i+1] {
			i--
		}

		if i < 0 {
			break
		}

		for bytes[j] <= bytes[i] {
			j--
		}

		// Swap characters at i and j
		bytes[i], bytes[j] = bytes[j], bytes[i]

		// Reverse the sequence after the position i
		for l, r := i+1, length-1; l < r; l, r = l+1, r-1 {
			bytes[l], bytes[r] = bytes[r], bytes[l]
		}

	}

	return result
}

