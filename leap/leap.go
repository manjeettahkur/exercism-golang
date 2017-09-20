package leap

const testVersion = 3

func IsLeapYear(year int) bool {
  if year%4 == 0 {
    if year%100 == 0 {
      if year%400 == 0 {
        return true
      } else {
        return false
      }
    } else {
      return true
    }
  } else {
    return false
  }
}
