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
