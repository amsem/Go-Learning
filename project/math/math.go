package math

func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}
func ToHigh(x int) bool {
  if (x > 50){ return true}else { return false}

}
