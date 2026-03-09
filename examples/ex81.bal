SUB GCD(n, m)
  IF m = 0 THEN
    LET GCD = n
  ELSE
    LET GCD = GCD(m, MOD(n, m))
  END IF
END SUB

SUB MOD(a, b)
  LET MOD = a - (a \ b) * b
END SUB

SUB Main
  PRINT GCD(1, 13)
  PRINT GCD(12, 24)
  PRINT GCD(27, 6)
END SUB
