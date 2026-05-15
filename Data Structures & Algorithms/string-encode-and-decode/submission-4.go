// just a clean code version of solution 1
//////////////////////////////////////////
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    var singleStringsEncoded []string
    for _, s := range strs {
        lenStr := strconv.Itoa(len(s))
        singleStringEncoded := lenStr + "#" + s
        singleStringsEncoded = append(singleStringsEncoded, singleStringEncoded)
    }
    return strings.Join(singleStringsEncoded, "")
}

func (s *Solution) Decode(encoded string) []string {
    var result []string
    for i:= 0; i < len(encoded); {
        // next delimiter index 
        j := strings.Index(encoded[i:], "#")
        if j == -1 {
            break
        }
        j += i// making j correct for the whole encoded string


        n, err := strconv.Atoi(encoded[i:j])
        if err != nil {
            fmt.Println("failed to convert length from str to int")
        }

        // extracting next n characters as its a string 
        stringStart := j+1
        stringEnd := j+n 
        str := encoded[stringStart : stringEnd +1]

        // append the string to result
        result = append(result, str)

        // moving i to first of next string msg
        i = stringEnd+1 
    }
    return result 
}
