package fraction

// Метод для умножения дроби на целое число
func (f Fraction) MultiplyByInt(a int) Fraction {
	return Fraction{
		numerator:   a * f.numerator,
		denominator: f.denominator,
	}
}

// Метод для умножения дроби на другую дробь
func (f1 Fraction) MultiplyByFraction(f2 Fraction) Fraction {
	return Fraction{
		numerator:   f1.numerator * f2.numerator,
		denominator: f1.denominator * f2.denominator,
	}
}

// Метод для деления дроби на целое число
func (f Fraction) DivideByInt(a int) Fraction {
	var n int = f.numerator
	var d uint
	if a > 0 {
		d = uint(a) * f.denominator
	} else if a < 0 {
		n *= -1
		d = uint(a) * f.denominator
	} else {
		panic("Denominator cannot be zero")
	}
	return Fraction{
		numerator:   n,
		denominator: d,
	}
}

// Метод для деления дроби на другую дробь
func (f1 Fraction) DivideByFraction(f2 Fraction) Fraction {
	var n int = f1.numerator * int(f2.denominator)
	var d uint
	if f2.numerator > 0 {
		d = uint(f2.numerator) * f1.denominator
	} else if f2.numerator < 0 {
		n *= -1
		d = uint(f2.numerator) * f1.denominator
	} else {
		panic("Denominator cannot be zero")
	}
	return Fraction{
		numerator:   n,
		denominator: d,
	}
}

// Метод для вычитания дроби на целое число
func (f Fraction) SubtractByInt(a int) Fraction {
	return Fraction{
		numerator:   f.numerator - a*int(f.denominator),
		denominator: f.denominator,
	}
}

// Метод для вычитания дроби на другую дробь
func (f1 Fraction) SubtractByFraction(f2 Fraction) Fraction {
	var a int
	var b uint
	if f1.denominator != f2.denominator {
		a = (f1.numerator*int(f2.denominator) - (f2.numerator * int(f1.denominator)))
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator - f2.numerator
		b = f1.denominator
	}
	return Fraction{
		numerator:   a,
		denominator: b,
	}
}

// Метод для суммирования дроби на целое число
func (f Fraction) SumByInt(a int) Fraction {
	return Fraction{
		numerator:   f.numerator + a*int(f.denominator),
		denominator: f.denominator,
	}
}

// Методы для суммирования дроби на другую дробь
func (f1 Fraction) SumByFraction(f2 Fraction) Fraction {
	var a int
	var b uint
	if f1.denominator != f2.denominator {
		a = (f1.numerator * int(f2.denominator)) + (f2.numerator * int(f1.denominator))
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator + f2.numerator
		b = f1.denominator
	}
	return Fraction{
		numerator:   a,
		denominator: b,
	}
}

func SumFractions(f1 Fraction, f2 Fraction) Fraction {
	var a int
	var b uint
	if f1.denominator != f2.denominator {
		a = (f1.numerator * int(f2.denominator)) + (f2.numerator * int(f1.denominator))
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator + f2.numerator
		b = f1.denominator
	}
	return Fraction{
		numerator:   a,
		denominator: b,
	}
}
