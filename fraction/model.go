package fraction

// Структура типа "Fraction" - "Дробь"
type Fraction struct {
	numerator   int  // Числитель
	denominator uint // Знаменатель
}

type Fraction32 struct {
	numerator   int32  // Числитель
	denominator uint32 // Знаменатель
}

type Fraction64 struct {
	numerator   int64  // Числитель
	denominator uint64 // Знаменатель
}

func (f Fraction) GetNumerator() int {
	return f.numerator
}
func (f *Fraction) SetNumerator(a int) {
	f.numerator = a
}
func (f Fraction) GetDenominator() uint {
	return f.denominator
}

func (f Fraction32) GetNumerator() int32 {
	return f.numerator
}
func (f *Fraction32) SetNumerator(a int32) {
	f.numerator = a
}
func (f Fraction32) GetDenominator() uint32 {
	return f.denominator
}

func (f Fraction64) GetNumerator() int64 {
	return f.numerator
}
func (f *Fraction64) SetNumerator(a int64) {
	f.numerator = a
}
func (f Fraction64) GetDenominator() uint64 {
	return f.denominator
}

// Метод упрощения дроби путем нахождения НОД с помощью алгоритма Евклида
func (f *Fraction) Simplification() {
	var a, b, r, i uint
	i = 0
	if f.numerator < 0 {
		f.numerator *= -1
	}
	a = uint(f.numerator)
	b = f.denominator
	if a < b {
		i = a
		a = b
		b = i
	}
	r = a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}

	f.denominator = f.denominator / b
	f.numerator = f.numerator / int(b)
}

func (f Fraction32) Simplification() Fraction32 {
	var a, b, r, i uint32
	i = 0
	if f.numerator < 0 {
		f.numerator *= -1
	}
	a = uint32(f.numerator)
	b = f.denominator
	if a < b {
		i = a
		a = b
		b = i
	}
	r = a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return Fraction32{
		denominator: f.denominator / b,
		numerator:   f.numerator / int32(b),
	}
}
func (f Fraction64) Simplification() Fraction64 {
	var a, b, r, i uint64
	i = 0
	if f.numerator < 0 {
		f.numerator *= -1
	}
	a = uint64(f.numerator)
	b = f.denominator
	if a < b {
		i = a
		a = b
		b = i
	}
	r = a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return Fraction64{
		denominator: f.denominator / b,
		numerator:   f.numerator / int64(b),
	}
}
