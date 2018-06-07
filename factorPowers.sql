SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		 Aloyts
-- Create date: 2018 June 06
-- Description: Factoring powers of 2
-- =============================================
ALTER PROCEDURE [dbo].[sp_powBO]
	@powInt bigint
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;


   declare @i int = floor(log(@powInt, 2))       /*highest power of 2 less than the argument*/
   declare @rem int = @powInt - power(2, @i)     /*remainder after subtracting the prior int*/
   /*
   select @i as 'highest power', @rem as 'remainder' , power(2, @i) as '2^'   /*return the initial values making up the power of 2 sum to the argument*/
   */

   select '2^' + cast(@i as varchar(16)) + ' = ' + cast(power(2, @i) as varchar(16)) as 'factors'

   while @i >= 1 and @rem > 0
   begin      
      select @i = floor(log(@rem, 2)) 
      select @rem = @rem - power(2, @i)
      /* select @i, @rem, power(2, @i) */
      select '2^' + cast(@i as varchar(16)) + ' = ' + cast(power(2, @i) as varchar(16)) as 'factors'
   end
  
END
