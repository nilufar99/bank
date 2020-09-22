package stats

import "github.com/nilufar99/bank/pkg/bank/types"

// Avg рассчитывает средную сумму платежа
func Avg(payments []types.Payment) types.Money  {

	var sum types.Money
	
	for _, payment := range payments{
		sum += payment.Amount
		
	}

	return sum / types.Money(len(payments))
}

// TotalInCategory находит сумму покупок в опреденной категории.
func TotalInCategory(payments []types.Payment, category types.Category)types.Money {
	sum := types.Money(0)

	for _, payment := range payments{
		if payment.Category == category{
			sum += payment.Amount
		}
	}
	return sum
	
}