program {
  var a = 11;
  var b = 10;
  var c = 5;

  if a == b {
    if a % c == 0 {
      c = a + 2;
    } else {
      c = a - 3;
    }
  } else {
    c = a - b;
  }
}