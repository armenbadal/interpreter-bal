
SUB NumberToText(num)
  IF num < 999999 THEN
    LET h = num \ 1000
    LET t = num - h * 1000

    IF h = 0 THEN
      LET NumberToText = helperNumberToText(t)
    ELSE
      LET NumberToText = helperNumberToText(h) & " հազար " & helperNumberToText(t)
    END IF
  ELSE
    LET NumberToText = "Շատ մեծ թիվ է։"
  END IF
END SUB

SUB helperNumberToText(num)
  LET ones = ["", "մեկ", "երկուս", "երեք", "չորս", "հինգ", 
              "վեց", "յոթ", "ութ", "ինը"]
  LET tens = ["", "տասն", "քսան", "երեսուն", "քառասուն", "հիսուն",
              "վաթսուն", "յոթանասուն", "ութսուն", "իննսուն"]

  IF num < 10 THEN  
    LET helperNumberToText = ones[num]
  ELSEIF num < 100 THEN
    LET t = num \ 10
    LET o = num - t * 10
    LET helperNumberToText = tens[t] & ones[o]
  ELSEIF num < 1000 THEN
    LET h = num \ 100
    LET ot = num - h * 100
    LET helperNumberToText = ones[h] & " հարյուր " & helperNumberToText(ot)
  END IF
END SUB

SUB TestNumberToText(number, expected)
  LET status = "FAIL"

  LET text = NumberToText(num)
  IF text = expected THEN
    LET status = "PASS"
  END IF

  PRINT status & ": " & STR(number) & " -> «" & text & "»"
END SUB

SUB Main
  LET numbers = [[8, "ութ"],
                 [28, "քսանութ"],
                 [108, "մեկ հարյուր ութ"],
                 [310, "երեք հարյուր տասն"],
                 [456, "չորս հարյուր հիսունվեց"],
                 [100, "մեկ հարյուր"],
                 [900, "ինը հարյուր"], 
                 [999, "ինը հարյուր իննսունինը"],
                 [1000, "մեկ հազար"]]
  FOR i = 0 TO LEN(numbers) - 1
    LET num = numbers[i][0]
    LET exp = numbers[i][1]
    CALL TestNumberToText num, exp
  END FOR
END SUB
