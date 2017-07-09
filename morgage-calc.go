package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  fmt.Println("This is the first test")
  choice := getUserDecision()
  if choice == "income" {
    incomeCalculation()
  } else if choice == "value" {
    valueCalculation()
  }
}

func getUserDecision() string {
  reader := bufio.NewReader(os.Stdin)
  text := "empty"
  for text != "income" && text != "value" {
    fmt.Println("Would you like to calculate your income or the value of the home? (income or value)")
    text, _ = reader.ReadString('\n')
    text = strings.TrimSpace(text)
    if text == "income" || text == "value" {
      fmt.Println("Thanks!")
    } else {
      fmt.Println("I didn't understand that.")
    }
  }
  return text
}

func incomeCalculation() float32 {
  reader := bufio.NewReader(os.Stdin)
  text := "empty"
  fmt.Println("What is the value of the home?")
  text, _ = reader.ReadString('\n')
  text = strings.TrimSpace(text)
  value, _ := strconv.Atoi(text)
  downPayment := float32(value) * 0.05
  repay := float32(value) - downPayment
  fmt.Println(repay)
  return repay
}

func valueCalculation() float32 {
  reader := bufio.NewReader(os.Stdin)
  text := "empty"
  fmt.Println("What is your income?")
  text, _ = reader.ReadString('\n')
  text = strings.TrimSpace(text)
  income, _ := strconv.Atoi(text)
  incomeF := float32(income)
  lowTotal := incomeF * 30.0 * 0.15
  highTotal := incomeF * 30.0 * 0.45
  medTotal := incomeF * 30.0 * 0.25
  fmt.Println("Range from " + strconv.FormatFloat(float64(lowTotal), 'f', -1, 32) + " to " + strconv.FormatFloat(float64(highTotal), 'f', -1, 32))
  fmt.Println("Ideal value: " + strconv.FormatFloat(float64(medTotal), 'f', -1, 32))
  
  return medTotal
}