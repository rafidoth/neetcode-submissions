type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var lenArr []int 
    for _, str := range strs {
        lenArr = append(lenArr, len(str))
    }
    if len(lenArr) != len(strs) {
        fmt.Println("Length array and original array doesn't match in length!")
        return ""
    }

    e := ""
    for indx , str := range strs {
        e += strconv.Itoa(lenArr[indx])
        e += "#"
        e += str
    }
    fmt.Println("encoded ", e)

    return e
}

func (s *Solution) Decode(encoded string) []string {
    var d []string
    length := ""
    lengthInt := 0
    str := ""
    flag := true   // true -> length , false -> string


    for _, r := range encoded {
        if flag {
            if r == '#'{
                flag = false
                integer, err := strconv.Atoi(length)
                if err != nil {
                    fmt.Println("Invalid length, conversion error!!", length)
                }
                lengthInt = integer
                if lengthInt == 0 {
                    d = append(d, "")
                    flag = true
                }
                length = ""
                continue
            }else{
                length += string(r)
            }
        }else{
            if lengthInt > 1 {
                str += string(r)
                lengthInt--
            }else if lengthInt == 1 {
                flag = true
                str += string(r)
                lengthInt--
                d = append(d, str)
                str = ""
                length = ""
            }else {
                fmt.Println("Length Integer error")
            }
        }
    }
    return d
}
