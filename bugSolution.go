func processData(data []int) (float64, error) {
    if len(data) == 0 {
        return 0, errors.New("data slice is empty")
    }

    sum := int64(0) // Use int64 to mitigate overflow
    for _, v := range data {
        sum += int64(v)
    }

    avg := float64(sum) / float64(len(data))
    return avg, nil
}

func main() {
    data := []int{1, 2, 3, 4, 5}
    avg, err := processData(data)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Average:", avg)
    }

    emptyData := []int{}
    avg, err = processData(emptyData)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Average:", avg)
    }
    
    largeData := []int{9223372036854775807, 1}
    avg, err = processData(largeData)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Average:", avg)
    }
}