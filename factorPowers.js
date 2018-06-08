function factorPowers(powInt){
/*
 =============================================
 Author:	  Aloyts
 Create date: 2018 June 06
 Description: Factoring powers of 2
 =============================================
*/
var maxPower = Math.floor(Math.log2(powInt));
var rem = (powInt - Math.pow(2, maxPower));

document.getElementById("text1").innerHTML += "maxPower=" + maxPower;

if(rem!=0)
{
	var factors = "";
    var powCount = powInt;
    
    for(var i=maxPower;i>=0;i--)
    {
    	
        if(powCount - Math.pow(2,i) >= 0)
        {
          factors += Math.pow(2,i) + ",";
          powCount -= Math.pow(2,i);
        }
    }
	factors = factors.substring(0, factors.length-1);  //remove trailing comma.
    
}

return factors;
}
