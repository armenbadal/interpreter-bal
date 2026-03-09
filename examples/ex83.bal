SUB QuickSort(arr)
  IF LEN(arr) <= 1 THEN
    LET QuickSort = arr
  ELSE
    LET parted = Partition(arr)

    LET left = QuickSort(parted[0])
    LET right = QuickSort(parted[2])

    LET QuickSort = Join(left, parted[1], right)
  END IF
END SUB

SUB Partition(arr)
    LET pivot = arr[0]
    
    LET left = []
    LET right = []
    FOR i = 1 TO LEN(arr) - 1
      LET e = arr[i]
      IF e < pivot THEN
        LET left = Append(left, e)
      ELSE
        LET right = Append(right, e)
      END IF
    END FOR

    LET Partition = [left, pivot, right]
END SUB

SUB Append(a, e)
  DIM res[LEN(a) + 1]
  FOR i = 0 TO LEN(a) - 1
    LET res[i] = a[i]
  END FOR
  LET res[LEN(a)] = e
  LET Append = res
END SUB

SUB Join(a, x, b)
  DIM res[LEN(a) + 1 + LEN(b)]
  
  LET j = 0
  FOR i = 0 TO LEN(a) - 1
    LET res[j] = a[i]
    LET j = j + 1
  END FOR

  LET res[j] = x
  LET j = j + 1

  FOR i = 0 TO LEN(b) - 1
    LET res[j] = b[i]
    LET j = j + 1
  END FOR

  LET Join = res
END SUB

SUB Main
  LET arr = [4, 3, 2, 0, 6, 1, 5]
  LET sorted = QuickSort(arr)
  PRINT sorted
END SUB
