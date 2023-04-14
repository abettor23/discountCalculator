package main

import (
	"fmt"
)

func main() {
	fmt.Println("Добрый день! Вы находитесь в супермаркете, где вы устанавливаете цену!")
	fmt.Println("По нашим правилам, вы можете выбрать только 3 позции.")
	fmt.Println("Если сумма вашего чека превысит 10000р., мы сделаем вам скидку 10%!")
	fmt.Println("Какой товар хотите купить первым?")

	var itemOne string
	fmt.Scan(&itemOne)

	fmt.Println("Отлично! Сколько вы готовы заплатить за него?")

	var itemOnePrice int
	fmt.Scan(&itemOnePrice)

	fmt.Println("Теперь назовите второй товар:")

	var itemTwo string
	fmt.Scan(&itemTwo)

	fmt.Println("Ваша цена за второй товар?")

	var itemTwoPrice int
	fmt.Scan(&itemTwoPrice)

	fmt.Println("Теперь назовите последний товар:")

	var itemThree string
	fmt.Scan(&itemThree)

	fmt.Println("Ваша цена за последний товар?")

	var itemThreePrice int
	fmt.Scan(&itemThreePrice)

	check := itemOnePrice + itemTwoPrice + itemThreePrice

	fmt.Print("Итак, у вас в корзине:", itemOne, ", ", itemTwo, ", ", itemThree, ".\n")
	fmt.Println("Сумма к оплате:", check)

	if check > 10000 {
		check := (itemOnePrice + itemTwoPrice + itemThreePrice) - ((itemOnePrice+itemTwoPrice+itemThreePrice)/100)*10
		fmt.Println("Сумма к оплате, с учетом скидки:", check)
	}

	if check <= 10000 {
		fmt.Println("Скидка не применена! (Товаров, менее, чем на 10000).")
	}
}
