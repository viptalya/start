package lasagna

const OvenTime = 40

// Определяет, сколько минут необходимо находиться лазанье в духовке.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// Считает время приготовления слоев лазаньи.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// Считает проведенное лазаньей время в духовке.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
