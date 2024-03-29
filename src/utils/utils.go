package utils

// type T int
// type K int

var MyTicket = Ticket{
	Name: "Hello",
	Date: 124314,
	Price: 111111,
}

func LogTicket() Ticket {
	return MyTicket;
}

func Find[T any](arr []T, cond func(T) bool) interface{} {
	for i := 0; i < len(arr); i++ {
		el := arr[i]
		if cond(el) {
			return el
		}
	}

	return nil
}

func Filter[T any](arr []T, fn func(value T, ind int) bool) []T {
	var filtered []T
	for i, el := range arr {
		needPush := fn(el, i)
		if needPush {
			filtered = append(filtered, el)
		}
	}
	return filtered
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	found := find(arr, func(num int) bool {
// 		return num >= 5
// 	})

// 	fmt.Println(found)

// 	const str string = "System"
// 	strToLower := strings.ToLower(str)

// 	var strSlice []string
// 	for _, char := range strToLower {
// 		strSlice = append(strSlice, string(char))
// 	}
// 	var sLetterArr []string
// 	sLetterArr = filter(strSlice, func(v string, ind int) bool {
// 		return string(v) == "s"
// 	})
// 	fmt.Println(sLetterArr)
// }
