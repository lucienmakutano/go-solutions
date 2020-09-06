package roman_num_decoder

func Decode(roman string) int {
  r_chars := []byte(roman)

  var res int
  var s_val string
  var s_val_1 string
  length := len(r_chars)

  r_v := map[string]int{
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
  }

  for idx, val := range r_chars {

    if (idx + 1) == length {
      s_val = string(val)
      s_val_1 = string(val)
    } else {
      s_val = string(val)
      s_val_1 = string(r_chars[idx + 1])
    }

    if r_v[s_val] >= r_v[s_val_1] {
      res += r_v[s_val]
    } else {
      res -= r_v[s_val]
    }

  }

  return res
}

