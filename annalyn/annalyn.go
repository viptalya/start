package annalyn

// Проверяет, можно ли провести быструю атаку.
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	} else {
		return true
	}
}

// Проверяет, есть ли смысл проводить слежку.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	}

	return false
}

// Проверяет, можно ли подать сигнал заключенному.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
		return true
	}

	return false
}

// Проверяет, можно ли освободить заключенного.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if !archerIsAwake && petDogIsPresent {
		return true
	} else if !knightIsAwake && !archerIsAwake && prisonerIsAwake && !petDogIsPresent {
		return true
	} else {
		return false
	}
}
