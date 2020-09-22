package stats

import (
	"fmt"
	"github.com/nilufar99/bank/pkg/bank/types"
)

func ExampleAvg()  {
	money := []types.Payment{
		{
			ID: 1,
			Amount: 100,
			Category: "auto",
		},{
			ID: 2,
			Amount: 200,
			Category: "auto",
		},{
			ID: 3,
			Amount: 300,
			Category: "auto",
		},
	}
	result := Avg(money)

	fmt.Println(result)

	//Output: 
	// 200

	
}


func ExampleTotalInCategory()  {
	money := []types.Payment{
		{
			ID: 1,
			Amount: 100,
			Category: "auto",

		},{
			ID: 2,
			Amount: 200,
			Category: "auto",

		},
		{
			ID: 2,
			Amount: 200,
			Category: "shop",

		},
	}
	total := TotalInCategory(money, "auto")
	fmt.Println(total)

	// Output: 
	// 300
}