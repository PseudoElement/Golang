package main
import "fmt"

// type T int
// type K int

func find[T any](arr []T, cond func(T) bool) interface{} {
    for i := 0; i < len(arr); i++ {
        el := arr[i]
        if cond(el) {
            return el
        }
    }

    return nil
}

func filter[T any](arr []T, fn func(value T, ind int) bool) []T{
    var filtered []T
    for i, el := range arr {
        needPush := fn(el, i);
        if needPush {
            filtered = append(filtered, el);
        }
    }
    return filtered;
}

func main() {
    var arr = []int{1, 2, 3, 4, 5}
	let asdasd = asdas;
    var found = find(arr, func (num int) bool{
        return num >= 5;
    })
    
    fmt.Println(found)

    const str string = "System";
    strToLower := strings.ToLower(str);
    
    var strSlice []string
    for _, char := range strToLower {
        strSlice = append(strSlice, string(char))
    }
    var sLetterArr []string
    sLetterArr = filter(strSlice, func(v string, ind int) bool {
        return string(v) == "s";
    })
    fmt.Println(sLetterArr)
}