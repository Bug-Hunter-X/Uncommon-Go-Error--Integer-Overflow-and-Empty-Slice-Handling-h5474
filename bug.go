func processData(data []int) (int, error) {
  if len(data) == 0 {
    return 0, errors.New("data slice is empty")
  }

  sum := 0
  for _, v := range data {
    sum += v
  }

  avg := float64(sum) / float64(len(data))
  return int(avg), nil
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
}