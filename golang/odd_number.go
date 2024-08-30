package main

func FindOddNumber(numbers []int) int {

    mappingNumbers := make(map[int]int)
    for _, number := range numbers {
		mappingNumbers[number]++
	}

    answerArray := []int{}
    for key, value := range mappingNumbers {
        if value %2 != 0 {
            answerArray = append(answerArray, key)
        }
    }

    if len(answerArray) == 1 {
        return answerArray[0]
    }

    //unexpected error
    return -1
}
