package fraction

import "log"

// Метод сравнения, где дробь f меньше числа a
func (f Fraction) IsSmallerThanINT(input any) bool {
	var b bool = false
	switch a := input.(type) {
	case uint:
		b = f.numerator < int(a*f.denominator)
	case int:
		b = f.numerator < (a * int(f.denominator))
	default:
		log.Panic("Error: wrong format")
	}
	return b
}

// Метод сравнения, где дробь f больше числа a
func (f Fraction) IsBiggerThanINT(input any) bool {
	var b bool = false
	switch a := input.(type) {
	case uint:
		b = f.numerator > int(a*f.denominator)
	case int:
		b = f.numerator > (a * int(f.denominator))
	default:
		log.Panic("Error: wrong format")
	}
	return b
}

// Метод сравнения, где дробь f равна числу a
func (f Fraction) IsEqualToINT(a int) bool {
	return f.numerator == (a * int(f.denominator))
}

// Метод сравнения, где дробь f1 меньше дроби f2
func (f1 Fraction) IsSmallerThanFraction(f2 Fraction) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * int(f2.denominator)) < (f2.numerator * int(f1.denominator)))
	} else {
		b = (f1.numerator < f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 больше дроби f2
func (f1 Fraction) IsBiggerThanFraction(f2 Fraction) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * int(f2.denominator)) > (f2.numerator * int(f1.denominator)))
	} else {
		b = (f1.numerator > f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 равна дроби f2
func (f1 Fraction) IsEqualToFraction(f2 Fraction) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * int(f2.denominator)) == (f2.numerator * int(f1.denominator)))
	} else {
		b = (f1.numerator == f2.numerator)
	}
	return b
}
