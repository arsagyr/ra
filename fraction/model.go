package fraction

// Структура типа "Fraction" - "Дробь"
type Fraction struct {
	numerator   int // Числитель
	denominator int // Знаменатель
}

type Fraction32 struct {
	numerator   int32 // Числитель
	denominator int32 // Знаменатель
}

type Fraction64 struct {
	numerator   int64 // Числитель
	denominator int64 // Знаменатель
}

func (f Fraction) GetNumerator() int {
	return f.numerator
}
func (f *Fraction) SetNumerator(a int) {
	f.numerator = a
}
func (f Fraction) GetDenominator() int {
	return f.denominator
}

func (f Fraction32) GetNumerator() int32 {
	return f.numerator
}
func (f *Fraction32) SetNumerator(a int32) {
	f.numerator = a
}
func (f Fraction32) GetDenominator() int32 {
	return f.denominator
}

func (f Fraction64) GetNumerator() int64 {
	return f.numerator
}
func (f *Fraction64) SetNumerator(a int64) {
	f.numerator = a
}
func (f Fraction64) GetDenominator() int64 {
	return f.denominator
}

// Метод упрощения дроби путем нахождения НОД с помощью алгоритма Евклида
func (f Fraction) Simplification() Fraction {
	var a, b, r int
	var i int = 0
	a = f.numerator
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
	return Fraction{
		denominator: f.denominator / b,
		numerator:   f.numerator /b,
	}
}

func (f Fraction32) Simplification() Fraction32 {
	var a, b, r int32
	var i int32 = 0
	a = f.numerator
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
		numerator:   f.numerator /b,
	}
}
func (f Fraction64) Simplification() Fraction64 {
	var a, b, r int64
	var i int64 = 0
	a = f.numerator
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
		numerator:   f.numerator /b,
	}
}
