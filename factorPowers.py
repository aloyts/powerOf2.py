import math
#powInt = 4294967295
#powInt = 895
powCount = powInt
maxPower = math.floor((math.log2(powInt)))

print("maxPower=", maxPower, "\n")

if (math.log2(powInt)%2 != 0) :
    print(powInt, "=\n")
    for i in range(maxPower, -1, -1):
        if powCount-math.pow(2, i) >= 0:
            powCount=powCount-math.pow(2, i)
            print("(2^",i,") = ", math.pow(2, i))
