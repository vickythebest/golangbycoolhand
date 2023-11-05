package advance

func rotateArray(array []int, postition int) {

	length := len(array)

	postition %= length

	rotateArray := make([]int, length)

	// fmt.Println("postition : ", postition)
	for i := 0; i < length; i++ {
		// fmt.Println(i, (i - postition), (i - postition + length), length, " New Index : ", (i-postition+length)%length)

		newIndex := (i - postition + length) % length
		rotateArray[i] = array[newIndex]
	}

	display(rotateArray)
}
