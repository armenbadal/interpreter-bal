SUB LinearSearch(arr, elem)
  LET length = LEN(arr)

  LET i = 0
  WHILE i < length AND arr[i] <> elem
	  LET i = i + 1
  END WHILE

  LET LinearSearch = NOT (i = length)
END SUB

SUB Main
  LET numbers = [0, 1, 2, 3, 4, 5, 7, 8, 9]
  PRINT LinearSearch(numbers, 6)
  PRINT LinearSearch(numbers, 2)

  LET words = ["this", "is", "a", "list", "of", "words"]
  PRINT LinearSearch(words, "of")
  PRINT LinearSearch(words, "ok")
END SUB
