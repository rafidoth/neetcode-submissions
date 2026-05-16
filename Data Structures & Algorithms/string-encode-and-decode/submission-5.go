// revise-1

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    encoded := ""
    for _ , str := range strs {
        lenStr := strconv.Itoa(len(str))
        encoded += (lenStr + "#" + str )
    }
    return encoded
}

//         j      i 
// 5#hello5#world_
// 0123456789ABCDE


func (s *Solution) Decode(encoded string) []string {
    var result []string
    for i := 0; i < len(encoded); {
        // finding next delimeter index
        j := strings.Index(encoded[i:], "#")
        if j == -1 {
            break
        }
        j += i // making it absolute for the whole string -> j = 8, i = 7 

        // extracting the length that lives before delimiter
        n, err := strconv.Atoi(encoded[i:j])
        if err != nil {
            fmt.Println("length string to integer conversion failed!")
        }
        str := encoded[j+1:j+n+1]
        result = append(result, str)
        i = j+n+1

    }
    return result
}
