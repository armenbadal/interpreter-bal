' կետի կոնստրուկտորը
SUB NewPoint(x, y)
  LET NewPoint = [x, y]
END SUB

' կետի x կոորդինատը
SUB XOf(point)
  LET XOf = point[0]
END SUB

' կետի y կոորդինատը
SUB YOf(point)
  LET YOf = point[1]
END SUB

' երկու կետերի հեռավորությունը
SUB Distance(a, b)
  LET dx = XOf(a) - XOf(b)
  LET dy = YOf(a) - YOf(b)
  LET Distance = SQR(dx^2 + dy^2)
END SUB

' շրջանի կոնստրուկտորը
SUB NewCircle(center, radius)
  LET NewCircle = [center, radius]
END SUB

' շրջանի կենտրոնը
SUB CenterOf(circle)
  LET CenterOf = circle[0]
END SUB

' շրջանի շառավիղը
SUB RadiusOf(circle)
  LET RadiusOf = circle[1]
END SUB

' կետի՝ շրջանի մեջ լինելը
SUB InCircle(circle, point)
  LET center = CenterOf(circle)
  LET dist = Distance(center, point)

  LET radius = RadiusOf(circle)
  LET InCircle = dist < radius
END SUB

' կետի՝ շրջանի մեջ լինելը (այլ տարբերակ)
SUB InCircle2(circle, point)
  LET Incircle = Distance(CenterOf(circle), point) < RadiusOf(circle)
END SUB


SUB Main
  LET circle = NewCircle(NewPoint(0, 0), 2)

  LET point1 = NewPoint(1.2, 1.2)
  PRINT InCircle(circle, point1)

  LET point2 = NewPoint(-4, 0)
  PRINT InCircle(circle, point2)
END SUB

