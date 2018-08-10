/*
 =============================================
 Author:      Aloyts
 Create date: 2018 August 06
 Description: Factoring powers of 2
 =============================================
*/
package main

import ("fmt")
import ("math")

func main() {
	var factors [50]float64
	
	var totalVal float64 = 17
	var powCount = totalVal
	
	var maxPower float64
	var rem float64
	
	//var i float64
	var j int = 0
	
	maxPower = math.Floor(math.Log2(totalVal))
	rem = totalVal - math.Pow(2, maxPower)
	fmt.Println(maxPower)
	fmt.Println(rem)
	
	if(rem!=0){
		for i:=maxPower; i >= 0 ; i-- {
			if(powCount - math.Pow(2,i) >= 0){
          			factors[j] = math.Pow(2,i)
			        powCount -= math.Pow(2,i)
				j++
				fmt.Printf("nextPow=%.0f\n", math.Pow(2,i))
				
				fmt.Println("j=", j)
        		}
		}
	}

}

