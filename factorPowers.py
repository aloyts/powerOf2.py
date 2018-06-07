def factorPowers(powInt):

    # =============================================
    # Author: Aloyts
    # Create date: 2018 June 06
    # Description: Factoring powers of 2
    # =============================================

    powCount = powInt
    maxPower = math.floor((math.log2(powInt)))

    print("maxPower=", maxPower, "\n")

    if (math.log2(powInt)%2 != 0) :
        print(powInt, "=\n")
        for i in range(maxPower, -1, -1):
            if powCount-math.pow(2, i) >= 0:
                powCount=powCount-math.pow(2, i)
                print("(2^",i,") = ", math.pow(2, i))
                
#factorPowers(4294967295) #sample run 2^32-1
#factorPowers(895)        #sample random number
#factorPowers(33792)      #sample 2^10 + 2^15
