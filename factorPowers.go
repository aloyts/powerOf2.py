/*
 =============================================
 Author:	  Aloyts
 Create date: 2018 August 16
 Description: Factoring powers of 2
 =============================================
*/

import 'dart:math';

void main() {
  
  print(factorPowers(15));
}

String factorPowers(powInt){

  int maxPower = (log(powInt)/log(2)).floor();
  
  print("maxPower=" + maxPower.toString());
  
  int rem = (powInt - pow(2, maxPower));
	
  print("Remainder=" + rem.toString());
  
  String factors = "";
  
  if(rem!=0)
  {
    
    int powCount = powInt;

      for(int i=maxPower;i>=0;i--)
      {

          if(powCount - pow(2,i) >= 0)
          {
            factors += pow(2,i).toString() + ",";
            powCount -= pow(2,i);
          }
      }
    factors = factors.substring(0, factors.length-1);  //remove trailing comma.
    
}

return factors;
}